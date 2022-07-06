package data

import (
	"project3/eventapp/features/participants"
	"time"

	"gorm.io/gorm"
)

type Participant struct {
	gorm.Model
	IdUser  int `json:"user_id" form:"user_id"`
	IdEvent int `json:"event_id" form:"event_id"`
	// Event   _event.Event
}

type Event struct {
	gorm.Model
	Name        string    `json:"name" form:"name"`
	Detail      string    `json:"detail" form:"detail"`
	URL         string    `json:"url" form:"url"`
	Date        time.Time `json:"time" form:"time"`
	Performers  string    `json:"performers" form:"performers"`
	HostedBy    string    `json:"hostedby" form:"hostedby"`
	City        string    `json:"city" form:"city"`
	Location    string    `json:"location" form:"location"`
	UserID      int       `json:"user_id" form:"user_id"`
	Participant Participant
}

func fromCore(core participants.Core) Participant {
	return Participant{
		IdUser:  core.IdUser,
		IdEvent: core.IdEvent,
	}
}

func (data *Event) toCore() participants.Event {
	return participants.Event{
		ID:          int(data.Participant.ID),
		Url:         data.URL,
		Name:        data.Name,
		HostedBy:    data.HostedBy,
		Performers:  data.Performers,
		Date:        data.Date,
		City:        data.City,
		Location:    data.Location,
		EventDetail: data.Detail,
	}
}

func ToCoreList(data []Event) []participants.Event {
	result := []participants.Event{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}
