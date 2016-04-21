package ewa

//Impulser interface
type Impulser interface {
	//W1 - ref to waves making impulse formation
	W1() *wave
	W2() *wave
	W3() *wave
	W4() *wave
	W5() *wave

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
