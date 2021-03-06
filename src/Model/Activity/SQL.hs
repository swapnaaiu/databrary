{-# LANGUAGE TemplateHaskell #-}
module Model.Activity.SQL
  ( selectActivityParty
  , selectActivityAccount
  , selectActivityAuthorize
  , selectActivityVolume
  , selectActivityAccess
  , selectActivityContainer
  , selectActivityRelease
  , selectActivityAsset
  , selectActivityAssetSlot
  , selectActivityExcerpt
  , activityQual
  ) where

import Data.List (stripPrefix)
import qualified Language.Haskell.TH as TH

import Model.SQL.Select
import Model.Audit.Types
import Model.Audit.SQL
import Model.Party.SQL
import Model.Authorize.SQL
import Model.Volume.SQL
import Model.VolumeAccess.SQL
import Model.Container.SQL
import Model.Slot.SQL
import Model.Release.SQL
import Model.Asset.SQL
import Model.Activity.Types

delim :: String -> Bool
delim "" = True
delim (' ':_) = True
delim (',':_) = True
delim _ = False

makeActivity :: Audit -> ActivityTarget -> Activity
makeActivity a x = Activity a x Nothing Nothing Nothing

targetActivitySelector :: String -> Selector -> Selector
targetActivitySelector t Selector{ selectOutput = o, selectSource = ts, selectJoined = (',':tj) }
  | Just s <- stripPrefix t ts, delim s, ts == tj =
    selector ("audit." ++ ts) $ OutputJoin False 'makeActivity [selectOutput (selectAudit t), o]
targetActivitySelector t Selector{ selectSource = ts } = error $ "targetActivitySelector " ++ t ++ ": " ++ ts

selectActivityParty :: Selector
selectActivityParty = targetActivitySelector "party" $
  selectMap (TH.ConE 'ActivityParty `TH.AppE`) selectPartyRow

selectActivityAccount :: Selector
selectActivityAccount = targetActivitySelector "account" $
  selectColumns 'ActivityAccount "account" ["id", "email", "password"]

selectActivityAuthorize :: TH.Name -> TH.Name -> Selector
selectActivityAuthorize p ident = targetActivitySelector "authorize" $
  selectMap (TH.ConE 'ActivityAuthorize `TH.AppE`) (selectAuthorizeChild p ident)

selectActivityVolume :: Selector
selectActivityVolume = targetActivitySelector "volume" $
  selectMap (TH.ConE 'ActivityVolume `TH.AppE`) selectVolumeRow

selectActivityAccess :: TH.Name -> TH.Name -> Selector
selectActivityAccess vol ident = targetActivitySelector "volume_access" $
  selectMap (TH.ConE 'ActivityAccess `TH.AppE`) (selectVolumeAccess vol ident)

selectActivityContainer :: Selector
selectActivityContainer = targetActivitySelector "container" $
  selectMap (TH.ConE 'ActivityContainer `TH.AppE`) selectContainerRow

selectActivityRelease :: Selector
selectActivityRelease = targetActivitySelector "slot_release" $
  addSelects 'ActivityRelease (selectSlotId "slot_release") [selectOutput releaseRow]

selectActivityAsset :: Selector
selectActivityAsset = targetActivitySelector "asset" $
  selectMap (TH.ConE 'ActivityAsset `TH.AppE`) selectAssetRow

selectActivityAssetSlot :: Selector
selectActivityAssetSlot = targetActivitySelector "slot_asset" $
  addSelects 'ActivityAssetSlot
    (selectColumn "slot_asset" "asset")
    [selectOutput $ selectSlotId "slot_asset"]

selectActivityExcerpt :: Selector
selectActivityExcerpt = targetActivitySelector "excerpt" $
  selectColumns 'ActivityExcerpt "excerpt" ["asset", "segment", "release"]

activityQual :: String
activityQual = "audit_action >= 'add' ORDER BY audit_time"
