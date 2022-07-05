package data

import (
	"project3/eventapp/features/participants"
	"time"

	"gorm.io/gorm"
)

type Participant struct {
	gorm.Model
	IdUser  int
	IdEvent int
	Event   Event
	User    User
}

type Event struct {
	gorm.Model
	Name        string    `json:"name" form:"name"`
	Detail      string    `json:"detail" form:"detail"`
	URL         string    `json:"url" form:"url"`
	Date        time.Time `json:"time" form:"time"`
	Performers  string    `json:"performers" form:"performers"`
	HostedBy    string    `json:"hostedby" form:"hostedby"`
	City        string    `json:"city" form:"city"`
	Location    string    `json:"location" form:"location"`
	IDUser      int
	User        User
	Participant []Participant
}

type User struct {
	gorm.Model
	Name        string
	Email       string
	Password    string
	Event       []Event
	Participant []Participant
}

func (data *Participant) toCore() participants.Core {
	return participants.Core{
		ID:        int(data.ID),
		IdEvent:   data.IdEvent,
		IdUser:    data.IdUser,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func ToCoreList(data []Participant) []participants.Core {
	result := []participants.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}
