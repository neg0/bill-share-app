package receipt

import "bsocial/internal/participant"

// Participants will return collection of participants
func (c *Collection) Participants() participant.Collection {
	participants := participant.NewCollection()
	participants = participants.From(c.ParticipantList())
	for _, individualBill := range *c {
		for _, individual := range participants {
			if individualBill.PayerName() == individual.Name() {
				individual.SetAmountPaid(individualBill.TotalAmount())
			}
		}

		for name, amount := range individualBill.SpendersCollection() {
			for _, individualParticipant := range participants {
				if name == individualParticipant.Name() {
					individualParticipant.SetAmountSpent(individualParticipant.AmountSpent().Value() + amount)
					if !individualParticipant.IsPayer() {
						individualParticipant.SetAmountPaid(0)
					}
				}
			}
		}
	}
	return participants
}
