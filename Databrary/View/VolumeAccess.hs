{-# LANGUAGE OverloadedStrings #-}
module Databrary.View.VolumeAccess
  ( htmlVolumeAccessForm
  ) where

import Data.Monoid ((<>), mempty)

import Databrary.Action
import Databrary.Model.Party
import Databrary.Model.Volume
import Databrary.Model.VolumeAccess
import Databrary.Controller.Paths
import Databrary.View.Form

import {-# SOURCE #-} Databrary.Controller.VolumeAccess

htmlVolumeAccessForm :: VolumeAccess -> AuthRequest -> FormHtml f
htmlVolumeAccessForm a@VolumeAccess{ volumeAccessVolume = vol, volumeAccessParty = p } = htmlForm
  ("Access to " <> volumeName vol <> " for " <> partyName p)
  postVolumeAccess (HTML, (volumeId vol, VolumeAccessTarget (partyId p)))
  (do
    field "individual" $ inputEnum $ Just $ volumeAccessIndividual a
    field "children" $ inputEnum $ Just $ volumeAccessChildren a
    field "delete" $ inputCheckbox False)
  (const mempty)
