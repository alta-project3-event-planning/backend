package routes

import (
	"project3/eventapp/factory"
	"project3/eventapp/middlewares"

	"github.com/labstack/echo/v4"
)

func New(presenter factory.Presenter) *echo.Echo {

	e := echo.New()
	e.Pre(middlewares.RemoveTrailingSlash())

	e.Use(middlewares.CorsMiddleware())

	e.POST("/register", presenter.UserPresenter.Insert)
	e.POST("/login", presenter.AuthPresenter.Login)

	e.GET("/users", presenter.UserPresenter.GetAll, middlewares.JWTMiddleware())
	e.GET("/users/details", presenter.UserPresenter.GetDataById, middlewares.JWTMiddleware())
	e.PUT("/users", presenter.UserPresenter.Update, middlewares.JWTMiddleware())
	e.DELETE("/users", presenter.UserPresenter.Delete, middlewares.JWTMiddleware())

	e.GET("/products", presenter.ProductPresenter.GetAll)
	e.GET("/products/:id", presenter.ProductPresenter.GetDataById)
	e.POST("/products", presenter.ProductPresenter.InsertData, middlewares.JWTMiddleware())
	e.PUT("/products/:id", presenter.ProductPresenter.UpdateData, middlewares.JWTMiddleware())
	e.DELETE("/products/:id", presenter.ProductPresenter.DeleteData, middlewares.JWTMiddleware())

	e.GET("/user-products", presenter.ProductPresenter.GetProductByUser, middlewares.JWTMiddleware())

	return e
}
