package test

import (
	"bsocial/internal/receipt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Participants", func() {
	var sut receipt.Collection

	BeforeEach(func() {
		sut = receipt.NewCollection()
	})

	Context("When called method to get participants with empty collection", func() {
		It("Should have no participants", func() {
			Expect(sut.Participants()).Should(BeNil())
		})
	})

	Context("When called method to get participants", func() {
		BeforeEach(func() {
			mockedReceipt := NewReceiptMock()
			sut.Add(mockedReceipt)
		})

		It("Should have three participants", func() {
			Expect(sut.Participants()).Should(HaveLen(3))
		})
	})
})
