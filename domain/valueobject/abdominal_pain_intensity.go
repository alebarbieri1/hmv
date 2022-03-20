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
