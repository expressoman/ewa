package ewa

//Correctioner interface
type Correctioner interface {
	Complex() bool
	Type() CorrectionType
}
