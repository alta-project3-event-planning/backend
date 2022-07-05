package response

import (
	"project3/eventapp/features/events"
	"time"
)

type Event struct {
	ID         int       `json:"id_event" form:"id_event"`
	Name       string    `json:"name" form:"name"`
	Detail     string    `json:"details" form:"details"`
	URL        string    `json:"image_url" form:"imageurl"`
	HostedBy   string    `json:"hostedby" form:"hostedby"`
	Performers string    `json:"performers" form:"performers"`
	Date       time.Time `json:"date" form:"date"`
	City       string    `json:"city" form:"city"`
	Location   string    `json:"location" form:"location"`
}

func FromCore(data events.Core) Event {
	return Event{
		ID:         data.ID,
		Name:       data.Name,
		Detail:     data.EventDetail,
		URL:        data.Url,
		HostedBy:   data.HostedBy,
		Performers: data.Performers,
		Date:       data.Date,
		City:       data.City,
		Location:   data.Location,
	}
}

func FromCoreList(data []events.Core) []Event {
	result := []Event{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
