package equipment_spec_test

import (
	. "github.com/jesseobrien/equipment-spec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Price", func() {

	var additionPriceA *Price
	var additionPriceB *Price
	var currency *Currency

	BeforeEach(func() {
		var err error
		currency = CurrencyFromString("CAD")

		additionPriceA, err = PriceFromString("1000", *currency)
		Expect(err).NotTo(HaveOccurred())
		additionPriceB = PriceFromInt(1000, *currency)
	})

	Describe("Adding two prices", func() {
		Context("Two $10.00USD items", func() {
			It("Should be $20.00USD", func() {
				expectedPrice := PriceFromInt(2000, *currency)

				Expect(additionPriceA.Add(additionPriceB)).To(Equal(expectedPrice))
			})
		})
	})

	Describe("Multiplying a price", func() {
		Context("A $10.00USD item multiplied by 2", func() {
			It("Should be $20.00USD", func() {
				expectedPrice := PriceFromInt(2000, *currency)

				Expect(additionPriceA.Multiply(2)).To(Equal(expectedPrice))
			})
		})
	})

})
