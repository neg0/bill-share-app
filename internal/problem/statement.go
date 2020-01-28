package problem

import (
	"bsocial/internal/bill"
	"bsocial/internal/receipt"

	"fmt"
	"log"
	"sync"
)

var receiptService receipt.Service

func init() {
	repository, err := bill.NewRepository()
	if err != nil {
		log.Fatal(err)
	}

	receiptService = receipt.NewService(repository)
}

// Statement is a problem statement application object
type Statement struct {
	ReceiptService receipt.Service
	WaitGroup      sync.WaitGroup
	Mutex          sync.Mutex
}

// NewStatement creates a new instance of Statement with receipt service already being injected during instantiation
func NewStatement() *Statement {
	return &Statement{
		ReceiptService: receiptService,
	}
}

// Handle will deliver the Acceptance Criteria inside the README.md
func (s *Statement) Handle() {
	s.WaitGroup.Add(1)

	go func(w *sync.WaitGroup) {
		receiptCollection := s.ReceiptService.ReceiptCollection()
		for _, payback := range receiptCollection.Split() {
			s.Mutex.Lock()
			fmt.Println(payback.String())
			s.Mutex.Unlock()
		}
		w.Done()
	}(&s.WaitGroup)
	s.WaitGroup.Wait()
}
