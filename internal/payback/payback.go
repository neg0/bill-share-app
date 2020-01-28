package payback

import (
	"bsocial/internal/participant"
	"bsocial/pkg/money"

	"fmt"
)

// ParticipantDescriber is definition of Payback
type API interface {
	Payee() participant.ParticipantDescriber
	Payer() participant.ParticipantDescriber
	Amount() money.API
	String() string
}

// Payback is an entity in which holds a single transaction payback
type Payback struct {
	payer  participant.ParticipantDescriber
	payee  participant.ParticipantDescriber
	amount money.API
}

// NewPayBack creates a new instance of PayBack
func NewPayBack(payer participant.ParticipantDescriber, payee participant.ParticipantDescriber) Payback {
	return Payback{
		payer:  payer,
		payee:  payee,
		amount: money.NewMoney(payer.AmountPaid().Value()),
	}
}

// Payee will return the payee during instantiation
func (p Payback) Payee() participant.ParticipantDescriber {
	return p.payee
}

// Payer will return the payer during instantiation
func (p Payback) Payer() participant.ParticipantDescriber {
	return p.payer
}

// Amount returns the amount of payback to payee
func (p Payback) Amount() money.API {
	return p.amount
}

// String returns English formatted string of for process of payback (from who to whom and the amount that should be paid)
func (p Payback) String() string {
	return fmt.Sprintf("%s pays %s %s", p.payer.Name(), p.payee.Name(), p.Amount().String())
}
