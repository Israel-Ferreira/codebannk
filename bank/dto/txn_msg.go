package dto

import "time"

type TxnMsg struct {
	ID string
	Name string
	Number string
	ExpirationMonth uint32
	ExpirationYear uint32
	CVV rune
	Amount float64
	Store string
	Description string
	CreatedAt time.Time
}


func ConvertTxnDtoToTxnMsg(txnDto TxnDto) TxnMsg {
	return TxnMsg{
		ID: txnDto.ID,
		Name: txnDto.Name,
		Number: txnDto.Number,
		ExpirationMonth: txnDto.ExpirationMonth,
		ExpirationYear: txnDto.ExpirationYear,
		CVV: txnDto.Cvv,
		Amount: txnDto.Amount,
		Store: txnDto.Store,
		Description: txnDto.Description,
		CreatedAt: txnDto.CreatedAt,
	}
}