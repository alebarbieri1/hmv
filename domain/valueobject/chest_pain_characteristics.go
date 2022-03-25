package valueobject

type ChestPainCharacteristics float64

// Float64 returns the ChestPainCharacteristics as a float64
func (c ChestPainCharacteristics) Float64() float64 { return float64(c) }

const (
	Undefined_ChestPainCharacteristics ChestPainCharacteristics = iota
	RadiatingToTheLeftArm_ChestPainCharacteristics
)

// String returns the string value of an ChestPainCharacteristics
func (c ChestPainCharacteristics) String() string {
	switch c {
	case RadiatingToTheLeftArm_ChestPainCharacteristics:
		return "radiating-to-the-left-arm"
	}

	return "undefined"
}

// NewChestPainCharacteristicsFromString creates a new ChestPainCharacteristics from a given string
func NewChestPainCharacteristicsFromString(s string) ChestPainCharacteristics {
	switch s {
	case "radiating-to-the-left-arm":
		return RadiatingToTheLeftArm_ChestPainCharacteristics
	}

	return Undefined_ChestPainCharacteristics
}
