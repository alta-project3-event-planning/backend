package request

import "project3/eventapp/features/participants"

type Participant struct {
	UserID  int `json:"user_id" form:"user_id"`
	EventID int `json:"event_id" form:"event_id"`
}

func ToCore(partReq Participant) participants.Core {
	return participants.Core{
		UserID:  partReq.UserID,
		EventID: partReq.EventID,
	}
}
