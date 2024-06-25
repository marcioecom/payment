package processtx

import "github.com/marcioecom/payment/domain/entity"

type TransactionDtoInput struct {
	ID                        string  `json:"id"`
	AccountID                 string  `json:"accountId"`
	CreditCardNumber          string  `json:"creditCardNumber"`
	CreditCardName            string  `json:"creditCardName"`
	CreditCardExpirationMonth int     `json:"creditCardExpirationMonth"`
	CreditCardExpirationYear  int     `json:"creditCardExpirationYear"`
	CreditCardCVV             int     `json:"creditCardCvv"`
	Amount                    float64 `json:"amount"`
}

type TransactionDtoOutput struct {
	ID           string                   `json:"id"`
	Status       entity.TransactionStatus `json:"status"`
	ErrorMessage string                   `json:"errorMessage"`
}
