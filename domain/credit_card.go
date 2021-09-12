package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type CreditCard struct {
	ID               string
	Name             string
	Number           string
	ExpirantionMonth int32
	ExpirantionYear  int32
	CVV              int32
	Balance          float64
	Limit            float64
	CreatedAt        time.Time
}

func NewCreditCard() CreditCard {
	c := CreditCard{}

	c.ID = uuid.NewV4().String()
	c.CreatedAt = time.Now()

	return c
}
