package test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestReceipt(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Receipt Suite")
}
