package equipment_spec

import (
	"time"
)

const RentalPeriodFormat = "2006/01/02 15:04 (EST)"

type RentalPeriod struct {
	startDate time.Time
	endDate   time.Time
}

func RentalPeriodFromString(start string, end string) (*RentalPeriod, error, error) {

	s, startErr := time.Parse(RentalPeriodFormat, start)

	e, endErr := time.Parse(RentalPeriodFormat, end)

	// @TODO Deal with errors here

	return &RentalPeriod{
		startDate: s,
		endDate:   e,
	}, startErr, endErr

}

func (rp *RentalPeriod) Equal(r *RentalPeriod) bool {
	return rp.startDate.Equal(r.startDate) && rp.endDate.Equal(r.endDate)
}

func (rp *RentalPeriod) IsWithinPeriod(p time.Time) bool {
	return rp.startDate.Before(p) && rp.endDate.After(p)
}

// @TODO fix this
func (rp *RentalPeriod) DurationInDays() int {
	return 1
}

func (rp *RentalPeriod) Duration() time.Duration {
	return rp.endDate.Sub(rp.startDate)
}
