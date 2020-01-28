package receipt

import (
	"bsocial/internal/payback"

	"math"
)

// Split is service to return a collection of payback entity
// Number of entities in collection defines the minimum number of transactions to payback the bill to the payer of bill
// Each payback and its order defines the process of payback and its sequential order
func (c *Collection) Split() payback.Collection {
	participants := c.Participants()
	payers := participants.Payers()
	spenders := participants.Spenders()

	payBackCollection := payback.NewCollection()
	settledCounter := 0
	for {
		if settledCounter == len(payers) {
			break
		}

		for _, payer := range payers {
			if payer.AmountPaid().Value() == 0 {
				settledCounter++
				continue
			}

			for _, spender := range spenders {
				if spender.AmountSpent() == spender.AmountPaid() {
					continue
				}

				newAmount := payer.AmountPaid().Value() - spender.AmountSpent().Value()
				payer.SetAmountPaid(newAmount)
				spender.SetAmountPaid(spender.AmountSpent().Value())
				if newAmount < 0 {
					payer.SetAmountPaid(0)
					spender.SetAmountPaid(spender.AmountSpent().Value() - math.Abs(newAmount))
				}

				payBackCollection = payBackCollection.Add(payback.NewPayBack(spender, payer))
				break
			}
		}
	}

	return payBackCollection
}
