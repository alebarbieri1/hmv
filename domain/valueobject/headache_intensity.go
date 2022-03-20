package valueobject

type HeadacheIntensity float64

func (i HeadacheIntensity) Float64() float64 { return float64(i) }

const (
	Undefined_HeadacheIntensity HeadacheIntensity = iota
	VeryLow_HeadacheIntensity
	Low_HeadacheIntensity
	Medium_HeadacheIntensity
	High_HeadacheIntensity
	VeryHigh_HeadacheIntensity
)

func (i HeadacheIntensity) String() string {
	switch i {
	case VeryLow_HeadacheIntensity:
		return "very-low"
	case Low_HeadacheIntensity:
		return "low"
	case Medium_HeadacheIntensity:
		return "medium"
	case High_HeadacheIntensity:
		return "high"
	case VeryHigh_HeadacheIntensity:
		return "very-high"
	}

	return "undefined"
}

func NewHeadacheIntensityFromString(s string) HeadacheIntensity {
	switch s {
	case "very-low":
		return VeryLow_HeadacheIntensity
	case "low":
		return Low_HeadacheIntensity
	case "medium":
		return Medium_HeadacheIntensity
	case "high":
		return High_HeadacheIntensity
	case "very-high":
		return VeryHigh_HeadacheIntensity
	}

	return Undefined_HeadacheIntensity
}
