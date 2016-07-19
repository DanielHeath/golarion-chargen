package types

import (
	"fmt"
	"math/rand"
)

// Character describes a generated characer
type Character struct {
	Name            string
	Surname         string
	Race            Race
	Sex             string
	Nationality     string
	Deity           string
	Mother          Parent
	Father          Parent
	Stats           Stats
	SpentFatePoints int
	Infancy         Infancy
	Childhood       Childhood
}

func (c Character) String() string {
	return fmt.Sprintf(`*****************
%s %s, a %s from %s.
Born to father %s
    and mother %s.
Stats:
%s
    `,
		c.Name,
		c.Surname,
		c.Race,
		c.Nationality,
		c.Father,
		c.Mother,
		c.Stats,
	)
}

// NewCharacter generates a random character with stats filled in
func NewCharacter() Character {
	return Character{}.FillInTheBlanks()
}

// FillInTheBlanks sets all unset properties to valid values
func (c Character) FillInTheBlanks() Character {
	if c.Sex == "" {
		c.Sex = SampleStr("Male", "Female")
	}

	if (c.Race == Race{}) {
		c.Race = RandomRace()
	}

	if c.Name == "" {
		c.Name = c.Race.RandomFemaleName()
		if c.Sex == "Male" {
			c.Name = c.Race.RandomMaleName()
		}
	}

	c.Mother = c.Mother.FillInTheBlanks(c.Father, c)
	c.Father = c.Father.FillInTheBlanks(c.Mother, c)

	if c.Surname == "" {
		c.Surname = SampleStr(
			// bias towards patrimony to balance golarion norms
			// with not wanting everything to be run by dudes
			c.Father.Surname,
			c.Father.Surname,
			c.Father.Surname,
			c.Father.Surname,
			c.Mother.Surname,
			c.Mother.Surname,
			c.Race.RandomSurname(),
		)
	}

	if c.Deity == "" {
		c.Deity = SampleStr(Gods...)
	}

	if c.Nationality == "" {
		c.Nationality = SampleStr(
			c.Father.Nationality,
			c.Father.Nationality,
			c.Father.Nationality,
			c.Mother.Nationality,
			c.Mother.Nationality,
			c.Mother.Nationality,
			SampleStr(Nationalities...),
		)
	}

	c.Stats = c.Stats.FillInTheBlanks()
	c.Infancy = c.Infancy.FillInTheBlanks(c.Deity)
	c.Childhood = c.Childhood.FillInTheBlanks()

	return c
}

// SampleStr returns one of the given strings at random
func SampleStr(options ...string) string {
	return options[rand.Intn(len(options))]
}
