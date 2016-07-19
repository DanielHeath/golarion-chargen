package types

type Childhood struct {
	Activity string
	Location string
}

func (c Childhood) FillInTheBlanks() Childhood {
	if c.Activity == "" {
		c.Activity = SampleStr(
			"[fixme] playing silly buggers",
			"[fixme] prancing",
		)
	}
	if c.Location == "" {
		c.Location = SampleStr(
			"on the castle walls",
			"in the forest",
		)
	}
	return c
}
