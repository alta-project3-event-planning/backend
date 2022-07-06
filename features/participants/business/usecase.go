package business

import (
	"errors"
	"project3/eventapp/features/participants"
)

type participantUseCase struct {
	participantData participants.Data
}

// GetAllEventbyParticipant implements participants.Business
func (uc *participantUseCase) GetAllEventbyParticipant(idUser int) (data []participants.Event, err error) {
	resp, err := uc.participantData.SelectDataEvent(idUser)
	return resp, err
}

// AddParticipant implements participants.Business
func (uc *participantUseCase) AddParticipant(partRequest participants.Core) error {
	if partRequest.IdEvent == 0 {
		return errors.New("data must be filled")
	}

	err := uc.participantData.Add(partRequest)
	return err
}

func NewParticipantBusiness(ptrData participants.Data) participants.Business {
	return &participantUseCase{
		participantData: ptrData,
	}
}
