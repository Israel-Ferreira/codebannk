package repository

import (
	"database/sql"
	"errors"

	"github.com/Israel-Ferreira/codebank/domain"
)

type TxnRepository struct {
	db *sql.DB
}

func NewTxnRepository(db *sql.DB) *TxnRepository {
	return &TxnRepository{db: db}
}

func (txnRepo TxnRepository) SaveTransaction(txn domain.Transaction, creditCard domain.CreditCard) error {
	stmt, err := txnRepo.db.Prepare(
		`insert into transactions(id, credit_card_id, amount, status, description, store, created_at) values ($1, $2, $3, $4, $5, $6, $7)`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		txn.ID,
		txn.CreditCardId,
		txn.Amount,
		txn.Status,
		txn.Description,
		txn.Store,
		txn.CreatedAt,
	)

	if err != nil {
		return err
	}

	if txn.Status == "APPROVED" {
		if err = txnRepo.updateBalance(creditCard); err != nil {
			return err
		}
	}

	if err = stmt.Close(); err != nil {
		return err
	}

	return nil
}

func (txnRepo TxnRepository) updateBalance(creditCard domain.CreditCard) error {
	_, err := txnRepo.db.Exec(
		"update credit_cards set balance = $1 where id = $2",
		creditCard.Balance,
		creditCard.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (txnRepo TxnRepository) CreateCreditCard(creditCard domain.CreditCard) error {
	stmt, err := txnRepo.db.Prepare(
		`insert into credit_cards(id, name, number, expiration_month,expiration_year, CVV,balance, balance_limit)  values($1,$2,$3,$4,$5,$6,$7,$8)`,
	)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		creditCard.ID,
		creditCard.Name,
		creditCard.Number,
		creditCard.ExpirationMonth,
		creditCard.ExpirationYear,
		creditCard.Cvv,
		creditCard.Balance,
		creditCard.Limit,
	)

	if err != nil {
		return err
	}

	return nil

}

func (txnRepo TxnRepository) GetCreditCard(creditCard domain.CreditCard) (domain.CreditCard, error) {
	var c domain.CreditCard

	stmt, err := txnRepo.db.Prepare(
		`select id, balance, balance_limit from credit_cards where number = $1`,
	)

	if err != nil {
		return c, err
	}

	if err = stmt.QueryRow(creditCard.Number).Scan(&c.ID, &c.Balance, &c.Limit); err != nil {
		return c, errors.New("credit card does not exists")
	}

	return c, nil
}
