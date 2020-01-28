package test

import (
	"bsocial/internal/bill"

	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Repository", func() {
	var sut *bill.Repository
	var hasErr error

	Context("When database file path is correct", func() {
		BeforeEach(func() {
			_ = os.Setenv("FILE_PATH_BILLS", "./mock-collection.json")
			actual, err := bill.NewRepository()

			sut = actual
			hasErr = err
		})

		AfterEach(func() {
			_ = os.Unsetenv("FILE_PATH_BILLS")
			sut, hasErr = nil, nil
		})

		Describe("When instantiation method is called", func() {
			It("should have no error", func() {
				Expect(hasErr).Should(BeNil())
			})

			It("Should have two bills loaded into repository", func() {
				Expect(sut.FindAll()).Should(HaveLen(2))
			})
		})
	})

	Describe("When database file path is incorrect", func() {
		BeforeEach(func() {
			_ = os.Setenv("FILE_PATH_BILLS", "./mock-none-existence.json")
			actual, err := bill.NewRepository()

			sut = actual
			hasErr = err
		})

		AfterEach(func() {
			_ = os.Unsetenv("FILE_PATH_BILLS")
			sut, hasErr = nil, nil
		})

		Context("When instantiation method is called", func() {
			It("Should have an error", func() {
				Expect(hasErr).ShouldNot(BeNil())
			})
		})
	})

	Describe("When database file path hasn't been defined", func() {
		BeforeEach(func() {
			_ = os.Unsetenv("FILE_PATH_BILLS")
			actual, err := bill.NewRepository()

			sut = actual
			hasErr = err
		})

		AfterEach(func() {
			_ = os.Unsetenv("FILE_PATH_BILLS")
			sut, hasErr = nil, nil
		})

		Context("When instantiation method is called", func() {
			It("Should have an error", func() {
				Expect(hasErr).ShouldNot(BeNil())
			})
		})
	})
})
