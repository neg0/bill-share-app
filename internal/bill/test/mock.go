package test

import (
	internalBill "bsocial/internal/bill"
	"bsocial/pkg/bill"
)

// RepositoryMock is mock object for Repository
type RepositoryMock struct{}

// NewRepositoryMock creates a new instance of RepositoryMock
func NewRepositoryMock() RepositoryMock {
	return RepositoryMock{}
}

// FindAll returns an array of bills
func (rm RepositoryMock) FindAll() bill.Collection {
	repository, _ := internalBill.NewRepository()

	return repository.FindAll()
}
