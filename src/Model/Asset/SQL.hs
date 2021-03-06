{-# LANGUAGE TemplateHaskell, OverloadedStrings #-}
module Model.Asset.SQL
  ( selectAssetRow
  , selectAsset
  -- , insertAsset
  -- , updateAsset
  , makeAssetRow
  , setAssetId
  -- for expanded queries
  ) where

import qualified Data.ByteString as BS
import Data.Int (Int64)
import qualified Data.Text as T
import qualified Language.Haskell.TH as TH

import Model.Offset
import Model.Format
import Model.Id.Types
import Model.Release.Types
import Model.SQL.Select
import Model.Volume.SQL
import Model.Asset.Types

makeAssetRow :: Id Asset -> Id Format -> Maybe Release -> Maybe Offset -> Maybe T.Text -> Maybe BS.ByteString -> Maybe Int64 -> AssetRow
makeAssetRow i = AssetRow i . getFormat'

selectAssetRow :: Selector -- ^ @'AssetRow'@
selectAssetRow = selectColumns 'makeAssetRow "asset" ["id", "format", "release", "duration", "name", "sha1", "size"]

selectAsset :: TH.Name -- ^ @'Identity'@
  -> Selector -- ^ @'Asset'@
selectAsset ident = selectJoin 'Asset
  [ selectAssetRow
  , joinOn "asset.volume = volume.id" $ selectVolume ident
  ]

setAssetId :: Asset -> Id Asset -> Asset
setAssetId a i = a{ assetRow = (assetRow a){ assetId = i } }

{-
assetKeys :: String -- ^ @'Asset'@
  -> [(String, String)]
assetKeys r =
  [ ("id", "${assetId $ assetRow " ++ r ++ "}") ]

assetSets :: String -- ^ @'Asset'@
  -> [(String, String)]
assetSets a =
  [ ("volume", "${volumeId $ volumeRow $ assetVolume " ++ a ++ "}")
  , ("format", "${formatId $ assetFormat $ assetRow " ++ a ++ "}")
  , ("release", "${assetRelease $ assetRow " ++ a ++ "}")
  , ("duration", "${assetDuration $ assetRow " ++ a ++ "}")
  , ("name", "${assetName $ assetRow " ++ a ++ "}")
  , ("sha1", "${assetSHA1 $ assetRow " ++ a ++ "}")
  , ("size", "${assetSize $ assetRow " ++ a ++ "}")
  ]

insertAsset :: TH.Name -- ^ @'AuditIdentity'@
  -> TH.Name -- ^ @'Asset'@
  -> TH.ExpQ -- ^ @'Asset'@
insertAsset ident a = auditInsert ident "asset"
  (assetSets (nameRef a))
  (Just $ selectOutput $ selectMap ((TH.VarE 'setAssetId `TH.AppE` TH.VarE a) `TH.AppE`) $ selectColumn "asset" "id")

updateAsset :: TH.Name -- ^ @'AuditIdentity'@
  -> TH.Name -- ^ @'Asset'@
  -> TH.ExpQ -- ^ @()@
updateAsset ident a = auditUpdate ident "asset"
  (assetSets (nameRef a))
  (whereEq $ assetKeys (nameRef a))
  Nothing
-}
