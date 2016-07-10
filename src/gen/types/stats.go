package types

import (
	"fmt"
	"math/rand"
)

// Stats stores the D&D standard ability scores
type Stats struct {
	Strength     int
	Dexterity    int
	Constitution int
	Wisdom       int
	Intelligence int
	Charisma     int
}

func (s Stats) String() string {
	return fmt.Sprintf(`Strength: %d
Dexterity: %d
Constitution: %d
Wisdom: %d
Intelligence: %d
Charisma: %d
Fate points: %d
	`,
		s.Strength,
		s.Dexterity,
		s.Constitution,
		s.Wisdom,
		s.Intelligence,
		s.Charisma,
		s.BaseFatePoints(),
	)
}

func rollStat() int {
	return 9 + rand.Intn(6)
}

// BaseFatePoints returns the number of fate points this character can spend on e.g. rerolls
func (s Stats) BaseFatePoints() int {
	// max stats - current stats
	return (14 * 6) - s.Total()
}

// Total returns the sum of the numeric stats
func (s Stats) Total() int {
	return s.Strength +
		s.Dexterity +
		s.Constitution +
		s.Wisdom +
		s.Intelligence +
		s.Charisma
}

// FillInTheBlanks sets all unset properties to valid values
func (s Stats) FillInTheBlanks() Stats {
	if s.Strength == 0 {
		s.Strength = rollStat()
	}

	if s.Dexterity == 0 {
		s.Dexterity = rollStat()
	}

	if s.Constitution == 0 {
		s.Constitution = rollStat()
	}

	if s.Wisdom == 0 {
		s.Wisdom = rollStat()
	}

	if s.Intelligence == 0 {
		s.Intelligence = rollStat()
	}

	if s.Charisma == 0 {
		s.Charisma = rollStat()
	}

	return s
}
