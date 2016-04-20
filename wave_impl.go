package ewa

import "time"

//Duration of the wave
func (w wave) Duration() time.Duration {
	return w.End.T.Sub(w.Base.T)
}

//Len - length of the wave
func (w wave) Len() float64 {
	if w.Up() {
		return w.End.P - w.Base.P
	}

	return w.Base.P - w.End.P
}

//Up - is wave heading up or down
func (w wave) Up() bool {
	return w.End.P > w.Base.P
}

//Fib - fibonacci retracement
func (w wave) Fib(level float64) float64 {
	if w.Up() {
		return w.Base.P + w.Len()*level
	}

	return w.End.P + w.Len()*level
}

//Begins - time when wave begins
func (w wave) Begins() time.Time {
	return w.Base.T
}

//Ends - time when wave ends
func (w wave) Ends() time.Time {
	return w.End.T
}

//Starts - price where wave starts
func (w wave) Starts() float64 {
	return w.Base.P
}

//Tops - price where wave Tops
func (w wave) Tops() float64 {
	return w.End.P
}
