package ewa

//DegreeType - type of wave degrees
type DegreeType uint

// Degrees of the wave
const (
	_                  = iota
	Primary DegreeType = iota
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

//CorrectionType type
type CorrectionType string

// Correction types
const (
	CTUnknown  CorrectionType = "unknown"
	CTZigzag                  = "zigzag"
	CTFlat                    = "flat"
	CTTriangle                = "triangle"
	CTCombo                   = "combo"
	CTTriple                  = "triple"
)

//Degree - gets degree type from string
func Degree(degree string) DegreeType {
	switch degree {
	case "Primary":
		return Primary
	case "Intermediate":
		return Intermediate
	case "Minor":
		return Minor
	case "Minute":
		return Minute
	case "Minuette":
		return Minuette
	case "Subminuette":
		return Subminuette
	case "Micro":
		return Micro
	case "Submicro":
		return Submicro
	case "Miniscule":
		return Miniscule
	case "Nano":
		return Nano
	case "Subnano":
		return Subnano
	case "Pico":
		return Pico
	}

	return Minuette
}
