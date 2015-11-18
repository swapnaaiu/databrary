module Databrary.Model.Activity.Types
  ( ActivityTarget(..)
  , Activity(..)
  ) where

import qualified Data.ByteString as BS

import Databrary.Model.Audit.Types
import Databrary.Model.Id.Types
import Databrary.Model.Release.Types
import Databrary.Model.Party.Types
import Databrary.Model.Authorize.Types
import Databrary.Model.Volume.Types
import Databrary.Model.VolumeAccess.Types
import Databrary.Model.Container.Types
import Databrary.Model.Slot.Types
import Databrary.Model.Record.Types
import Databrary.Model.RecordSlot.Types
import Databrary.Model.Asset.Types
import Databrary.Model.AssetSlot.Types

data ActivityTarget
  = ActivityParty       { activityPartyRow :: !PartyRow }
  | ActivityAccount
    { activityAccountId :: !(Id Party)
    , activityAccountEmail :: !BS.ByteString
    , activityAccountPassword :: !(Maybe BS.ByteString)
    }
  | ActivityAuthorize   { activityAuthorize :: !Authorize }
  | ActivityVolume      { activityVolumeRow :: !VolumeRow }
  | ActivityAccess      { activityAccess :: !VolumeAccess }
  | ActivityContainer   { activityContainer :: !ContainerRow }
  | ActivityRelease
    { activityReleaseSlotId :: !(Id Slot)
    , activityRelease :: !Release
    }
  | ActivityRecord      { activityRecordRow :: !RecordRow }
  | ActivityRecordSlot  { activityRecordSlot :: !RecordSlot }
  | ActivityAsset       { activityAssetRow :: !AssetRow }
  | ActivityAssetSlot   { activityAssetSlot :: !(Id AssetSlot) }
  | ActivityExcerpt
    { activityExcerptId :: !(Id AssetSlot)
    , activityExcerptRelease :: !(Maybe Release)
    }

data Activity = Activity
  { activityAudit :: !Audit
  , activityUser :: !PartyRow
  , activityTarget :: !ActivityTarget
  , activityPrev :: Maybe ActivityTarget
  }

