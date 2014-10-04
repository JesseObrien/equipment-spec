package equipment_spec

import (
	"errors"
	"log"
	"strconv"
)

type Price struct {
	Value    int
	Currency Currency
}

func PriceFromString(v string, c Currency) (*Price, error) {
	pv, err := strconv.Atoi(v)
	return &Price{Value: pv, Currency: c}, err
}

func PriceFromInt(v int, c Currency) *Price {
	return &Price{Value: v, Currency: c}
}

func (price *Price) Add(other *Price) *Price {
	if _, err := price.hasMatchingCurrency(other); err != nil {
		log.Print(err)
	}
	return PriceFromInt(price.Value+other.Value, price.Currency)
}

func (price *Price) Multiply(multiplier float64) *Price {
	v := float64(price.Value) * multiplier
	return PriceFromInt(int(v), price.Currency)
}

func (price *Price) hasMatchingCurrency(other *Price) (bool, error) {
	if !price.Currency.Equals(other.Currency) {
		return false, errors.New("Currency's do not match.")
	}
	return true, nil
}
