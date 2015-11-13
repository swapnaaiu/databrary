{-# LANGUAGE OverloadedStrings #-}
module Databrary.View.Container
  ( htmlContainerEdit
  ) where

import Data.Foldable (fold)
import Data.Monoid ((<>), mempty)

import Databrary.Model.Volume
import Databrary.Model.Container
import Databrary.Model.Slot
import Databrary.Action.Types
import Databrary.Action
import Databrary.View.Form

import {-# SOURCE #-} Databrary.Controller.Container

htmlContainerForm :: Maybe Container -> FormHtml f
htmlContainerForm cont = do
  field "name" $ inputText (containerName . containerRow =<< cont)
  field "date" $ inputDate (containerDate . containerRow =<< cont)
  field "release" $ inputEnum False (containerRelease =<< cont)

htmlContainerEdit :: Either Volume Container -> RequestContext -> FormHtml f
htmlContainerEdit (Left v)  = htmlForm "Create container" createContainer (HTML, volumeId $ volumeRow v) (htmlContainerForm Nothing) (const mempty)
htmlContainerEdit (Right c) = htmlForm ("Edit container " <> fold (containerName $ containerRow c)) postContainer (HTML, containerSlotId $ containerId $ containerRow c) (htmlContainerForm $ Just c) (const mempty)
