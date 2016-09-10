package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

type HealthCtrl struct{}

func NewHealthCtrl() *HealthCtrl {
	return &HealthCtrl{}
}

func (ctrl *HealthCtrl) Check(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
