package test

import (
	"bsocial/internal/payback"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Collection", func() {
	var sut payback.Collection

	Context("When NewCollection is called", func() {
		sut = payback.NewCollection()

		Context("And when new payback added to the collection", func() {
			mockPayBack := NewPayBackMock("Alice", "John", 78.95)
			actual := sut.Add(mockPayBack)

			It("Should have one in collection", func() {
				Expect(actual).To(HaveLen(1))
			})

			Context("When another new payback added to the collection", func() {
				mockPayBack := NewPayBackMock("Tom", "Henry", 78.95)
				actual := sut.Add(mockPayBack)

				It("Should have two in collection", func() {
					Expect(actual).To(HaveLen(2))
				})

				Context("When same participant added to the collection", func() {
					actual := sut.Add(mockPayBack)

					It("Should ignore the duplication and add a payback again to the collection", func() {
						Expect(actual).To(HaveLen(3))
					})
				})
			})
		})
	})
})
