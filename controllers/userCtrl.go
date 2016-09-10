package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/parnurzeal/gorequest"
	"github.com/vitorsalgado/la-democracia/gateway/endpoints"
	"github.com/vitorsalgado/la-democracia/gateway/models"
	"github.com/vitorsalgado/la-democracia/lib/go/messages"
)

type UserCtrl struct{}

func NewUserCtrl() *UserCtrl {
	return &UserCtrl{}
}

func (ctrl *UserCtrl) Register(c echo.Context) error {
	endpoint := endpoints.UserEndpoint
	req := new(models.RegistrationRequest)
	res := messages.Response{}

	if err := c.Bind(req); err != nil {
		return err
	}

	request := gorequest.New()

	resp, _, errs := request.Post(endpoint).
		SendStruct(req).
		EndStruct(&res)

	if errs != nil {
		return c.JSON(resp.StatusCode, res)
	}

	return c.JSON(http.StatusOK, res)
}
