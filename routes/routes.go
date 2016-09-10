package routes

import (
	"github.com/labstack/echo"
	"github.com/vitorsalgado/la-democracia/gateway/controllers"
)

type (
	Router struct{}
)

func (router *Router) SetUp() *echo.Echo {
	e := echo.New()
	userCtrl := controllers.NewUserCtrl()
	healthCtrl := controllers.NewHealthCtrl()

	e.SetHTTPErrorHandler(controllers.GtHTTPErrorHandler)

	e.POST("/api/user", userCtrl.Register)

	e.Get("/api/chk", healthCtrl.Check)

	return e
}
