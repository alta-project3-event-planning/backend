package participants

import "time"

type Core struct {
	ID        int
	UserID    int
	EventID   int
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
	UserID      int
	User        User
}

type User struct {
	ID    int
	Name  string
	Email string
}

type Business interface {
	AddParticipant(data Core) (row int, err error)
}

type Data interface {
	Add(data Core) (row int, err error)
}
