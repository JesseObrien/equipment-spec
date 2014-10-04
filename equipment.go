package equipment_spec

type Equipment struct {
	Name string
	Rate Rate
}

func EquipmentFromString(n string, r Rate) *Equipment {
	return &Equipment{Name: n, Rate: r}
}
