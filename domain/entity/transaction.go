package entity

import "errors"

type TransactionStatus string

const (
	REJECTED TransactionStatus = "rejected"
	APPROVED TransactionStatus = "approved"
)

type Transaction struct {
	ID           string
	AccountID    string
	Amount       float64
	CreditCard   *CreditCard
	Status       TransactionStatus
	ErrorMessage string
}

func NewTransaction() *Transaction {
	return &Transaction{}
}

func (t *Transaction) Validate() error {
	if t.Amount > 1000 {
		return errors.New("you don't have limit for this transaction")
	}

	if t.Amount < 1 {
		return errors.New("the amount must be greater than 1")
	}

	return nil
}

func (t *Transaction) SetCreditCard(card *CreditCard) {
	t.CreditCard = card
}
