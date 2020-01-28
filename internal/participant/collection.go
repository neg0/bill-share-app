package participant

import (
	"sort"
)

// Collection is an array of Participant
// Implements Sort interface for sorting participants by amount paid and spent
type Collection []ParticipantDescriber

// NewCollection creates a new instance of Collection
func NewCollection() Collection {
	var instance Collection
	return instance
}

// From will sort the map of string and float by float value
func (c *Collection) From(participants []string) Collection {
	for _, participantName := range participants {
		participant := &Participant{}
		participant.SetName(participantName)
		participant.SetAmountPaid(0)
		participant.SetAmountSpent(0)

		*c = append(*c, participant)
	}
	return *c
}

// Payers will list the payers among participants
func (c Collection) Payers() Collection {
	if len(c) < 1 {
		return c
	}

	var payerCollection Collection
	for _, individual := range c {
		if individual.IsPayer() {
			newBalance := individual.AmountPaid().Value() - individual.AmountSpent().Value()
			individual.SetAmountPaid(newBalance)
			payerCollection = append(payerCollection, individual)
		}
	}

	sort.Sort(payerCollection)
	return payerCollection
}

// Spenders will list spenders among participants
func (c Collection) Spenders() Collection {
	if len(c) < 1 {
		return c
	}

	var spenderCollection Collection
	for _, individual := range c {
		if !individual.IsPayer() {
			spenderCollection = append(spenderCollection, individual)
		}
	}

	sort.Sort(spenderCollection)
	return spenderCollection
}

func (c Collection) Len() int {
	return len(c)
}

func (c Collection) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c Collection) Less(i, j int) bool {
	if !c[i].IsPayer() && c[i].AmountSpent().Value() > c[j].AmountSpent().Value() {
		return true
	}

	if c[i].IsPayer() && c[i].AmountPaid().Value() > c[j].AmountPaid().Value() {
		return true
	}

	return false
}
