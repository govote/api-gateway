package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/deputadosemfoco/api-gateway/config"
	"github.com/deputadosemfoco/api-gateway/models"
	"github.com/deputadosemfoco/go-libs/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/parnurzeal/gorequest"
)

type (
	// UserCtrl ...
	UserCtrl struct {
		Config config.Provider
	}

	// RegistrationRequest represents the required information to register or login a User
	RegistrationRequest struct {
		Name                string
		Email               string
		PhotoURL            string
		FacebookID          string
		FacebookAccessToken string
	}

	// User represents the application user
	User struct {
		domain.Entity

		ID         string `json:"id"`
		Name       string `json:"name"`
		Email      string `json:"email"`
		PhotoURL   string `json:"photoUrl"`
		FacebookID string `json:"facebookId"`
		GoogleID   string `json:"googleId"`
	}
)

// AuthAndRegister sign up and sign in user.
// If successful, returns a jwt token required for all application secure operations
func (ctrl *UserCtrl) AuthAndRegister(c echo.Context) error {
	endpoint := ctrl.Config.Get().UserEndpoint
	req := new(RegistrationRequest)

	if err := c.Bind(req); err != nil {
		return err
	}

	res, body, err := gorequest.New().
		Post(endpoint).
		SendStruct(req).
		End()

	if len(err) == 0 && (res.StatusCode == http.StatusCreated || res.StatusCode == http.StatusOK) {
		user := new(User)
		json.Unmarshal([]byte(body), user)

		token := ctrl.buildJwtToken(user.Name, user.FacebookID)
		js := &struct {
			Token string `json:"token"`
			User  *User  `json:"user"`
		}{token, user}

		return c.JSON(res.StatusCode, js)
	}

	return c.JSONBlob(res.StatusCode, []byte(body))
}

func (ctrl *UserCtrl) buildJwtToken(name, fbID string) string {
	secret := ctrl.Config.Get().TokenSecret

	claims := &models.Claims{
		StandardClaims: jwt.StandardClaims{
			Id:        fbID,
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
		Name:       name,
		FacebookID: fbID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))

	if err != nil {
		panic(err)
	}

	return t
}
