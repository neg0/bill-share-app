package receipt

import (
	"bsocial/internal/bill"
)

// Service for receipt entity
type Service struct {
	billRepository bill.CollectionFinder
}

// NewService creates a new instance of service for receipt context
func NewService(repository bill.CollectionFinder) Service {
	return Service{
		billRepository: repository,
	}
}

// ReceiptCollection returns an array of Bills from injected Repository and presents them in receipt entity
func (s Service) ReceiptCollection() Collection {
	bills := s.billRepository.FindAll()
	collection := NewCollection()

	for _, individualBill := range bills {
		collection.Add(individualBill)
	}

	return collection
}
