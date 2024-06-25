package repository

import "github.com/marcioecom/payment/domain/entity"

type TransactionRepository interface {
	Insert(id string, account string, amount float64, status entity.TransactionStatus, errorMessage string) error
}
