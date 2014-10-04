package equipment_spec

type Rate struct {
	Price        Price
	RentalPeriod RentalPeriod
}

func (rate *Rate) GetPrice() Price {
	return rate.Price
}

func (rate *Rate) PriceForDays(days int) Price {
	var total = PriceFromInt(0, rate.Price.Currency)

	for days > 0 {
		days = days - 1
		total = total.Add(&rate.Price)
	}

	return *total
}
