package data

import (
	"project3/eventapp/features/comments"

	"gorm.io/gorm"
)

type mysqlCommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(conn *gorm.DB) comments.Data {
	return &mysqlCommentRepository{
		db: conn,
	}
}

func (repo *mysqlCommentRepository) Add(data comments.Core) (row int, err error) {
	commentData := fromCore(data)
	result := repo.db.Create(&commentData)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}

func (repo *mysqlCommentRepository) GetComment(limit, offset, eventId int) (response []comments.Core, err error) {
	var dataComment []Comment

	result := repo.db.Where("event_id = ?", eventId).Preload("User").Preload("Event").Limit(limit).Offset(offset).Find(&dataComment).Order("id DESC")

	if result.Error != nil {
		return []comments.Core{}, result.Error
	}
	return ToCoreList(dataComment), result.Error
}
