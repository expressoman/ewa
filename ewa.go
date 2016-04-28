package ewa

import "fmt"

func (p Point) String() string {
	return fmt.Sprintf("%.2f-%s", p.P, p.T.Format("Jan 02 15:04"))
}

func (m Move) String() string {
	return fmt.Sprintf("%.2f-%.2f %s", m.Base.P, m.End.P, m.Base.T.Format("Jan 02 15:04"))
}
