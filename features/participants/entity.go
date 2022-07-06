package participants

import "time"

type Core struct {
	ID        int
	IdUser    int
	IdEvent   int
	CreatedAt time.Time
	UpdatedAt time.Time
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
}

type Business interface {
	AddParticipant(data Core) error
	GetAllEventbyParticipant(idUser int) (data []Event, err error)
}

type Data interface {
	Add(data Core) error
	SelectDataEvent(idUser int) (data []Event, err error)
}
