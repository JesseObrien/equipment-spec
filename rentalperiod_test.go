package equipment_spec_test

import (
	. "github.com/jesseobrien/equipment-spec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Rentalperiod", func() {

	var firstPeriod *RentalPeriod
	var secondPeriod *RentalPeriod

	BeforeEach(func() {
		var startErr error
		var endErr error

		firstPeriod, startErr, endErr = RentalPeriodFromString("2014/10/01 22:39 (EST)", "2014/10/02 12:00 (EST)")

		Expect(startErr).NotTo(HaveOccurred())
		Expect(endErr).NotTo(HaveOccurred())

		secondPeriod, startErr, endErr = RentalPeriodFromString("2014/10/01 22:39 (EST)", "2014/10/02 12:00 (EST)")

		Expect(startErr).NotTo(HaveOccurred())
		Expect(endErr).NotTo(HaveOccurred())

	})

	Describe("Comparing Rental Periods", func() {
		Context("Two Rental Periods from 2014/10/1 to 2014/10/2", func() {
			It("Should be equal", func() {
				Expect(firstPeriod.Equal(secondPeriod)).To(Equal(true))
			})
		})
	})

	Describe("Verify dates on Rental Periods", func() {
		Context("A rental period from 2014/10/02 11:00", func() {
			It("Should be within 2014/10/01 11:39 - 2014/10/02 12:00", func() {

				var aNewPeriod, _ = time.Parse(RentalPeriodFormat, "2014/10/02 10:00 (EST)")
				Expect(firstPeriod.IsWithinPeriod(aNewPeriod)).To(Equal(true))
			})
		})

		Context("A rental period from 2014/10/01 00:00 to 2014/10/03 00:00", func() {
			It("Should have a duration of two days", func() {
				var twoDayPeriod, startErr, endErr = RentalPeriodFromString("2014/10/01 01:00 (EST)", "2014/10/03 01:00 (EST)")

				Expect(startErr).NotTo(HaveOccurred())
				Expect(endErr).NotTo(HaveOccurred())
				Expect(twoDayPeriod.DurationInDays()).To(Equal(2))
			})
		})
	})

	// @TODO Finish testing duration

})
