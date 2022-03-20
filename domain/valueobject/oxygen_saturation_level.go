package valueobject

type OxygenSaturationLevel float64

func (i OxygenSaturationLevel) Float64() float64 { return float64(i) }

const (
	Undefined_OxygenSaturationLevel OxygenSaturationLevel = iota
	Normal_OxygenSaturationLevel
	Hypoxic_OxygenSaturationLevel
	SeverelyHypoxic_OxygenSaturationLevel
)

// Reference: https://cdn.shopify.com/s/files/1/0059/3992/files/image18.png
func (i OxygenSaturationLevel) String() string {
	switch i {
	case Normal_OxygenSaturationLevel:
		return "normal"
	case Hypoxic_OxygenSaturationLevel:
		return "hypoxic"
	case SeverelyHypoxic_OxygenSaturationLevel:
		return "severely-hypoxic"
	}

	return "undefined"
}
