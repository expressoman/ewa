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
	*wave
	degree     DegreeType
	w1, w3, w5 *impulse
	w2, w4     Correctioner
}

//correction
type correction struct {
	*wave
	degree DegreeType
}

//zigzag
type zigzag struct {
	*wave
	degree DegreeType
	a, c   *impulse
	b      Correctioner
}

//flat
type flat struct {
	*wave
	degree DegreeType
	a, b   Correctioner
	c      *impulse
}

//triangle
type triangle struct {
	*wave
	degree        DegreeType
	a, b, c, d, e Correctioner
}

//combo
type combo struct {
	*wave
	degree  DegreeType
	w, x, y Correctioner
}

//triple
type triple struct {
	*wave
	degree         DegreeType
	w, x, y, x1, z Correctioner
}
