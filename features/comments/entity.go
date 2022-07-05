package comments

import (
	"time"
)

type Core struct {
	ID        int
	IdEvent   int
	IdUser    int
	Comment   string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User
	Event     Event
}

type Event struct {
	ID          int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	EventName   string
	EventDetail string
	Url         string
	Date        time.Time
	Performers  string
	HostedBy    string
	City        string
	Location    string
	IDUser      int
	User        User
}

type User struct {
	ID    int
	Name  string
	Email string
}

type Business interface {
	AddComment(data Core) (row int, err error)
}

type Data interface {
	Add(data Core) (row int, err error)
}
