package ewa

import "time"

//Duration of the wave
func (m Move) Duration() time.Duration {
	return m.End.T.Sub(m.Base.T)
}

//Len - length of the wave
func (m Move) Len() float64 {
	if m.Up() {
		return m.End.P - m.Base.P
	}

	return m.Base.P - m.End.P
}

//Up - is wave heading up or down
func (m Move) Up() bool {
	return m.End.P > m.Base.P
}

//Retrace - fibonacci retracement
func (m Move) Retrace(level float64) float64 {
	if m.Up() {
		return m.Base.P + m.Len()*level
	}

	return m.End.P + m.Len()*level
}

//Begins - time when wave begins
func (m Move) Begins() time.Time {
	return m.Base.T
}

//Ends - time when wave ends
func (m Move) Ends() time.Time {
	return m.End.T
}

//Starts - price where wave starts
func (m Move) Starts() float64 {
	return m.Base.P
}

//Tops - price where wave Tops
func (m Move) Tops() float64 {
	return m.End.P
}

//Slope of the wave
func (m Move) Slope() float64 {
	return m.Len() / float64(m.Duration())
}
