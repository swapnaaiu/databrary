{-# LANGUAGE RecordWildCards, OverloadedStrings, ViewPatterns #-}
module Store.Transcode
  ( startTranscode
  , forkTranscode
  , stopTranscode
  , collectTranscode
  , transcodeEnabled
  ) where

import Control.Concurrent (ThreadId)
import Control.Monad (guard, unless, void)
import Control.Monad.IO.Class (liftIO)
import Control.Monad.Trans.Resource (InternalState)
import qualified Data.ByteString as BS
import qualified Data.ByteString.Builder as BSB
import qualified Data.ByteString.Char8 as BSC
import qualified Data.ByteString.Lazy.Char8 as BSLC
import Data.Maybe (fromMaybe, isNothing)
import Data.Monoid ((<>))
import System.Exit (ExitCode(..))
import Text.Read (readMaybe)
import qualified Web.Route.Invertible as R

import Ops
import Has (view, peek, peeks, focusIO, MonadHas)
import Service.DB (MonadDB)
import Service.Log
import HTTP.Route (routeURL)
import Model.Audit (MonadAudit)
import Model.Segment
import Model.Asset
import Model.AssetSlot
import Model.Format
import Model.Time
import Model.Transcode
import Files
import Service.Types (Secret)
import Store.Types
import Store.Temp
import Store.Asset
import Store.Transcoder
import Store.AV
import Store.Probe
import Action.Types
import Action.Run

import {-# SOURCE #-} Controller.Transcode

ctlTranscode :: (MonadDB c m, MonadHas Timestamp c m, MonadLog c m, MonadStorage c m) => Transcode -> TranscodeArgs -> m (ExitCode, String, String)
ctlTranscode tc args = do
  t <- peek
  Just ctl <- peeks storageTranscoder
  let args'
        = "-i" : show (transcodeId tc)
        : "-f" : BSC.unpack (head $ formatExtension $ assetFormat $ assetRow $ transcodeAsset tc)
        : args
  r@(c, o, e) <- liftIO $ runTranscoder ctl args'
  focusIO $ logMsg t ("transcode " ++ unwords args' ++ ": " ++ case c of { ExitSuccess -> "" ; ExitFailure i -> ": exit " ++ show i ++ "\n" } ++ o ++ e)
  return r

transcodeArgs :: (MonadStorage c m, MonadHas Secret c m, MonadAudit c m) => Transcode -> m TranscodeArgs
transcodeArgs t@Transcode{..} = do
  Just f <- getAssetFile (transcodeOrig t)
  req <- peek
  auth <- peeks $ transcodeAuth t
  fp <- liftIO $ unRawFilePath f
  return $
    [ "-s", fp
    , "-r", BSLC.unpack $ BSB.toLazyByteString $ routeURL (Just req) (R.requestActionRoute remoteTranscode (transcodeId t)) [("auth", Just auth)]
    , "--" ]
    ++ maybe [] (\l -> ["-ss", show l]) lb
    ++ maybe [] (\u -> ["-t", show $ u - fromMaybe 0 lb]) (upperBound rng)
    ++ transcodeOptions
  where
  rng = segmentRange transcodeSegment
  lb = lowerBound rng

startTranscode :: (MonadStorage c m, MonadHas Secret c m, MonadAudit c m, MonadHas Timestamp c m, MonadLog c m) => Transcode -> m (Maybe TranscodePID)
startTranscode tc = do
  tc' <- updateTranscode tc lock Nothing
  unless (transcodeProcess tc' == lock) $ fail $ "startTranscode " ++ show (transcodeId tc)
  findMatchingTranscode tc >>= maybe
    (do
      args <- transcodeArgs tc
      (r, out, err) <- ctlTranscode tc' args
      let pid = guard (r == ExitSuccess) >> readMaybe out
      _ <- updateTranscode tc' pid $ (isNothing pid `thenUse` out) <> (null err `unlessUse` err)
      return pid)
    (\(transcodeAsset -> match) -> do
      a <- changeAsset (transcodeAsset tc)
        { assetRow = (assetRow $ transcodeAsset tc)
          { assetSHA1 = assetSHA1 $ assetRow match
          , assetDuration = assetDuration $ assetRow match
          , assetSize = assetSize $ assetRow match
          }
        } Nothing
      void $ changeAssetSlotDuration a
      _ <- updateTranscode tc' Nothing (Just $ "reuse " ++ show (assetId $ assetRow match))
      return Nothing)
  where lock = Just (-1)

forkTranscode :: Transcode -> Handler ThreadId
forkTranscode tc = focusIO $ \ctx ->
  forkAction
    (startTranscode tc) ctx
    (either
      (\e -> logMsg (view ctx) ("forkTranscode: " ++ show e) (view ctx))
      (const $ return ()))

stopTranscode :: (MonadDB c m, MonadHas Timestamp c m, MonadLog c m, MonadStorage c m) => Transcode -> m Transcode
stopTranscode tc@Transcode{ transcodeProcess = Just pid } | pid >= 0 = do
  tc' <- updateTranscode tc Nothing (Just "aborted")
  (r, out, err) <- ctlTranscode tc ["-k", show pid]
  unless (r == ExitSuccess) $
    fail ("stopTranscode: " ++ out ++ err)
  return tc'
stopTranscode tc = return tc

collectTranscode
  :: (MonadHas AV c m, MonadDB c m, MonadHas InternalState c m, MonadLog c m, MonadStorage c m, MonadHas Timestamp c m, MonadAudit c m)
  => Transcode -> Int -> Maybe BS.ByteString -> String -> m ()
collectTranscode tc 0 sha1 logs = do
  tc' <- updateTranscode tc (Just (-2)) (Just logs)
  f <- makeTempFile (const $ return ())
  (r, out, err) <- ctlTranscode tc ["-c", BSC.unpack $ tempFilePath f]
  _ <- updateTranscode tc' Nothing (Just $ out ++ err)
  if r /= ExitSuccess
    then fail $ "collectTranscode " ++ show (transcodeId tc) ++ ": " ++ show r ++ "\n" ++ out ++ err
    else do
      av <- focusIO $ avProbe (tempFilePath f)
      unless (avProbeCheckFormat (assetFormat $ assetRow $ transcodeAsset tc) av)
        $ fail $ "collectTranscode " ++ show (transcodeId tc) ++ ": format error"
      let dur = avProbeLength av
      a <- changeAsset (transcodeAsset tc)
        { assetRow = (assetRow $ transcodeAsset tc)
          { assetSHA1 = sha1
          , assetDuration = dur
          }
        } (Just $ tempFilePath f)
      focusIO $ releaseTempFile f
      void $ changeAssetSlotDuration a
collectTranscode tc e _ logs =
  void $ updateTranscode tc Nothing (Just $ "exit " ++ show e ++ '\n' : logs)
