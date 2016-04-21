package ewa

func (c combo) HasSub() bool {
	return c.w != nil && c.x != nil && c.y != nil
}

func (c combo) Degree() DegreeType {
	return c.degree
}

func (c combo) Type() CorrectionType {
	return Combo
}
