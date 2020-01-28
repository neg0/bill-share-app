package test

import (
	"bsocial/internal/participant"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Participant", func() {
	// SUT stands for System/Subject Under Test
	var sut participant.Participant

	Context("When set a new name for participant", func() {
		BeforeEach(func() {
			sut = participant.NewParticipant()
			sut.SetName("Alice")
		})

		It("Should have a new name", func() {
			Expect(sut.Name()).To(Equal("Alice"))
		})
	})

	Context("When set a new paid amount for participant", func() {
		BeforeEach(func() {
			sut = participant.NewParticipant()
			sut.SetAmountPaid(21.95)
		})

		It("Should have a new set paid amount", func() {
			Expect(sut.AmountPaid().Value()).To(Equal(21.95))
		})
	})

	Context("When set a new spent amount for participant", func() {
		BeforeEach(func() {
			sut = participant.NewParticipant()
			sut.SetAmountSpent(95.15)
		})

		It("Should have a new spent amount", func() {
			Expect(sut.AmountSpent().Value()).To(Equal(95.15))
		})
	})

	Context("When participant has amount paid higher than zero", func() {
		BeforeEach(func() {
			sut = participant.NewParticipant()
			sut.SetAmountPaid(92.90)
		})

		It("Should be a payer among participants", func() {
			Expect(sut.IsPayer()).Should(BeTrue())
		})
	})

	Context("When participant has amount paid is zero", func() {
		BeforeEach(func() {
			sut = participant.NewParticipant()
			sut.SetAmountPaid(0)
		})

		It("Should NOT be a payer among participants", func() {
			Expect(sut.IsPayer()).Should(BeFalse())
		})
	})

	Context("When participant has paid the bill and also spent", func() {
		BeforeEach(func() {
			sut = participant.NewParticipant()
			sut.SetName("Alice")
			sut.SetAmountSpent(15)
			sut.SetAmountPaid(120.25)
		})

		It("Should have owned the amount paid minus amount spent", func() {
			Expect(sut.Owned().Value()).Should(Equal(105.25))
		})

		It("Should owe nobody", func() {
			Expect(sut.Owes(nil)).Should(BeNil())
		})
	})

	Context("When need to know whom participants owes to", func() {
		BeforeEach(func() {
			sut = participant.NewParticipant()
			sut.SetName("Alice")
			sut.SetAmountSpent(15)
			sut.SetAmountPaid(45.25)
		})

		It("Should owe nobody", func() {
			collectionOfParticipants := []participant.ParticipantDescriber{
				NewParticipantMock("Neo", 0, 10.95),
				NewParticipantMock("Morpheus", 0, 10.95),
				NewParticipantMock("Trinity", 0, 10.95),
			}

			Expect(sut.Owes(collectionOfParticipants)).Should(HaveLen(0))
		})

		It("Should owe one person (Neo)", func() {
			collectionOfParticipants := []participant.ParticipantDescriber{
				NewParticipantMock("Neo", 55.50, 10.95),
				NewParticipantMock("Morpheus", 0, 10.95),
				NewParticipantMock("Trinity", 0, 10.95),
			}

			actual := sut.Owes(collectionOfParticipants)
			Expect(actual).Should(HaveLen(1))
			Expect(actual[0].Name()).Should(Equal("Neo"))
		})
	})
})
