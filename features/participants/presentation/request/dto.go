package request

import "project3/eventapp/features/participants"

type Participant struct {
	IdUser  int `json:"user_id" form:"user_id"`
	IdEvent int `json:"event_id" form:"event_id"`
}

func ToCore(partReq Participant) participants.Core {
	return participants.Core{
		IdUser:  partReq.IdUser,
		IdEvent: partReq.IdEvent,
	}
}
