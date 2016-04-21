package ewa

import "time"

//Impulser interface
type Impulser interface {

	//Duration of the wave
	Duration() time.Duration

	//Len length of the price move of the wave
	Len() float64

	// Up - uptrend = true, downtrend = false
	Up() bool

	//Retrace - price level of retracement - Len()*float64
	Retrace(float64) float64

	//Begins - time when wave begins
	Begins() time.Time

	//Begins - time when wave ends
	Ends() time.Time

	//Starts - price where wave starts
	Starts() float64

	//Tops - price where wave tops
	Tops() float64

	//Slope - slope of the wave = Len() / Duration()
	Slope() float64

	// Waver interface

	//Degree of the wave
	Degree() DegreeType

	//Shows if wave has subdivisions or only start and end
	HasSub() bool

	//Waves of the impulse formation
	W1() Impulser
	W2() Correctioner
	W3() Impulser
	W4() Correctioner
	W5() Impulser

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
