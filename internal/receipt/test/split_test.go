package test

import (
	"bsocial/internal/receipt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Split", func() {
	var sut receipt.Collection

	BeforeEach(func() {
		sut = receipt.NewCollection()
	})

	Context("When called method to get participants with empty collection", func() {
		It("Should have no splitting ways", func() {
			Expect(sut.Split()).Should(BeNil())
		})
	})

	Context("When a receipt is added to collection", func() {
		BeforeEach(func() {
			mockedReceipt := NewReceiptMock2()
			sut.Add(mockedReceipt)
		})

		It("Should be able to split the receipt (bill)", func() {
			Expect(sut.Split()).Should(HaveLen(4))
		})
	})
})
