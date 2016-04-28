package ewa

//HasSub - does it have subwaves
func (i *Impulse) HasSub() bool {
	return i.W1 != nil && i.W2 != nil && i.W3 != nil && i.W4 != nil && i.W5 != nil
}

//Rule1 - Wave 2 never bigger than Wave 1
func (i *Impulse) Rule1() bool {
	return i.W1.Len() >= i.W2.Len()
}

//Rule2 - Wave 4 never goes into territory of Wave 1
func (i *Impulse) Rule2() bool {
	if i.W1.Up() {
		return i.W1.Tops() < i.W4.Tops()
	}

	return i.W1.Tops() > i.W4.Tops()
}

//Rule3 - Wave 3 is never the shortest
func (i *Impulse) Rule3() bool {
	return i.W3.Len() >= i.W1.Len() || i.W3.Len() >= i.W5.Len()
}

//Check - all 3 rules together
func (i *Impulse) Check() bool {
	return i.Rule1() && i.Rule2() && i.Rule3()
}

//Failed or not
func (i *Impulse) Failed() bool {
	return i.W3.Tops() >= i.W5.Tops()
}

//Extended or not
func (i *Impulse) Extended() bool {
	return i.W3.Len() >= i.W5.Len()*2 || i.W3.Len() >= i.W1.Len()*2
}

//Diagonal or not
func (i *Impulse) Diagonal() bool {
	return i.W3.Len() >= i.W5.Len()*2 || i.W3.Len() >= i.W1.Len()*2
}
