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
  result := repo.db.Create(&data)
  if result.Error != nil {
    return 0, result.Error
  }
  return int(result.RowsAffected), nil
}