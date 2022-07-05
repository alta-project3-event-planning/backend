package data

import (
	"project3/eventapp/features/events"
	"project3/eventapp/features/users/data"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name       string    `json:"name" form:"name"`
	Detail     string    `json:"detail" form:"detail"`
	URL        string    `json:"url" form:"url"`
	Date       string	 `json:"time" form:"time"`
	Performers string    `json:"performers" form:"performers"`
	HostedBy   string    `json:"hostedby" form:"hostedby"`
	City       string    `json:"city" form:"city"`
	Location   string    `json:"location" form:"location"`
	UserID     int
	User       data.User
}

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

//DTO

func (data *Event) toCore() events.Core {
	return events.Core{
		ID:          int(data.ID),
		Name:        data.Name,
		EventDetail: data.Detail,
		Url:         data.URL,
		City:        data.City,
		HostedBy:    data.HostedBy,
		Performers:  data.Performers,
		Location:    data.Location,
		IDUser:      data.UserID,
	}
}

func ToCoreList(data []Event) []events.Core {
	result := []events.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core events.Core) Event {
	return Event{
		Name:       core.Name,
		Detail:     core.EventDetail,
		URL:        core.Url,
		HostedBy:   core.HostedBy,
		Performers: core.Performers,
		UserID:     core.IDUser,
		User: data.User{
			Name:  core.User.Name,
			Email: core.User.Email,
		},
	}
}

func toCore(data Event) events.Core {
	return data.toCore()
}
