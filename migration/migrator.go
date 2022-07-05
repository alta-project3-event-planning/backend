package migration

import (
	_mEvent "project3/eventapp/features/events/data"
	_mUser "project3/eventapp/features/users/data"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(_mUser.User{})
	db.AutoMigrate(_mEvent.Event{})
}
