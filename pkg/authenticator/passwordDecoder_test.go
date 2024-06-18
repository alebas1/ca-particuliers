package authenticator_test

import (
	"testing"

	"github.com/alebas1/ca-particuliers/pkg/authenticator"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPass(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PasswordDecoder Suite")
}

var _ = Describe("ComputePasswordCombination", func() {
	var _ = Describe("Nominal cases", func() {
		It("should return the combination of a passcode", func() {
			By("computing the combination of the passcode 1234 with the key layout 1234")
			combination := authenticator.ComputePasswordCombination([]string{"1", "2", "3", "4"}, []string{"1", "2", "3", "4"})
			Expect(combination).To(Equal([]string{"0", "1", "2", "3"}))

			By("computing the combination of the passcode 1234 with the key layout 4321")
			combination = authenticator.ComputePasswordCombination([]string{"1", "2", "3", "4"}, []string{"4", "3", "2", "1"})
			Expect(combination).To(Equal([]string{"3", "2", "1", "0"}))
		})
	})
})
