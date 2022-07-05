package data

import (
	_event "project3/eventapp/features/events/data"
	"project3/eventapp/features/participants"
	_user "project3/eventapp/features/users/data"

	"gorm.io/gorm"
)

type Participant struct {
	gorm.Model
	UserID  int
	EventID int
	Event   _event.Event
	User    _user.User
}

// type Event struct {
// 	gorm.Model
// 	Name        string    `json:"name" form:"name"`
// 	Detail      string    `json:"detail" form:"detail"`
// 	URL         string    `json:"url" form:"url"`
// 	Date        time.Time `json:"time" form:"time"`
// 	Performers  string    `json:"performers" form:"performers"`
// 	HostedBy    string    `json:"hostedby" form:"hostedby"`
// 	City        string    `json:"city" form:"city"`
// 	Location    string    `json:"location" form:"location"`
// 	UserID      int
// 	User        User
// 	Participant []Participant
// }

// type User struct {
// 	gorm.Model
// 	Name     string
// 	Email    string
// 	Password string
// 	Event    []Event
// }

func (data *Participant) toCore() participants.Core {
	return participants.Core{
		ID:        int(data.ID),
		IdEvent:   data.EventID,
		IdUser:    data.EventID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func ToCoreList(data []Participant) []participants.Core {
	result := []participants.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}
