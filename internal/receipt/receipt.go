package receipt

// Reader is definition of receipt
type Reader interface {
	SpendersName() []string
	SpendersCollection() map[string]float64
	PayerName() string
	TotalAmount() float64
}
