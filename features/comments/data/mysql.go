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

func (repo *mysqlCommentRepository) GetComment(offset, eventId int) (response []comments.Core, err error) {
	var dataComment []Comment
	result := repo.db.Order("id DESC").Where("event_id = ?", eventId).Preload("User").Limit(5).Offset(offset).Find(&dataComment)

	if result.Error != nil {
		return []comments.Core{}, result.Error
	}
	return ToCoreList(dataComment), result.Error
}
