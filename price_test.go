package equipment_spec_test

import (
	. "github.com/jesseobrien/equipment-spec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Price", func() {
	var additionPriceA *Price
	var additionPriceB *Price
	var expectedPriceC *Price

	BeforeEach(func() {
		var err error
		additionPriceA, err = PriceFromString("1000", "USD")
		Expect(err).NotTo(HaveOccurred())
		additionPriceB = PriceFromInt(1000, "USD")
		expectedPriceC = PriceFromInt(2000, "USD")
	})

	Describe("Adding two prices", func() {
		Context("Two $10.00USD items", func() {
			It("Should be $20.00USD", func() {
				Expect(additionPriceA.Add(additionPriceB)).To(Equal(expectedPriceC))
			})
		})
	})

	Describe("Multiplying a price", func() {
		Context("A $10.00USD item multiplied by 2", func() {
			It("Should be $20.00USD", func() {
				Expect(additionPriceA.Times(2)).To(Equal(expectedPriceC))
			})
		})
	})

})
