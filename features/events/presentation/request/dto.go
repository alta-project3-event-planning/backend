package request

import (
	"project3/eventapp/features/events"
	"time"
)

type Event struct {
	Name       string    `json:"name" form:"name"`
	Detail     string    `json:"details" form:"details"`
	HostedBy   string    `json:"hostedby" form:"hostedby"`
	Performers string    `json:"performers" form:"performers"`
	Date       time.Time `json:"date" form:"date"`
	City       string    `json:"city" form:"city"`
	Location   string    `json:"location" form:"location"`
}

func ToCore(eventReq Event) events.Core {
	eventCore := events.Core{
		Name:        eventReq.Name,
		EventDetail: eventReq.Detail,
		HostedBy:    eventReq.HostedBy,
		Performers:  eventReq.Performers,
		Date:        eventReq.Date,
		City:        eventReq.City,
		Location:    eventReq.Location,
	}
	return eventCore
}
