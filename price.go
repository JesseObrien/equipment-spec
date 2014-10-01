package equipment_spec

import (
	"errors"
	"log"
	"strconv"
)

type Price struct {
	value    int
	currency string
}

func PriceFromString(v string, c string) (*Price, error) {
	pv, err := strconv.Atoi(v)
	return &Price{value: pv, currency: c}, err
}

func PriceFromInt(v int, c string) *Price {
	return &Price{value: v, currency: c}
}

func (price *Price) Add(other *Price) *Price {
	if _, err := price.hasMatchingCurrency(other); err != nil {
		log.Print(err)
	}
	return PriceFromInt(price.value+other.value, price.currency)
}

func (price *Price) Times(multiplier float64) *Price {
	v := float64(price.value) * multiplier
	return PriceFromInt(int(v), price.currency)
}

func (price *Price) hasMatchingCurrency(other *Price) (bool, error) {
	if price.currency != other.currency {
		return false, errors.New("Currency's do not match.")
	}
	return true, nil
}
