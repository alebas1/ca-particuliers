package cav1

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPass(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "")
}

var _ = Describe("ComputePasswordCombination", func() {
	var _ = Describe("Nominal cases", func() {
		It("should return the combination of a passcode", func() {
			By("computing the combination of the passcode 1234 with the key layout 1234")
			combination, err := computePasscodeCombination([]string{"1", "2", "3", "4"}, []string{"1", "2", "3", "4"})
			Expect(err).To(BeNil())
			Expect(combination).To(Equal([]string{"0", "1", "2", "3"}))

			By("computing the combination of the passcode 1234 with the key layout 4321")
			combination, err = computePasscodeCombination([]string{"1", "2", "3", "4"}, []string{"4", "3", "2", "1"})
			Expect(err).To(BeNil())
			Expect(combination).To(Equal([]string{"3", "2", "1", "0"}))
		})

		It("should return error if the key layout is invalid", func() {
			passcode := []string{"1", "2", "3", "4"}
			keyLayout := []string{"1", "2", "4"} // missing "3"
			combination, err := computePasscodeCombination(passcode, keyLayout)
			Expect(err).ToNot(BeNil())
			Expect(combination).To(Equal([]string{"0", "1"}))
		})
	})
})
