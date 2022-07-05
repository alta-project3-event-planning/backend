package data

import (
	"errors"
	"project3/eventapp/features/events"

	"gorm.io/gorm"
)

type mysqlEventRepository struct {
	db *gorm.DB
}

func NewEventRepository(conn *gorm.DB) events.Data {
	return &mysqlEventRepository{
		db: conn,
	}
}

func (repo *mysqlEventRepository) SelectData() (response []events.Core, err error) {
	var dataEvent []Event
	result := repo.db.Find(&dataEvent)
	if result.Error != nil {
		return []events.Core{}, result.Error
	}
	return ToCoreList(dataEvent), result.Error
}

func (repo *mysqlEventRepository) SelectDataByID(id int) (response events.Core, err error) {
	dataEvent := Event{}
	result := repo.db.Find(&dataEvent, id)
	if result.Error != nil {

		return events.Core{}, result.Error
	}

	return toCore(dataEvent), err
}

func (repo *mysqlEventRepository) InsertData(EventData events.Core) error {
	EventModel := fromCore(EventData)
	result := repo.db.Create(&EventModel)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed insert data")
	}
	return nil
}

func (repo *mysqlEventRepository) DeleteDataByID(id int, userId int) error {
	dataEvent := Event{}
	result := repo.db.Where("user_id = ?", userId).Delete(&dataEvent, id)
	if result.RowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return result.Error
}

func (repo *mysqlEventRepository) UpdateDataByID(dataReq map[string]interface{}, id int, userId int) error {

	model := Event{}
	model.ID = uint(id)
	result := repo.db.Model(model).Where("user_id = ?", userId).Updates(dataReq)
	if result.RowsAffected == 0 {
		return errors.New("no row affected")
	}
	if result != nil {
		return result.Error
	}

	return nil

}

func (repo *mysqlEventRepository) SelectDataByUserID(id_user int) (response []events.Core, err error) {
	var dataEvent []Event
	result := repo.db.Where("user_id = ?", id_user).Find(&dataEvent)
	if result.Error != nil {
		return []events.Core{}, result.Error
	}
	return ToCoreList(dataEvent), result.Error
}