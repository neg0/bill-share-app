package test

import (
	"bsocial/pkg/money"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Money", func() {
	// sut (SUT) stands for System/Subject Under Test
	var sut money.Money

	BeforeEach(func() {
		sut = money.NewMoney(49.95)
	})

	Describe("When new Money is instantiated", func() {
		It("should have a value used for amount during instantiation", func() {
			Expect(sut.Value()).To(Equal(49.95))
		})

		It("Should print money formatted with two decimal points and default GBP currency symbol", func() {
			Expect(sut.String()).To(Equal("£49.95"))
		})
	})
})
