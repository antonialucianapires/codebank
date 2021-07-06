package domain

import (
	"github.com/antonialucianapires/codebank/enum"
	_ "github.com/antonialucianapires/codebank/enum"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Transaction struct {
	ID string
	Anount float64
	Status string
	Description string
	Store string
	CreditCardId string
	CreatedAt time.Time
}

func NewTransaction() *Transaction {
	t := &Transaction{}
	t.ID = uuid.NewV4().String()
	t.CreatedAt = time.Now()
	return t
}

func(t *Transaction) ProcessAndValidate(creditCard *CreditCard) {

	//RN =
	if t.Anount + creditCard.Balance > creditCard.Limit {
		t.Status = enum.StatusTransaction(1).String()
	} else {
		t.Status = enum.StatusTransaction(2).String()
		creditCard.Balance = creditCard.Balance + t.Anount
	}
}