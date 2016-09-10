package models

import "github.com/vitorsalgado/la-democracia/lib/go/messages"

type RegistrationRequest struct {
	messages.Request
	Name       string
	Email      string
	PhotoURL   string
	FacebookID string
	GoogleID   string
}
