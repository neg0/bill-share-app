package test

import (
	"bsocial/internal/participant"
	participantTest "bsocial/internal/participant/test"
	"bsocial/pkg/money"

	"fmt"
)

type PayBackMock struct {
	payer  participant.ParticipantDescriber
	payee  participant.ParticipantDescriber
	amount money.API
}

func NewPayBackMock(payerName string, payeeName string, amount float64) *PayBackMock {
	payer := participantTest.NewParticipantMock(payerName, 0, 0)
	payee := participantTest.NewParticipantMock(payeeName, 0, 0)

	return &PayBackMock{payer, payee, money.NewMoney(amount)}
}

func (pbm *PayBackMock) Payee() participant.ParticipantDescriber {
	return pbm.payee
}

func (pbm *PayBackMock) Payer() participant.ParticipantDescriber {
	return pbm.payer
}

func (pbm *PayBackMock) Amount() money.API {
	return pbm.amount
}

func (pbm *PayBackMock) String() string {
	return fmt.Sprintf("%s pays %s %s", pbm.payer.Name(), pbm.payee.Name(), pbm.Amount().String())
}
