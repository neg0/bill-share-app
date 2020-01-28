package test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBill(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bill internal Suite")
}
