package ewa

func ShortCorrection(base *Point, end *Point, degree uint8) *correction {
	return &Correction{Wave: &Wave{Base: base, End: end}, Degree: degree}
}

func FullCorrection(corrType interface{}, degree uint8) *correction {
	if zigzag, ok := corrType.(ZigZag); ok {
		return &Correction{ZigZag: &zigzag, Degree: degree}
	}

	if flat, ok := corrType.(Flat); ok {
		return &Correction{Flat: &flat, Degree: degree}
	}

	if triangle, ok := corrType.(Triangle); ok {
		return &Correction{Triangle: &triangle, Degree: degree}
	}

	if combo, ok := corrType.(Combo); ok {
		return &Correction{Combo: &combo, Degree: degree}
	}

	if triple, ok := corrType.(Triple); ok {
		return &Correction{Triple: &triple, Degree: degree}
	}

	return &Correction{Degree: degree}
}
