package ewa

//Impulser interface
type Impulser interface {
	//Degree of the wave
	Degree() DegreeType

	//Shows if wave has subdivisions or only start and end
	HasSub() bool

	//Waves of the impulse formation
	W1() *impulse
	W2() *correction
	W3() *impulse
	W4() *correction
	W5() *impulse

	//Impulse rules
	Rule1() bool
	Rule2() bool
	Rule3() bool

	//Check - Rule1() && Rule2() && Rule3()
	Check() bool

	//Failed - price of the W5 does not go further W5
	Failed() bool

	//Extended - Len() of this wave is the biggest
	//among all child waves of the parent wave
	Extended() bool

	//Diagonal - wave 4 overlaps wave 1
	Diagonal() bool
}
