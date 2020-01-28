package test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestParticipant(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Participant Suite")
}
