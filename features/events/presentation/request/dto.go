package request

import "project3/eventapp/features/events"

type Event struct {
	Name   string `json:"name" form:"name"`
	Detail string `json:"detail" form:"detail"`
	Stock  int    `json:"stock" form:"stock"`
	Price  int    `json:"price" form:"price"`
}

func ToCore(eventReq Event) events.Core {
	eventCore := events.Core{
		Name:          eventReq.Name,
		EventDetail:	eventReq.Detail,
		Stock:         eventReq.Stock,
		Price:         eventReq.Price,
	}
	return eventCore
}
