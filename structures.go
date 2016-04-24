package ewa

import "time"

//Point struc
type Point struct {
	T time.Time
	P float64
}

//Move struc
type Move struct {
	Base *Point
	End  *Point
}

//Wave struc
type Wave struct {
	*Move
	Degree DegreeType

	Prev   *Wave
	Next   *Wave
	Parent *Wave
}

//Impulse struc
type Impulse struct {
	*Wave

	W1 *Impulse
	W2 *Correction
	W3 *Impulse
	W4 *Correction
	W5 *Impulse
}

//Correction struc
type Correction struct {
	*Wave

	Zigzag   *Zigzag
	Flat     *Flat
	Triangle *Triangle
	Combo    *Combo
	Triple   *Triple
}

// Corrective structures

//Zigzag struct
type Zigzag struct {
	A, C *Impulse
	B    *Correction
}

//Flat struct
type Flat struct {
	A, B *Correction
	C    *Impulse
}

//Triangle struct
type Triangle struct {
	A, B, C, D, E *Correction
}

//Combo struct
type Combo struct {
	W, X, Y *Correction
}

//Triple struct
type Triple struct {
	W, X, Y, X2, Z *Correction
}
