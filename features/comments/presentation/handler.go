package presentation

import (
  "project3/eventapp/features/comments"
  "net/http"
  _request_comment "project3/eventapp/features/comments/presentation/request"
  
  "project3/eventapp/helper"
	"project3/eventapp/middlewares"
	// "strconv"

	"github.com/labstack/echo/v4"
)

type CommentHandler struct {
	commentBusiness comments.Business
}

func NewCommentHandler(business comments.Business) *CommentHandler {
	return &CommentHandler{
		commentBusiness: business,
	}
}

func (h *CommentHandler) Add(c echo.Context) error {
  userID_token, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get user id",
		})
	}
	
	comment := _request_comment.Comment{}
	err_bind := c.Bind(&comment)
	if err_bind != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to bind insert data"))
	}
	
	commentCore := _request_comment.ToCore(comment)
	commentCore.IdUser = userID_token
	
	row, err := h.commentBusiness.AddComment(commentCore)
	if err != nil {
	  return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed insert your comment"))
	}
	if row == 0 {
	  return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed insert your comment"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success insert your comment"))
}