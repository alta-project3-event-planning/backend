package response

import (
	"project3/eventapp/features/comments"
	"time"
)

type Comment struct {
	ID        int       `json:"id" form:"id"`
	UserID    int       `json:"user_id" form:"user_id"`
	Name      string    `json:"name" form:"name"`
	Comment   string    `json:"comment" form:"comment"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
	User      []User
}

type Event struct {
	ID         int       `json:"id_event" form:"id_event"`
	Name       string    `json:"name" form:"name"`
	URL        string    `json:"image_url" form:"imageurl"`
	HostedBy   string    `json:"hostedby" form:"hostedby"`
	Performers string    `json:"performers" form:"performers"`
	Date       time.Time `json:"date" form:"date"`
	City       string    `json:"city" form:"city"`
	Location   string    `json:"location" form:"location"`
	Detail     string    `json:"details" form:"details"`
}

type User struct {
	ID        int       `json:"id" form:"id"`
	URL       string    `json:"url" form:"url"`
	Name      string    `json:"name" form:"name"`
	Email     string    `json:"email" form:"email"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
}

func FromCore(data comments.Core) Comment {
	return Comment{
		ID:      data.ID,
		Name:    data.User.Name,
		Comment: data.Comment,
	}
}

func FromCoreList(data []comments.Core) []Comment {
	result := []Comment{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
