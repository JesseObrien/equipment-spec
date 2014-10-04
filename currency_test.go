package equipment_spec_test

import (
	. "github.com/jesseobrien/equipment-spec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Currency", func() {

	var currency *Currency

	BeforeEach(func() {
		currency = CurrencyFromString("CAD")
	})

	Describe("Compare currencies", func() {
		Context("two currencies, both with a name of CAD", func() {
			It("should be equal", func() {
				c := CurrencyFromString("CAD")
				Expect(currency.Equals(*c)).To(Equal(true))
			})
		})

		Context("two currencies, one CAD and one USD", func() {
			It("should no be equal", func() {
				c := CurrencyFromString("USD")
				Expect(currency.Equals(*c)).To(Equal(false))
			})
		})
	})

})
