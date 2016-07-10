package types

import (
	"fmt"
	"math/rand"
)

type Stats struct {
	Strength     uint
	Dexterity    uint
	Constitution uint
	Wisdom       uint
	Intelligence uint
	Charisma     uint
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

func rollStat() uint {
	return uint(9 + rand.Intn(6))
}

type Character struct {
	Name        string
	Surname     string
	Race        Race
	Sex         string
	Nationality string
	Mother      Parent
	Father      Parent
	Stats       Stats
}

func (s Stats) BaseFatePoints() uint {
	// max stats - current stats
	return (14 * 6) - s.Total()
}

func (s Stats) Total() uint {
	return s.Strength +
		s.Dexterity +
		s.Constitution +
		s.Wisdom +
		s.Intelligence +
		s.Charisma
}

func (c Character) String() string {
	// todo print stats and fatepoints
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

func NewCharacter() Character {
	return Character{}.FillInTheBlanks()
}

func (c Character) RepickDad() Character {
	c.Father = Parent{}
	c.Surname = ""
	c.Nationality = ""
	return c.FillInTheBlanks()
}

func (c Character) RepickMum() Character {
	c.Mother = Parent{}
	c.Surname = ""
	c.Nationality = ""
	return c.FillInTheBlanks()
}

func (c Character) FillInTheBlanks() Character {
	if c.Sex == "" {
		c.Sex = SampleStr("Male", "Female")
	}

	if (c.Race == Race{}) {
		c.Race = RandomRace()
	}

	if c.Race.Halfbreed() {
		c.Race = GetRaceMatchingParents(
			c.Race.Name,
			c.Father.Race.Name,
			c.Mother.Race.Name,
		)
	}

	if c.Name == "" {
		c.Name = c.Race.RandomFemaleName()
		if c.Sex == "Male" {
			c.Name = c.Race.RandomMaleName()
		}
	}

	if (c.Mother == Parent{}) {
		c.Mother = RandomMum(c.Race.Mother())

		if c.Surname != "" {
			c.Mother.Surname = SampleStr(
				c.Mother.Surname,
				c.Surname,
				c.Surname,
				c.Surname,
			)
		}

		if (c.Father != Parent{}) {
			// Dad was set but mum wasnt.
			// 2/3 of couples come from the same place, so
			// Mum probably comes from where Dad comes from
			c.Mother.Nationality = SampleStr(
				c.Father.Nationality,
				c.Father.Nationality,
				c.Mother.Nationality,
			)
		}
	}

	if (c.Father == Parent{}) {
		c.Father = RandomDad(c.Race.Father())
		if c.Race.Halfbreed() && !c.Mother.Race.Halfbreed() {
			// dads race needs to be the other half!

		}

		if c.Surname != "" && c.Mother.Surname != c.Surname {
			c.Father.Surname = SampleStr(
				c.Father.Surname,
				c.Surname,
				c.Surname,
				c.Surname,
			)
		}

		// 2/3 of couples come from the same place
		// so a random dad probably comes from the same
		// place as mum does
		c.Father.Nationality = SampleStr(
			c.Father.Nationality,
			c.Mother.Nationality,
			c.Mother.Nationality,
		)
	}

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
	return c
}

func (c Stats) FillInTheBlanks() Stats {
	if c.Strength == 0 {
		c.Strength = rollStat()
	}

	if c.Dexterity == 0 {
		c.Dexterity = rollStat()
	}

	if c.Constitution == 0 {
		c.Constitution = rollStat()
	}

	if c.Wisdom == 0 {
		c.Wisdom = rollStat()
	}

	if c.Intelligence == 0 {
		c.Intelligence = rollStat()
	}

	if c.Charisma == 0 {
		c.Charisma = rollStat()
	}

	return c
}

func SampleStr(options ...string) string {
	return options[rand.Intn(len(options))]
}
