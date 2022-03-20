package valueobject

type AbdominalPainIntensity float64

func (i AbdominalPainIntensity) Float64() float64 { return float64(i) }

const (
	Undefined_AbdominalPainIntensity AbdominalPainIntensity = iota
	VeryLow_AbdominalPainIntensity
	Low_AbdominalPainIntensity
	Medium_AbdominalPainIntensity
	High_AbdominalPainIntensity
	VeryHigh_AbdominalPainIntensity
)

func (i AbdominalPainIntensity) String() string {
	switch i {
	case VeryLow_AbdominalPainIntensity:
		return "very-low"
	case Low_AbdominalPainIntensity:
		return "low"
	case Medium_AbdominalPainIntensity:
		return "medium"
	case High_AbdominalPainIntensity:
		return "high"
	case VeryHigh_AbdominalPainIntensity:
		return "very-high"
	}

	return "undefined"
}

func NewAbdominalPainIntensityFromString(s string) AbdominalPainIntensity {
	switch s {
	case "very-low":
		return VeryLow_AbdominalPainIntensity
	case "low":
		return Low_AbdominalPainIntensity
	case "medium":
		return Medium_AbdominalPainIntensity
	case "high":
		return High_AbdominalPainIntensity
	case "very-high":
		return VeryHigh_AbdominalPainIntensity
	}

	return Undefined_AbdominalPainIntensity
}
