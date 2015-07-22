{-# LANGUAGE OverloadedStrings #-}
module Databrary.View.Login
  ( htmlLogin
  , htmlUserForm
  ) where

import Data.Monoid (mempty)

import Databrary.Model.Party.Types
import Databrary.HTTP.Form.View
import Databrary.Action
import Databrary.View.Form

import {-# SOURCE #-} Databrary.Controller.Login

htmlLogin :: AuthRequest -> FormHtml f
htmlLogin = htmlForm "Login" postLogin HTML (do
  field "email" $ inputText (Nothing :: Maybe String)
  field "password" inputPassword
  field "superuser" $ inputCheckbox False)
  (const mempty)

htmlUserForm :: Account -> AuthRequest -> FormHtml f
htmlUserForm a = htmlForm "Edit account" postUser HTML (do
  field "auth" $ inputPassword
  field "email" $ inputText $ Just $ accountEmail a
  "password" .:> do
    field "once" inputPassword
    field "again" inputPassword)
  (const mempty)
