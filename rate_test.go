package equipment_spec_test

import (
	. "github.com/jesseobrien/equipment-spec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rate", func() {

	var standardRate *Rate

	BeforeEach(func() {
		var p, err = PriceFromString("1025", *CurrencyFromString("CAD"))
		var rp, startErr, endErr = RentalPeriodFromString("2014/10/01 12:00 (EST)", "2014/10/03 12:00 (EST)")

		Expect(err).NotTo(HaveOccurred())
		Expect(startErr).NotTo(HaveOccurred())
		Expect(endErr).NotTo(HaveOccurred())

		standardRate = &Rate{Price: *p, RentalPeriod: *rp}
	})

	Describe("Determine rental rates", func() {
		Context("Rate for 2 days @ $10.25/day", func() {
			It("Should be $20.50", func() {
				days := standardRate.RentalPeriod.DurationInDays()
				price := standardRate.PriceForDays(days)

				Expect(price.Value).To(Equal(2050))
			})
		})
	})

})
