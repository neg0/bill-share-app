package test

import (
	"bsocial/internal/participant"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Collection", func() {
	var sut participant.Collection

	BeforeEach(func() {
		sut = participant.NewCollection()
	})

	AfterEach(func() {
		sut = nil
	})

	Context("When NewCollection called", func() {
		It("Should have an empty list of participants", func() {
			Expect(sut).To(HaveLen(0))
		})

		It("Should have an empty list of spenders", func() {
			Expect(sut.Spenders()).To(HaveLen(0))
		})

		It("Should have an empty list of payers", func() {
			Expect(sut.Payers()).To(HaveLen(0))
		})

		Context("When `From` factory method called with list of three participants and their names", func() {
			actual := sut.From([]string{
				"Alice",
				"Tom",
				"John",
			})

			It("Should have three participants", func() {
				Expect(actual).To(HaveLen(3))
			})
		})
	})
})
