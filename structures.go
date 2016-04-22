package ewa

import "time"

//point on chart
type point struct {
	p float64
	t time.Time
}

//wave
type wave struct {
	base *point
	end  *point
}

//impulse
type impulse struct {
	Waver
	degree     DegreeType
	w1, w3, w5 Impulser
	w2, w4     Correctioner
}

//correction
type correction struct {
	Waver
	degree DegreeType
}

//zigzag
type zigzag struct {
	Waver
	degree DegreeType
	a, c   Impulser
	b      Correctioner
}

//flat
type flat struct {
	Waver
	degree DegreeType
	a, b   Correctioner
	c      Impulser
}

//triangle
type triangle struct {
	Waver
	degree        DegreeType
	a, b, c, d, e Correctioner
}

//combo
type combo struct {
	Waver
	degree  DegreeType
	w, x, y Correctioner
}

//triple
type triple struct {
	Waver
	degree         DegreeType
	w, x, y, x1, z Correctioner
}
