package ewa

import "time"

//Waver - interface for base wave
type Waver interface {

	//Finished - Formula: Ends() - time of last tick <= 0
	Finished() bool

	//Duration of the wave
	Duration() time.Duration

	//Len length of the price move of the wave
	Len() float64

	// Up - uptrend = true, downtrend = false
	Up() float64

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

	//Next - ref to obj representing next wave following after this
	Next() *Wave

	//Prev - ref to obj representing prev wave following after this
	Prev() *Wave

	//Parent - ref to parent wave obj
	Parent() *Wave

	//Sub - array of pointers of the sub waves
	Sub() []*Wave
}

//Correctioner interface
type Correctioner interface {
	Complex() bool
	Type() WaveType
}

//Markup structure
type Markup interface {
	Wave(WaveType) []*Wave
}
