package data

import (
	"project3/eventapp/features/events"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name   string `json:"name" form:"name"`
	Detail string `json:"detail" form:"detail"`
	Photo  string `json:"photo" form:"photo"`
	URL    string `json:"url" form:"url"`
	Stock  int    `json:"stock" form:"stock"`
	Price  int    `json:"price" form:"price"`
	UserID int
	User   User
}

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Event  []Event
}

//DTO

func (data *Event) toCore() events.Core {
	return events.Core{
		ID:            int(data.ID),
		Name:          data.Name,
		EventDetail: data.Detail,
		Photo:         data.Photo,
		PhotoUrl:      data.URL,
		Stock:         data.Stock,
		Price:         data.Price,
		UserID:        data.UserID,
		User: events.User{
			ID:    int(data.User.ID),
			Name:  data.User.Name,
			Email: data.User.Email,
		},
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
		Name:   core.Name,
		Detail: core.EventDetail,
		Photo:  core.Photo,
		URL:    core.PhotoUrl,
		Stock:  core.Stock,
		Price:  core.Price,
		UserID: core.UserID,
		User: User{
			Name:  core.User.Name,
			Email: core.User.Email,
		},
	}
}

func toCore(data Event) events.Core {
	return data.toCore()
}
