package test

import (
	"fmt"
)

// MoneyMock is Mock type of Money
type MoneyMock struct {
	amount float64
}

// NewMoneyMock creates a new Mock for Money
func NewMoneyMock(amount float64) *MoneyMock {
	return &MoneyMock{amount}
}

// Value returns the floating value for mocked money
func (mm *MoneyMock) Value() float64 {
	return mm.amount
}

// String returns the formatted money with currency in string
func (mm *MoneyMock) String() string {
	return fmt.Sprintf("£%0.2f", mm.amount)
}
