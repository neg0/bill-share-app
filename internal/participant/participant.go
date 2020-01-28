package participant

import (
	"bsocial/pkg/money"

	"math"
)

// Participant object
type Participant struct {
	name        string
	amountPaid  money.API
	amountSpent money.API
}

// ParticipantDescriber is a definition of Participant
type ParticipantDescriber interface {
	Name() string
	AmountPaid() money.API
	AmountSpent() money.API
	SetName(string)
	SetAmountPaid(float64)
	SetAmountSpent(float64)
	Owes(Collection) Collection
	Owned() money.API
	IsPayer() bool
}

// NewParticipant creates a new instance of Participant pointer
func NewParticipant() Participant {
	var instance Participant
	return instance
}

// Name returns a name of the participant
func (p *Participant) Name() string {
	return p.name
}

// SetName sets a name name for the participant
func (p *Participant) SetName(name string) {
	p.name = name
}

// SetAmountPaid will sets a new amountPaid for the participant
func (p *Participant) SetAmountPaid(amount float64) {
	p.amountPaid = money.NewMoney(amount)
}

// SetAmountSpent will sets a new amountSpent for the participant
func (p *Participant) SetAmountSpent(amount float64) {
	p.amountSpent = money.NewMoney(amount)
}

// AmountSpent returns amountPaid in Money ValueObject
func (p *Participant) AmountPaid() money.API {
	return p.amountPaid
}

// AmountSpent returns amountSpent in Money ValueObject
func (p *Participant) AmountSpent() money.API {
	return p.amountSpent
}

// Owes returns a list of participants that this participants owes money to
func (p *Participant) Owes(participants Collection) Collection {
	var owes Collection
	for _, participant := range participants {
		if p.Name() != participant.Name() &&
			participant.IsPayer() &&
			p.AmountPaid().Value() < participant.AmountPaid().Value() {
			owes = append(owes, participant)
		}
	}
	return owes
}

// Owned Money by the participant
func (p *Participant) Owned() money.API {
	return money.NewMoney(math.Abs(p.AmountPaid().Value() - p.AmountSpent().Value()))
}

// IsPayer indicates the type of participant in the collection if is a payer or a spender
func (p *Participant) IsPayer() bool {
	return p.AmountPaid().Value() > 0
}
