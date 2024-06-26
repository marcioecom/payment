package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransaction_Validate(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 900

	assert.Nil(t, transaction.Validate())
}

func TestTransaction_IsNotValidWithAmountGreaterThan1000(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 1001
	err := transaction.Validate()
	assert.Error(t, err)
	assert.Equal(t, "you don't have limit for this transaction", err.Error())
}

func TestTransaction_IsNotValidWithAmountLessThan1(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 0
	err := transaction.Validate()
	assert.Error(t, err)
	assert.Equal(t, "the amount must be greater than 1", err.Error())
}
