package equipment_spec_test

import (
	. "github.com/jesseobrien/equipment-spec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rate", func() {

	var standardRate *Rate

	BeforeEach(func() {
		var p, err = PriceFromString("1000", "CAD")
		var rp, startErr, endErr = RentalPeriodFromString("2014/10/01 12:00 (EST)", "2014/10/02 12:00 (EST)")
		Expect(err).NotTo(HaveOccurred())
		Expect(startErr).NotTo(HaveOccurred())
		Expect(endErr).NotTo(HaveOccurred())
		standardRate = &Rate{price: p, rentalPeriod: rp}
	})

})
