package business

import (
	"fmt"
	"project3/eventapp/features/participants"
	"testing"

	"github.com/stretchr/testify/assert"
)

//	mock data success case
type mockParticipantDataSuccess struct{}

func (mock mockParticipantDataSuccess) SelectDataEvent(idUser int) (data []participants.Core, err error) {
	return []participants.Core{
		{
			ID: 1,
			Event: participants.Event{
				Url:         "example.com",
				Name:        "user 1",
				HostedBy:    "andi",
				Performers:  "av7 band",
				City:        "surabaya",
				Location:    "1.2, 1.3",
				EventDetail: "detail av7 band",
			},
		},
	}, nil
}

func (mock mockParticipantDataSuccess) AddData(data participants.Core) (err error) {
	return nil
}

func (mock mockParticipantDataSuccess) DeleteData(param, userID int) (err error) {
	return nil
}

//	mock data failed case
type mockParticipantDataFailed struct{}

func (mock mockParticipantDataFailed) SelectDataEvent(idUser int) (data []participants.Core, err error) {
	return data, fmt.Errorf("Failed to select data event")
}

func (mock mockParticipantDataFailed) AddData(data participants.Core) (err error) {
	return fmt.Errorf("Failed to insert data event")
}

func (mock mockParticipantDataFailed) DeleteData(param, userID int) (err error) {
	return fmt.Errorf("Failed to delete data event")
}

func TestGetAllEventParticipant(t *testing.T) {
	t.Run("Test Get Event Data by Participant Success", func(t *testing.T) {
		id := 1
		participantBusiness := NewParticipantBusiness(mockParticipantDataSuccess{})
		result, err := participantBusiness.GetAllEventbyParticipant(id)
		assert.Nil(t, err)
		assert.Equal(t, "user 1", result[0].Event.Name)
	})

	t.Run("Test Get Event Data by Participant Failed", func(t *testing.T) {
		id := 1
		participantBusiness := NewParticipantBusiness(mockParticipantDataFailed{})
		result, err := participantBusiness.GetAllEventbyParticipant(id)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestAddParticipant(t *testing.T) {
	t.Run("Test Add Participant Success", func(t *testing.T) {
		participantBusiness := NewParticipantBusiness(mockParticipantDataSuccess{})
		newParticipant := participants.Core{
			EventID: 1,
		}
		err := participantBusiness.AddParticipant(newParticipant)
		assert.Nil(t, err)
	})

	t.Run("Test Add Participant Failed", func(t *testing.T) {
		participantBusiness := NewParticipantBusiness(mockParticipantDataFailed{})
		newParticipant := participants.Core{
			EventID: -1,
		}
		err := participantBusiness.AddParticipant(newParticipant)
		assert.NotNil(t, err)
	})
}

func TestDeleteData(t *testing.T) {
	t.Run("Test Delete Data Event Success", func(t *testing.T) {
		id := 1
		userId := 1
		participantBusiness := NewParticipantBusiness(mockParticipantDataSuccess{})
		err := participantBusiness.DeleteParticipant(id, userId)
		assert.Nil(t, err)
	})

	t.Run("Test Delete Data Event Failed", func(t *testing.T) {
		id := -1
		userId := -1
		participantBusiness := NewParticipantBusiness(mockParticipantDataFailed{})
		err := participantBusiness.DeleteParticipant(id, userId)
		assert.NotNil(t, err)
	})
}