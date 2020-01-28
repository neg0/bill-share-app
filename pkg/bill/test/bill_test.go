package test

import (
	"bsocial/pkg/bill"
	"encoding/json"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bill", func() {
	var sut bill.Bill
	var testCase []byte
	var hasError error

	BeforeEach(func() {
		fileContents, _ := ioutil.ReadFile("./mock.json")
		testCase = fileContents
	})

	Context("When un-marshalling valid JSON data", func() {
		BeforeEach(func() {
			err := json.Unmarshal(testCase, &sut)
			hasError = err
		})

		It("Should successfully un-marshall without any error", func() {
			Expect(hasError).Should(BeNil())
		})

		It("Should Have a total amount of bill", func() {
			Expect(sut.TotalAmount(), 45.5)
		})

		It("Should have a payer name", func() {
			Expect(sut.PayerName()).Should(Equal("Tommen"))
		})

		It("Should have a list of five participants", func() {
			Expect(sut.SpendersName()).Should(HaveLen(5))
		})

		It("Should have a list of five participants with their name and spent amount", func() {
			Expect(sut.SpendersCollection()).Should(HaveLen(5))
		})
	})

	Context("When un-marshalling invalid JSON data", func() {
		BeforeEach(func() {
			testCase := []byte(`{ "title": 123 }`)
			err := json.Unmarshal(testCase, &sut)
			hasError = err
		})

		It("Should have an error", func() {
			Expect(hasError).ShouldNot(BeNil())
		})
	})
})
