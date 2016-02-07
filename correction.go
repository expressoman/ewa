package ewa

type Correction struct {
	*Wave
	ZigZag   *ZigZag
	Flat     *Flat
	Triangle *Triangle
	Combo    *Combo
	Triple   *Triple
}

type ZigZag struct {
	A, C Impulse
	B    Correction
}
type Flat struct {
	A, B Correction
	C    Impulse
}
type Triangle struct {
	A, B, C, D, E Correction
}
type Combo struct {
	W, X, Y Correction
}
type Triple struct {
	W, X, Y, X1, Z Correction
}
