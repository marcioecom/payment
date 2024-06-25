package transaction

import (
	"encoding/json"

	"github.com/marcioecom/payment/domain/entity"
	"github.com/marcioecom/payment/usecase/processtx"
)

type KafkaPresenter struct {
	ID           string                   `json:"id"`
	Status       entity.TransactionStatus `json:"status"`
	ErrorMessage string                   `json:"errorMessage"`
}

func NewKafkaPresenter() *KafkaPresenter {
	return &KafkaPresenter{}
}

func (k *KafkaPresenter) Show() ([]byte, error) {
	data, err := json.Marshal(k)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (k *KafkaPresenter) Bind(input any) error {
	k.ID = input.(processtx.TransactionDtoOutput).ID
	k.Status = input.(processtx.TransactionDtoOutput).Status
	k.ErrorMessage = input.(processtx.TransactionDtoOutput).ErrorMessage
	return nil
}
