package ewa

//Sub - does it has subwaves
func (c *Correction) Sub() bool {
	return c.Zigzag != nil || c.Flat != nil ||
		c.Triangle != nil || c.Combo != nil || c.Triple != nil
}

//Type - gets correction type
func (c *Correction) Type() CorrectionType {
	if c.Zigzag != nil {
		return CTZigzag
	}

	if c.Flat != nil {
		return CTFlat
	}

	if c.Triangle != nil {
		return CTTriangle
	}

	if c.Combo != nil {
		return CTCombo
	}

	if c.Triple != nil {
		return CTTriple
	}

	return CTUnknown
}
