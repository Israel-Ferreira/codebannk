package usecase

import (
	"github.com/Israel-Ferreira/codebank/domain"
	"github.com/Israel-Ferreira/codebank/dto"
)

type UseCaseTxn struct {
	TxnRepository domain.TransactionRepository
}

func NewUseCaseTxn(txnRepo domain.TransactionRepository) UseCaseTxn {
	return UseCaseTxn{TxnRepository: txnRepo}
}

func (u UseCaseTxn) ProcessTxn(txnDto dto.TxnDto) (domain.Transaction, error) {
	creditCard := u.hydrateCreditCard(txnDto)
	ccBalanceAndLimit, err := u.TxnRepository.GetCreditCard(*creditCard)

	if err != nil {
		return domain.Transaction{}, err
	}

	creditCard.ID = ccBalanceAndLimit.ID
	creditCard.Balance = ccBalanceAndLimit.Balance
	creditCard.Limit = ccBalanceAndLimit.Limit


	newTxn := u.hydrateTxn(txnDto, ccBalanceAndLimit)

	newTxn.ProcessAndValidate(creditCard)

	return *newTxn, nil
}

func (u UseCaseTxn) hydrateCreditCard(txnDto dto.TxnDto) *domain.CreditCard {
	creditCard := domain.NewCreditCard()

	creditCard.Name = txnDto.Name
	creditCard.Number = txnDto.Number

	creditCard.Cvv = txnDto.Cvv

	creditCard.ExpirationMonth = txnDto.ExpirationMonth
	creditCard.ExpirationYear = txnDto.ExpirationYear

	return creditCard
}

func (u UseCaseTxn) hydrateTxn(txnDto dto.TxnDto, creditCard domain.CreditCard) *domain.Transaction {
	txn := domain.NewTransaction()

	txn.Amount = txnDto.Amount
	txn.Store = txnDto.Store
	txn.CreditCardId = creditCard.ID
	txn.Description = txnDto.Description

	return txn
}
