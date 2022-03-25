package valueobject

type OxygenSaturationLevel float64

// Float64 returns the OxygenSaturationLevel as a float64
func (i OxygenSaturationLevel) Float64() float64 { return float64(i) }

const (
	Undefined_OxygenSaturationLevel OxygenSaturationLevel = iota
	Normal_OxygenSaturationLevel
	Hypoxic_OxygenSaturationLevel
	SeverelyHypoxic_OxygenSaturationLevel
)

// String returns the string value of an OxygenSaturationLevel
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

// NewOxygenSaturationLevelFromString creates a new OxygenSaturationLevel from a given string
func NewOxygenSaturationLevelFromString(s string) OxygenSaturationLevel {
	switch s {
	case "normal":
		return Normal_OxygenSaturationLevel
	case "hypoxic":
		return Hypoxic_OxygenSaturationLevel
	case "severely-hypoxic":
		return SeverelyHypoxic_OxygenSaturationLevel
	}

	return Undefined_OxygenSaturationLevel
}
