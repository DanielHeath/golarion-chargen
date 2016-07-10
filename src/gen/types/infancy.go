package types

// What were your first years like
type Infancy struct {
	Carer    string
	Location string
}

// FillInTheBlanks sets all unset properties to valid values
func (s Infancy) FillInTheBlanks() Infancy {
	if s.Carer == "" {
		s.Carer = SampleStr(
			"your mum",
			"your mum",
			"your mum",
			"your mum",
			"your dad",
			"your dad",
			"your dad",
			"your loving foster parents",
			"the clergy of {god}",
		)
	}

	if s.Location == "" {
		s.Location = SampleStr(
			"in a farmhouse",
			"in a castle",
			"in a wizards tower",
			"in the poor part of town",
			"in the ritzy part of town",
			"aboard a ship",
			"in the wilds",
			"in an army camp, wherever it went",
		)
	}
	return s
}
