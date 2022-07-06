package data

import (
	"errors"
	"project3/eventapp/features/participants"

	"gorm.io/gorm"
)

type mysqlParticipantRepository struct {
	db *gorm.DB
}

// SelectDataEvent implements participants.Data
func (repo *mysqlParticipantRepository) SelectDataEvent(idUser int) (data []participants.Event, err error) {
	var dataParticipant []Event

	result := repo.db.Joins("Participant", repo.db.Where("user_id = ?", idUser)).Find(&dataParticipant)
	if result.Error != nil {
		return []participants.Event{}, result.Error
	}

	return ToCoreList(dataParticipant), result.Error
}

// Add implements participants.Data
func (repo *mysqlParticipantRepository) Add(ParticipantData participants.Core) error {
	Model := fromCore(ParticipantData)
	result := repo.db.Create(&Model)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed insert data")
	}
	return nil
}

func NewParticipantRepository(conn *gorm.DB) participants.Data {
	return &mysqlParticipantRepository{
		db: conn,
	}
}
