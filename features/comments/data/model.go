package data

import (
	"project3/eventapp/features/comments"
	_event "project3/eventapp/features/events/data"
	_user "project3/eventapp/features/users/data"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	EventID int
	UserID  int
	Comment string `gorm:"text" json:"text" binding:"required"`
	Event   _event.Event
	User    _user.User
}

func (data *Comment) toCore() comments.Core {
	return comments.Core{
		ID:      int(data.ID),
		EventID: data.EventID,
		UserID:  data.UserID,
		User: comments.User{
			Name:   data.User.Name,
			Avatar: data.User.URL,
		},
		Comment:   data.Comment,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func ToCoreList(data []Comment) []comments.Core {
	result := []comments.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core comments.Core) Comment {
	return Comment{
		UserID:  core.UserID,
		EventID: core.EventID,
		Comment: core.Comment,
	}
}
