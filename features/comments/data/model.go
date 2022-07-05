package data

import (
	"project3/eventapp/features/comments"
	_event "project3/eventapp/features/events/data"
	_user "project3/eventapp/features/users/data"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	IdEvent int
	IdUser  int
	Comment string `gorm:"text" json:"text" binding:"required"`
	Event   _event.Event
	User    _user.User
}

func (data *Comment) toCore() comments.Core {
	return comments.Core{
		ID:        int(data.ID),
		IdEvent:   data.IdEvent,
		IdUser:    data.IdUser,
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

// func fromCore(core comment.Core) Comment {
// 	return Comment{

// 	}
// }
