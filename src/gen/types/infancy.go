package types

import "fmt"

// What were your first years like
type Infancy struct {
	Carer    string
	Location string
	Deity    string
}

func (s Infancy) String() string {
	if s.Deity != "" {
		return fmt.Sprintf("%s %s %s", s.Carer, s.Deity, s.Location)
	}
	return fmt.Sprintf("%s %s", s.Carer, s.Location)
}

// FillInTheBlanks sets all unset properties to valid values
func (s Infancy) FillInTheBlanks(deity string) Infancy {
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
			"the clergy of",
		)
		if s.Carer == "the clergy of" {
			if deity == "" {
				deity = SampleStr(Gods...)
			}

			s.Deity = SampleStr(
				deity,
				deity,
				deity,
				deity,
				deity,
				deity,
				deity,
				SampleStr(Gods...),
			)
		}
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
