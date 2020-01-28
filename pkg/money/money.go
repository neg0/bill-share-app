package money

import "fmt"

// defaultCurrencySymbolGBP is Currency Symbol UK (GBP)
const defaultCurrencySymbolGBP = "£"

// Money is a ValueObject for Money with float64 value
type Money struct {
	amount         float64
	currencySymbol string
}

// API is a definition of Money
type API interface {
	Value() float64
	String() string
}

// NewMoney will create a new instance of Money Value Object
// Properties in Money are private to make Money as a ValueObject Immutable
func NewMoney(amount float64) Money {
	return Money{
		amount:         amount,
		currencySymbol: defaultCurrencySymbolGBP,
	}
}

// Value returns a value of Money in float
func (m Money) Value() float64 {
	return m.amount
}

// String prints formatted Money in string with symbol and two decimal points
func (m Money) String() string {
	return fmt.Sprintf("%s%0.2f", m.currencySymbol, m.amount)
}
