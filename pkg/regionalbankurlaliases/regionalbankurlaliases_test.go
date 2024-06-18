package regionalbankurlaliases_test

import (
	"testing"

	"github.com/alebas1/ca-particuliers/pkg/regionalbankurlaliases"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRegionalBankAliases(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RegionalBankAliases Suite")
}

var _ = Describe("getRegionalBankAlias", func() {
	var _ = Describe("Nominal cases", func() {
		It("should return the alias of a region", func() {
			By("getting the alias of the region 62")
			alias, err := regionalbankurlaliases.GetRegionalBankAlias("62")
			Expect(err).To(BeNil())
			Expect(alias).To(Equal("ca-norddefrance"))

			By("getting the alias of the region 01")
			alias, err = regionalbankurlaliases.GetRegionalBankAlias("01")
			Expect(err).To(BeNil())
			Expect(alias).To(Equal("ca-centrest"))
		})

		It("should return an error if the region does not exist", func() {
			alias, err := regionalbankurlaliases.GetRegionalBankAlias("does-not-exist")
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(Equal("Region not found"))
			Expect(alias).To(Equal(""))
		})
	})
})
