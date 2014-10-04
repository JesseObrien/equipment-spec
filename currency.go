package equipment_spec

type Currency struct {
	Name string
}

func CurrencyFromString(n string) *Currency {
	return &Currency{Name: n}
}

func (c *Currency) Equals(cu Currency) bool {
	return c.Name == cu.Name
}
