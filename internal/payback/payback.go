package payback

import (
	"bsocial/internal/participant"
	"bsocial/pkg/money"

	"fmt"
)

// Builder is definition of Payback
type Reader interface {
	Payee() participant.Builder
	Payer() participant.Builder
	Amount() money.Reader
	String() string
}

// Payback is an entity in which holds a single transaction payback
type Payback struct {
	payer  participant.Builder
	payee  participant.Builder
	amount money.Reader
}

// NewPayBack creates a new instance of PayBack
func NewPayBack(payer participant.Builder, payee participant.Builder) Payback {
	return Payback{
		payer:  payer,
		payee:  payee,
		amount: money.NewMoney(payer.AmountPaid().Value()),
	}
}

// Payee will return the payee during instantiation
func (p Payback) Payee() participant.Builder {
	return p.payee
}

// Payer will return the payer during instantiation
func (p Payback) Payer() participant.Builder {
	return p.payer
}

// Amount returns the amount of payback to payee
func (p Payback) Amount() money.Reader {
	return p.amount
}

// String returns English formatted string of for process of payback (from who to whom and the amount that should be paid)
func (p Payback) String() string {
	return fmt.Sprintf("%s pays %s %s", p.payer.Name(), p.payee.Name(), p.Amount().String())
}
