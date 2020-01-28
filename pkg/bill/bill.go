package bill

import "reflect"

// Spenders is a type of participants is map with string key and float value e.g. "Alice": 24.95
type Spenders map[string]float64

// Bill is a representation of data for marshall and un-marshalling JSON
type Bill struct {
	Title    string   `json:"title"`
	Spenders Spenders `json:"spend"`
	Payer    string   `json:"payer"`
	Total    float64  `json:"total"`
}

// SpendersList returns the list of participant names
func (b Bill) SpendersName() []string {
	participantList := reflect.ValueOf(b.Spenders).MapKeys()
	participants := make([]string, len(participantList))

	for index, value := range participantList {
		participants[index] = value.String()
	}

	return participants
}

// SpendersCollection returns a map of all participants with their name and their spending balance
func (b Bill) SpendersCollection() map[string]float64 {
	return b.Spenders
}

// PayerName returns a name of payer in each bill
func (b Bill) PayerName() string {
	return b.Payer
}

// TotalAmount returns the total amount spent for each bill
func (b Bill) TotalAmount() float64 {
	return b.Total
}
