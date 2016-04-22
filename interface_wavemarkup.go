package ewa

import "time"

//WaveMarkuper - interface for wave object inside markup
type WaveMarkuper interface {

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

	// ============ Waver interface ends ============

	//Next - ref to obj representing next wave following after this
	Next() Waver

	//Prev - ref to obj representing prev wave following after this
	Prev() Waver

	//Parent - ref to parent wave obj
	Parent() Waver

	//Sub - array of pointers of the sub waves
	Sub() Wavers
}
