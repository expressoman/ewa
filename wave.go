package ewa

//ParentWave getter
func (w Wave) ParentWave() *Wave {
	return w.Parent
}

//NextWave getter
func (w Wave) NextWave() *Wave {
	return w.Next
}

//PrevWave getter
func (w Wave) PrevWave() *Wave {
	return w.Prev
}
