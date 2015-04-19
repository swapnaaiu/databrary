{-# LANGUAGE OverloadedStrings #-}
module Databrary.Controller.Comment
  ( postComment
  ) where

import Data.Maybe (maybeToList)

import Databrary.Ops
import Databrary.Model.Permission
import Databrary.Model.Id
import Databrary.Model.Slot
import Databrary.Model.Comment
import Databrary.HTTP.Form.Deform
import Databrary.Action
import Databrary.Controller.Permission
import Databrary.Controller.Form
import Databrary.Controller.Slot
import Databrary.View.Comment

postComment :: API -> Id Slot -> AppRAction
postComment api si = action POST (api, si) $ withAuth $ do
  u <- authAccount
  s <- getSlot PermissionSHARED si
  c <- runForm (api == HTML ?> htmlCommentForm s) $ do
    text <- "text" .:> (deformRequired =<< deform)
    parent <- "parent" .:> deformNonEmpty deform
    return (blankComment u s)
      { commentText = text
      , commentParents = maybeToList parent
      }
  c' <- addComment c
  case api of
    JSON -> okResponse [] $ commentJSON c'
    HTML -> redirectRouteResponse [] $ viewSlot api $ slotId $ commentSlot c'
