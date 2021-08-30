package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Transaction struct {
	ID           string
	Amount       float64
	Status       string
	CreditCardId string
	Store        string
	CreatedAt    time.Time
}

func (txn *Transaction) ProcessAndValidate(creditCard *CreditCard) {
	if (creditCard.Balance + txn.Amount) > creditCard.Limit {
		txn.Status = "REJECTED"
	} else {
		txn.Status = "APPROVED"
		creditCard.Balance += txn.Amount
	}
}

func NewTransaction() *Transaction {
	txn := &Transaction{}
	txn.ID = uuid.NewV4().String()
	txn.CreatedAt = time.Now()

	return txn
}
