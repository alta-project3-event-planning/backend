package factory

import (
	_userBusiness "project3/eventapp/features/users/business"
	_userData "project3/eventapp/features/users/data"
	_userPresentation "project3/eventapp/features/users/presentation"

	_authBusiness "project3/eventapp/features/auth/business"
	_authData "project3/eventapp/features/auth/data"
	_authPresentation "project3/eventapp/features/auth/presentation"

	_eventBusiness "project3/eventapp/features/events/bussiness"
	_eventData "project3/eventapp/features/events/data"
	_eventPresentation "project3/eventapp/features/events/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	UserPresenter    *_userPresentation.UserHandler
	AuthPresenter    *_authPresentation.AuthHandler
	EventPresenter *_eventPresentation.EventHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {

	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	userPresentation := _userPresentation.NewUserHandler(userBusiness)

	authData := _authData.NewAuthRepository(dbConn)
	authBusiness := _authBusiness.NewAuthBusiness(authData)
	authPresentation := _authPresentation.NewAuthHandler(authBusiness)

	eventData := _eventData.NewEventRepository(dbConn)
	eventBusiness := _eventBusiness.NewEventBusiness(eventData)
	eventPresentation := _eventPresentation.NewEventHandler(eventBusiness)

	return Presenter{
		UserPresenter:    userPresentation,
		AuthPresenter:    authPresentation,
		EventPresenter: eventPresentation,
	}
}
