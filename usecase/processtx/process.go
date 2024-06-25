package processtx

import (
	"github.com/marcioecom/payment/adapter/broker"
	"github.com/marcioecom/payment/domain/entity"
	"github.com/marcioecom/payment/domain/repository"
)

type ProcessTransaction struct {
	repository repository.TransactionRepository
	producer   broker.ProducerInterface
	topic      string
}

func NewProcessTransaction(repository repository.TransactionRepository, producer broker.ProducerInterface, topic string) *ProcessTransaction {
	return &ProcessTransaction{
		repository: repository,
		producer:   producer,
		topic:      topic,
	}
}

func (p *ProcessTransaction) Execute(input TransactionDtoInput) (TransactionDtoOutput, error) {
	transaction := entity.NewTransaction()
	transaction.ID = input.ID
	transaction.AccountID = input.AccountID
	transaction.Amount = input.Amount

	cc, invalidCC := entity.NewCreditCard(
		input.CreditCardNumber,
		input.CreditCardName,
		input.CreditCardExpirationMonth,
		input.CreditCardExpirationYear,
		input.CreditCardCVV,
	)
	if invalidCC != nil {
		return p.rejectTransaction(transaction, invalidCC)
	}

	transaction.SetCreditCard(cc)
	if invalidTx := transaction.Validate(); invalidTx != nil {
		return p.rejectTransaction(transaction, invalidTx)
	}

	if err := p.repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, entity.APPROVED, ""); err != nil {
		return TransactionDtoOutput{}, err
	}

	output := TransactionDtoOutput{
		ID:           transaction.ID,
		Status:       entity.APPROVED,
		ErrorMessage: "",
	}

	if err := p.publish(output, []byte(transaction.ID)); err != nil {
		return TransactionDtoOutput{}, err
	}

	return output, nil
}

func (p *ProcessTransaction) rejectTransaction(transaction *entity.Transaction, invalidTx error) (TransactionDtoOutput, error) {
	if err := p.repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, entity.REJECTED, invalidTx.Error()); err != nil {
		return TransactionDtoOutput{}, err
	}

	output := TransactionDtoOutput{
		ID:           transaction.ID,
		Status:       entity.REJECTED,
		ErrorMessage: invalidTx.Error(),
	}

	if err := p.publish(output, []byte(transaction.ID)); err != nil {
		return TransactionDtoOutput{}, err
	}

	return output, nil
}

func (p *ProcessTransaction) publish(output TransactionDtoOutput, key []byte) error {
	err := p.producer.Publish(output, key, p.topic)
	if err != nil {
		return err
	}
	return nil
}
