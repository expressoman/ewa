package ewa

//DegreeType - type of wave degrees
type DegreeType uint

// Degrees of the wave
const (
	_              = iota
	Primary Degree = iota
	Intermediate
	Minor
	Minute
	Minuette
	Subminuette
	Micro
	Submicro
	Miniscule
	Nano
	Subnano
	Pico
)

//waves
type waves []*wave

//WaveType type
type WaveType string

//Wave types
const (
	Impulse    WaveType = "impulse"
	Correction          = "correction"
)

//ImpulseType type
type ImpulseType string

//ImpulseType
const (
	Simple   ImpulseType = "simple"
	Extended             = "extended"
	Diagonal             = "diagonal"
)

//CorrectionType type
type CorrectionType string

//Correction types
const (
	ZigZag   CorrectionType = "zigzag"
	Flat                    = "flat"
	Triangle                = "triangle"
	Combo                   = "combo"
	Triple                  = "triple"
)

//LabelType types
type LabelType string

//Label types
const (
	W1 LabelType = "1"
	W2           = "2"
	W3           = "3"
	W4           = "4"
	W5           = "5"
	A            = "a"
	B            = "b"
	C            = "c"
	D            = "d"
	E            = "e"
	W            = "w"
	X            = "x"
	Y            = "y"
	Z            = "z"
)
