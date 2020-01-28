package test

import (
	"bsocial/pkg/bill"

	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Collection", func() {
	var sut bill.Collection

	Context("When NewCollection instantiated", func() {
		sut = bill.NewCollection()

		Context("Then factory method `FromJSON` with valid JSON containing two bills is called", func() {
			testCase, _ := ioutil.ReadFile("./mock-collection.json")
			actual, err := sut.FromJSON(testCase)

			It("Should successfully create a new collection", func() {
				Expect(err).Should(BeNil())
			})
			It("Should have two bills", func() {
				Expect(actual).Should(HaveLen(2))
			})
		})

		Context("Then factory method `FromJSON` with invalid JSON is called", func() {
			actual, err := sut.FromJSON([]byte(`[{title: 123}]`))

			It("Should fail to create a new collection", func() {
				Expect(err).ShouldNot(BeNil())
			})

			It("Should have an empty result", func() {
				Expect(actual).Should(BeNil())
			})
		})
	})
})
