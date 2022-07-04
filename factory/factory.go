package factory

import (
	_userBusiness "project3/eventapp/features/users/business"
	_userData "project3/eventapp/features/users/data"
	_userPresentation "project3/eventapp/features/users/presentation"

	_authBusiness "project3/eventapp/features/auth/business"
	_authData "project3/eventapp/features/auth/data"
	_authPresentation "project3/eventapp/features/auth/presentation"

	_productBusiness "project3/eventapp/features/products/bussiness"
	_productData "project3/eventapp/features/products/data"
	_productPresentation "project3/eventapp/features/products/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	UserPresenter    *_userPresentation.UserHandler
	AuthPresenter    *_authPresentation.AuthHandler
	ProductPresenter *_productPresentation.ProductHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {

	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	userPresentation := _userPresentation.NewUserHandler(userBusiness)

	authData := _authData.NewAuthRepository(dbConn)
	authBusiness := _authBusiness.NewAuthBusiness(authData)
	authPresentation := _authPresentation.NewAuthHandler(authBusiness)

	productData := _productData.NewProductRepository(dbConn)
	productBusiness := _productBusiness.NewProductBusiness(productData)
	productPresentation := _productPresentation.NewProductHandler(productBusiness)

	return Presenter{
		UserPresenter:    userPresentation,
		AuthPresenter:    authPresentation,
		ProductPresenter: productPresentation,
	}
}
