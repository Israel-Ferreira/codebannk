package dto

import "time"


type TxnDto struct {
	ID string
	Name string
	Number string
	ExpirationMonth	uint32
	ExpirationYear uint32
	Cvv int32
	Amount float64
	Store string
	Description string
	CreatedAt time.Time
}
