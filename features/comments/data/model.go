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

// type Event struct {
// 	gorm.Model
// 	Name       string    `json:"name" form:"name"`
// 	Detail     string    `json:"detail" form:"detail"`
// 	URL        string    `json:"url" form:"url"`
// 	Date       time.Time `json:"time" form:"time"`
// 	Performers string    `json:"performers" form:"performers"`
// 	HostedBy   string    `json:"hostedby" form:"hostedby"`
// 	City       string    `json:"city" form:"city"`
// 	Location   string    `json:"location" form:"location"`
// 	UserID     int
// 	User       User
// 	Comment    []Comment
// }

// type User struct {
// 	gorm.Model
// 	Name     string
// 	Email    string
// 	Password string
// 	Event    []Event
// }

func (data *Comment) toCore() comments.Core {
	return comments.Core{
		ID:        int(data.ID),
		IdEvent:   data.EventID,
		IdUser:    data.UserID,
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
