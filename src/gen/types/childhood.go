package types

type Childhood struct {
	Activity string
	Location string
}

func (c Childhood) FillInTheBlanks() Childhood {
	if c.Activity == "" {
		c.Activity = SampleStr(
			"playing silly buggers",
		)
	}
	if c.Location == "" {
		c.Location = SampleStr(
			"on the castle walls",
		)
	}
	return c
}
