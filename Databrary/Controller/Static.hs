{-# LANGUAGE OverloadedStrings #-}
module Databrary.Controller.Static
  ( staticPublicFile
  ) where

import qualified Data.Text as T

import Databrary.Action.Route
import Databrary.Action
import Databrary.HTTP.File

staticPublicFile :: StaticPath -> AppRAction
staticPublicFile sp = action GET ("public" :: T.Text, sp) $ do
  serveStaticFile "public" sp
