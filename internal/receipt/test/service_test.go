package test

import (
	"bsocial/internal/bill/test"
	"bsocial/internal/payback"
	"bsocial/internal/receipt"

	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Service", func() {
	var sut receipt.Service

	BeforeEach(func() {
		_ = os.Setenv("FILE_PATH_BILLS", "./mock-collection.json")
		repositoryMock := test.NewRepositoryMock()

		sut = receipt.NewService(repositoryMock)
	})

	AfterEach(func() {
		_ = os.Unsetenv("FILE_PATH_BILLS")
	})

	Context("When new service is instantiated", func() {
		It("Should have a collection of receipts", func() {
			Expect(sut.ReceiptCollection()).Should(HaveLen(2))
		})

		It("Should have four transaction to payback the bill", func() {
			bills := sut.ReceiptCollection()
			actual := bills.Split()
			Expect(actual).Should(HaveLen(4))
		})

	})

	Context("When paying back the bills", func() {
		var paybackCollection payback.Collection
		allowedPayers := []string{
			"Ola",
			"Sam",
			"Sandy",
		}

		BeforeEach(func() {
			bills := sut.ReceiptCollection()
			paybackCollection = bills.Split()
		})

		It("Should have payer", func() {
			for _, transaction := range paybackCollection {
				Expect(allowedPayers).Should(ContainElement(transaction.Payer().Name()))
			}
		})

		It("Should have payee and should not be in payer list", func() {
			for _, transaction := range paybackCollection {
				Expect(allowedPayers).ShouldNot(ContainElement(transaction.Payee().Name()))
			}
		})
	})
})
