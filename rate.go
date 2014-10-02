package equipment_spec

type Rate struct {
	price        Price
	rentalPeriod RentalPeriod
}

func (rate *Rate) getPrice() Price {
	return rate.price
}

func (rate *Rate) getPriceForDays(days int) {
	var total = PriceFromInt(0, rate.price.currency)

	for days > 0 {
		days = days - 1
		total = total.Add(&rate.price)
	}
}
