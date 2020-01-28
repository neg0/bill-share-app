package test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPayback(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Payback Suite")
}
