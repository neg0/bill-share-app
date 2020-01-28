package test

import (
	"bsocial/internal/participant"
	participantTest "bsocial/internal/participant/test"
	"bsocial/pkg/money"

	"fmt"
)

type PayBackMock struct {
	payer  participant.Builder
	payee  participant.Builder
	amount money.Reader
}

func NewPayBackMock(payerName string, payeeName string, amount float64) *PayBackMock {
	payer := participantTest.NewParticipantMock(payerName, 0, 0)
	payee := participantTest.NewParticipantMock(payeeName, 0, 0)

	return &PayBackMock{payer, payee, money.NewMoney(amount)}
}

func (pbm *PayBackMock) Payee() participant.Builder {
	return pbm.payee
}

func (pbm *PayBackMock) Payer() participant.Builder {
	return pbm.payer
}

func (pbm *PayBackMock) Amount() money.Reader {
	return pbm.amount
}

func (pbm *PayBackMock) String() string {
	return fmt.Sprintf("%s pays %s %s", pbm.payer.Name(), pbm.payee.Name(), pbm.Amount().String())
}
