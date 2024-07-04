package cav1

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("getRegionalBankAlias", func() {
	var _ = Describe("Nominal cases", func() {
		It("should return the alias of a region", func() {
			By("getting the alias of the region 62")
			alias, err := getRegionalBankAlias("62")
			Expect(err).To(BeNil())
			Expect(alias).To(Equal("ca-norddefrance"))
		})

		It("should return an error if the region does not exist", func() {
			alias, err := getRegionalBankAlias("does-not-exist")
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(Equal("Region not found"))
			Expect(alias).To(Equal(""))
		})
	})
})
