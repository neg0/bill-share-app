package test

import (
	. "bsocial/internal/participant/test"
	"bsocial/internal/payback"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Payback", func() {
	var sut payback.Payback
	var mockedPayer *ParticipantMock
	var mockedPayee *ParticipantMock

	Context("When NewPayBack is called with payer and payee", func() {
		BeforeEach(func() {
			mockedPayer = NewParticipantMock("Alice", 29.95, 12.00)
			mockedPayee = NewParticipantMock("John", 9.50, 12.00)
			sut = payback.NewPayBack(mockedPayer, mockedPayee)
		})

		AfterEach(func() {
			var emptyInstance payback.Payback
			sut = emptyInstance
		})

		It("Should return the parsed Payer during initialisation", func() {
			Expect(sut.Payer()).To(Equal(mockedPayer))
		})

		It("Should return the parsed Payee during initialisation", func() {
			Expect(sut.Payee()).To(Equal(mockedPayee))
		})

		It("Should return the Paid Amount of payer", func() {
			Expect(sut.Amount().Value()).To(Equal(29.95))
		})

		It("Should print the process of payback from payer to payee and the amount", func() {
			Expect(sut.String()).To(Equal("Alice pays John £29.95"))
		})
	})
})
