package request

import (
	"project3/eventapp/features/comments"
)

type Comment struct {
	EventID int    `json:"event_id" form:"event_id"`
	UserID  int    `json:"user_id" form:"user_id"`
	Comment string `json:"comment" form:"comment"`
}

func ToCore(req Comment) comments.Core {
	commentCore := comments.Core{
		EventID: req.EventID,
		UserID:  req.UserID,
		Comment: req.Comment,
	}
	return commentCore
}
