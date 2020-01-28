package test

import (
	"bsocial/internal/receipt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Collection", func() {
	var sut receipt.Collection

	Context("When NewCollection is called", func() {
		BeforeEach(func() {
			sut = receipt.NewCollection()
		})

		It("Should have an empty collection of receipts (Bills)", func() {
			Expect(sut).Should(HaveLen(0))
		})

		It("Should have an empty list of participants", func() {
			Expect(sut.ParticipantList()).Should(HaveLen(0))
		})
	})

	Context("When a Receipt (Bill) added to the collection", func() {
		BeforeEach(func() {
			sut = receipt.NewCollection()

			mockedReceipt := NewReceiptMock()
			sut.Add(mockedReceipt)
		})

		It("Should have one receipt (Bill) inside the collection", func() {
			Expect(sut).Should(HaveLen(1))
		})
	})
})
