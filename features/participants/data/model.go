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

func (data *Participant) toCore() participants.Core {
	return participants.Core{
		ID:        int(data.ID),
		EventID:   data.EventID,
		UserID:    data.EventID,
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
