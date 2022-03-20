package valueobject

type ChestPainCharacteristics float64

func (c ChestPainCharacteristics) Float64() float64 { return float64(c) }

const (
	Undefined_ChestPainCharacteristics ChestPainCharacteristics = iota
	RadiatingToTheLeftArm_ChestPainCharacteristics
)

func (c ChestPainCharacteristics) String() string {
	switch c {
	case RadiatingToTheLeftArm_ChestPainCharacteristics:
		return "radiating-to-the-left-arm"
	}

	return "undefined"
}

func NewChestPainCharacteristicsFromString(s string) ChestPainCharacteristics {
	switch s {
	case "radiating-to-the-left-arm":
		return RadiatingToTheLeftArm_ChestPainCharacteristics
	}

	return Undefined_ChestPainCharacteristics
}
