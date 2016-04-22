package ewa

import "time"

//Waver - interface for base wave
type Waver interface {

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
}
