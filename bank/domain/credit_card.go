package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)


type CreditCard struct {
	ID string
	Name string
	Number string
	Balance float64
	Limit float64
	ExpirationMonth uint32
	ExpirationYear uint32
	Cvv int32
	CreatedAt time.Time
}



func NewCreditCard() *CreditCard {
	c := &CreditCard{}
	c.ID = uuid.NewV4().String()
	c.CreatedAt = time.Now()

	return c
}