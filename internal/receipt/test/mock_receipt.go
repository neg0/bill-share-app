package test

type ReceiptMock struct {
	spendersName       []string
	spendersCollection map[string]float64
	payerName          string
	totalAmount        float64
}

func NewReceiptMock() *ReceiptMock {
	return &ReceiptMock{
		spendersName: []string{
			"Alice",
			"Tom",
			"Henry",
		},
		spendersCollection: map[string]float64{
			"Alice": 15.25,
			"Tom":   10,
			"Henry": 20,
		},
		payerName:   "Alice",
		totalAmount: 45.25,
	}
}

func NewReceiptMock2() *ReceiptMock {
	return &ReceiptMock{
		spendersName: []string{
			"Alice",
			"Tom",
			"Henry",
			"Tammy",
			"Andy",
		},
		spendersCollection: map[string]float64{
			"Alice": 15.25,
			"Tom":   10,
			"Henry": 20,
			"Tammy": 10,
			"Andy":  15,
		},
		payerName:   "Alice",
		totalAmount: 70.25,
	}
}

func (r *ReceiptMock) SpendersName() []string {
	return r.spendersName
}

func (r *ReceiptMock) SpendersCollection() map[string]float64 {
	return r.spendersCollection
}

func (r *ReceiptMock) PayerName() string {
	return r.payerName
}

func (r *ReceiptMock) TotalAmount() float64 {
	return r.totalAmount
}
