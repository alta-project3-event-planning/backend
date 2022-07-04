package migration

import (
	_mProduct "project3/eventapp/features/products/data"
	_mUser "project3/eventapp/features/users/data"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(_mUser.User{})
	db.AutoMigrate(_mProduct.Product{})
}
