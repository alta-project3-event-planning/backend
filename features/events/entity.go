package events

import (
	"time"
)

type Core struct {
	ID            int
	Name          string
	EventName   string
	EventDetail string
	Stock         int
	Price         int
	Photo         string
	PhotoUrl      string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	UserID        int
	User          User
}

type User struct {
	ID    int
	Name  string
	Email string
}

type Business interface {
	GetAllEvent() (data []Core, err error)
	GetEventByID(param int) (data Core, err error)
	InsertEvent(dataReq Core) (err error)
	DeleteEventByID(id int, userId int) (err error)
	UpdateEventByID(dataReq Core, id int, userId int) (err error)
	GetEventByUserID(id_user int) (data []Core, err error)
}

type Data interface {
	SelectData() (data []Core, err error)
	SelectDataByID(param int) (data Core, err error)
	InsertData(dataReq Core) (err error)
	DeleteDataByID(id int, userId int) (err error)
	UpdateDataByID(dataReq map[string]interface{}, id int, userId int) (err error)
	SelectDataByUserID(id_user int) (data []Core, err error)
}
