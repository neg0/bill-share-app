package test

import (
	"bsocial/internal/participant"
	"bsocial/pkg/money"
	testMoney "bsocial/pkg/money/test"
)

// ParticipantMock is a Mock Object for testing purpose to avoid duplication in our test cases
type ParticipantMock struct {
	name        string
	paidAmount  money.API
	spentAmount money.API
}

// NewParticipantMock creates a new instance of ParticipantMock with three default parameters
func NewParticipantMock(name string, paid float64, spent float64) *ParticipantMock {
	return &ParticipantMock{
		name:        name,
		paidAmount:  testMoney.NewMoneyMock(paid),
		spentAmount: testMoney.NewMoneyMock(spent),
	}
}

// Name returns name of mocked participant
func (pm *ParticipantMock) Name() string {
	return pm.name
}

// SetName sets a new name for mocked participant
func (pm *ParticipantMock) SetName(name string) {
	pm.name = name
}

// AmountPaid returns amountPaid of mocked participant
func (pm *ParticipantMock) AmountPaid() money.API {
	return pm.paidAmount
}

// SetAmountPaid sets a new amountPaid for mocked participant
func (pm *ParticipantMock) SetAmountPaid(amount float64) {
	pm.paidAmount = testMoney.NewMoneyMock(amount)
}

// AmountSpent returns amountSpent of mocked participant
func (pm *ParticipantMock) AmountSpent() money.API {
	return pm.spentAmount
}

// SetAmountSpent sets a new amountSpent for mocked participant
func (pm *ParticipantMock) SetAmountSpent(amount float64) {
	pm.spentAmount = testMoney.NewMoneyMock(amount)
}

// Owes returns a list of participants the participant mocks owes
func (pm *ParticipantMock) Owes(collection participant.Collection) participant.Collection {
	return nil
}

// Owned is the amount owned by participant
func (pm *ParticipantMock) Owned() money.API {
	return pm.paidAmount
}

// IsPayer is an indication of participant to be payer or not
func (pm *ParticipantMock) IsPayer() bool {
	return pm.paidAmount.Value() > 0
}
