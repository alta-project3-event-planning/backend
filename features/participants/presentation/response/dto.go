package response

import (
	"time"

	"project3/eventapp/features/participants"
)

type Participant struct {
	ID      int `json:"id" form:"id"`
	IdUser  int `json:"user_id" form:"user_id"`
	IdEvent int `json:"event_id" form:"event_id"`
	// Event   _event.Event
}

type Event struct {
	ID         int       `json:"id_event" form:"id_event"`
	Name       string    `json:"name" form:"name"`
	Detail     string    `json:"detail" form:"detail"`
	URL        string    `json:"url" form:"url"`
	Date       time.Time `json:"time" form:"time"`
	Performers string    `json:"performers" form:"performers"`
	HostedBy   string    `json:"hostedby" form:"hostedby"`
	City       string    `json:"city" form:"city"`
	Location   string    `json:"location" form:"location"`
	UserID     int       `json:"user_id" form:"user_id"`
}

func FromCore(data participants.Event) Event {
	return Event{
		ID:         data.ID,
		URL:        data.Url,
		Name:       data.Name,
		HostedBy:   data.HostedBy,
		Performers: data.Performers,
		Date:       data.Date,
		City:       data.City,
		Location:   data.Location,
		Detail:     data.EventDetail,
	}
}

func FromCoreList(data []participants.Event) []Event {
	result := []Event{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
