package response

import (
	"project3/eventapp/features/events"
)

type Event struct {
	ID     int    `json:"id" form:"id"`
	Name   string `json:"name" form:"name"`
	Detail string `json:"detail" form:"detail"`
	Photo  string `json:"photo" form:"photo"`
	URL    string `json:"url" form:"url"`
	Stock  int    `json:"stock" form:"stock"`
	Price  int    `json:"price" form:"price"`
	UserID int    `json:"userid" form:"userid"`
}

func FromCore(data events.Core) Event {
	return Event{
		ID:     data.ID,
		Name:   data.Name,
		Detail: data.EventDetail,
		Photo:  data.Photo,
		URL:    data.PhotoUrl,
		Stock:  data.Stock,
		Price:  data.Price,
		UserID: data.UserID,
	}
}

func FromCoreList(data []events.Core) []Event {
	result := []Event{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
