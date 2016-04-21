package ewa

import "time"

//Duration of the wave
func (w wave) Duration() time.Duration {
	return time.Second
}

//Len - length of the wave
func (w wave) Len() float64 {
	if w.Up() {
		return w.end.p - w.base.p
	}

	return w.base.p - w.end.p
}

//Up - is wave heading up or down
func (w wave) Up() bool {
	return w.end.p > w.base.p
}

//Fib - fibonacci retracement
func (w wave) Retrace(level float64) float64 {
	if w.Up() {
		return w.base.p + w.Len()*level
	}

	return w.end.p + w.Len()*level
}

//Begins - time when wave begins
func (w wave) Begins() time.Time {
	return w.base.t
}

//Ends - time when wave ends
func (w wave) Ends() time.Time {
	return w.end.t
}

//Starts - price where wave starts
func (w wave) Starts() float64 {
	return w.base.p
}

//Tops - price where wave Tops
func (w wave) Tops() float64 {
	return w.end.p
}

//Tops - price where wave Tops
func (w wave) Slope() float64 {
	return w.Len() / float64(w.Duration())
}
