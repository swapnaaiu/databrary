{-# LANGUAGE TemplateHaskell, ScopedTypeVariables #-}
module Model.VolumeAccess.SQL
  ( selectVolumeAccess
  , selectVolumeAccessParty
  , selectPartyVolumeAccess
  , updateVolumeAccess
  , insertVolumeAccess
  , deleteVolumeAccess
  , selectVolumeActivity
  ) where

import qualified Language.Haskell.TH as TH

import Model.Permission.Types
import Model.Party.Types
import Model.Volume.Types
import Model.SQL.Select
import Model.Audit.SQL
import Model.Party.SQL
import Model.Volume.SQL
import Model.VolumeAccess.Types

volumeAccessRow :: Selector -- ^ @'Party' -> 'Volume' -> 'VolumeAccess'@
volumeAccessRow = selectColumns 'VolumeAccess "volume_access" ["individual", "children", "sort", "share_full"]

selectVolumeAccess :: TH.Name -- ^ 'Volume'
  -> TH.Name -- ^ 'Identity'
  -> Selector -- ^ 'VolumeAccess'
selectVolumeAccess vol ident = selectMap (`TH.AppE` TH.VarE vol) $ selectJoin '($)
  [ volumeAccessRow
  , joinOn ("volume_access.party = party.id AND volume_access.volume = ${volumeId $ volumeRow " ++ nameRef vol ++ "}")
    $ selectAuthParty ident
  ]

makeVolumeAccessParty :: Party -> Maybe (Party -> Volume -> VolumeAccess) -> Volume -> VolumeAccess
makeVolumeAccessParty p Nothing v = VolumeAccess PermissionNONE PermissionNONE Nothing (getShareFullDefault p PermissionNONE) p v
makeVolumeAccessParty p (Just af) v = af p v

selectVolumeAccessParty :: TH.Name -- ^ 'Volume'
  -> TH.Name -- ^ 'Identity'
  -> Selector -- ^ 'VolumeAccess'
selectVolumeAccessParty vol ident = selectMap (`TH.AppE` TH.VarE vol) $ selectJoin 'makeVolumeAccessParty
  [ selectAuthParty ident
  , maybeJoinOn ("party.id = volume_access.party AND volume_access.volume = ${volumeId $ volumeRow " ++ nameRef vol ++ "}")
    volumeAccessRow
  ]

selectPartyVolumeAccess :: TH.Name -- ^ 'Party'
  -> TH.Name -- ^ 'Identity'
  -> Selector -- ^ 'VolumeAccess'
selectPartyVolumeAccess p ident = selectJoin '($)
  [ selectMap (`TH.AppE` TH.VarE p) volumeAccessRow
  , joinOn ("volume_access.volume = volume.id AND volume_access.party = ${partyId $ partyRow " ++ nameRef p ++ "}")
    $ selectVolume ident
  ]

type ColumnName = String
type ColumnValueExp = String

volumeAccessKeys :: String -- ^ @'VolumeAccess'@
  -> [(ColumnName, ColumnValueExp)]
volumeAccessKeys a =
  [ ("volume", "${volumeId $ volumeRow $ volumeAccessVolume " ++ a ++ "}")
  , ("party", "${partyId $ partyRow $ volumeAccessParty " ++ a ++ "}")
  ]

volumeAccessSets :: String -- ^ @'VolumeAccess'@
  -> [(ColumnName, ColumnValueExp)]
volumeAccessSets a =
  [ ("individual", "${volumeAccessIndividual " ++ a ++ "}")
  , ("children", "${volumeAccessChildren " ++ a ++ "}")
  , ("sort", "${volumeAccessSort " ++ a ++ "}")
  , ("share_full", "${volumeAccessShareFull " ++ a ++ "}")
  ]

type SQLFilterClause = String

noReturning :: Maybe SelectOutput
noReturning = Nothing

updateVolumeAccess :: TH.Name -- ^ @'AuditIdentity'@
  -> TH.Name -- ^ @'VolumeAccess'@
  -> TH.ExpQ
updateVolumeAccess ident a = auditUpdate ident "volume_access"
  (volumeAccessSets as)
  (whereEq $ volumeAccessKeys as :: SQLFilterClause)
  noReturning
  where as = nameRef a

insertVolumeAccess :: TH.Name -- ^ @'AuditIdentity'@
  -> TH.Name -- ^ @'VolumeAccess'@
  -> TH.ExpQ
insertVolumeAccess ident a = auditInsert ident "volume_access"
  (volumeAccessKeys as ++ volumeAccessSets as)
  noReturning
  where as = nameRef a

deleteVolumeAccess :: TH.Name -- ^ @'AuditIdentity'@
  -> TH.Name -- ^ @'VolumeAccess'@
  -> TH.ExpQ
deleteVolumeAccess ident a = auditDelete ident "volume_access"
  (whereEq $ volumeAccessKeys as)
  Nothing
  where as = nameRef a

selectVolumeActivity :: TH.Name -- ^@'Identity'@
  -> Selector -- ^ @('Timestamp', 'Volume')@
selectVolumeActivity ident = selectJoin '(,)
  [ selectAuditActivity "volume_access"
  , joinOn "audit.volume = volume.id"
    $ selectVolume ident
  ]
