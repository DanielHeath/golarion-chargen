package types

import "fmt"

type Parent struct {
	Name        string
	Surname     string
	Race        Race
	Occupation  string
	Nationality string
	// Hobbies
	// Stat adjustments
	// Wealth
}

func vowel(r byte) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'u' || r == 'o' || r == 'A' || r == 'E' || r == 'I' || r == 'U' || r == 'O'
}

func an(r Race) string {
	if r.Name == "" {
		return ""
	}

	if vowel(r.Name[0]) {
		return "an"
	}
	return "a"
}

func (p Parent) String() string {
	return fmt.Sprintf("%s %s, %s %s %s from %s", p.Name, p.Surname, an(p.Race), p.Race, p.Occupation, p.Nationality)
}

func RandomMum(race Race) Parent {
	return Parent{
		Name:        race.RandomFemaleName(),
		Surname:     race.RandomSurname(),
		Race:        race,
		Occupation:  RandomOccupation(),
		Nationality: RandomNationality(),
	}
}

func RandomDad(race Race) Parent {
	return Parent{
		Name:        race.RandomMaleName(),
		Surname:     race.RandomSurname(),
		Race:        race,
		Occupation:  RandomOccupation(),
		Nationality: RandomNationality(),
	}
}

func (p Parent) FillInTheBlanks(other Parent, child Character) Parent {
	if p.Race.Name == "" {
		p.Race = child.Race
		if child.Race.Halfbreed() && !other.Race.Halfbreed() {
			if other.Race.Name == child.Race.Mix1 {
				p.Race = GetRace(SampleStr(child.Race.Name, child.Race.Mix2))
			} else {
				p.Race = GetRace(SampleStr(child.Race.Name, child.Race.Mix1))
			}
		}
	}

	if p.Name == "" {
		p.Name = p.Race.RandomMaleName()
	}

	if p.Surname == "" {
		p.Surname = p.Race.RandomSurname()
		if child.Surname != "" && other.Surname != child.Surname {
			p.Surname = SampleStr(
				p.Surname,
				child.Surname,
				child.Surname,
				child.Surname,
			)
		}
	}

	if p.Occupation == "" {
		p.Occupation = RandomOccupation()
	}

	if p.Nationality == "" {
		p.Nationality = RandomNationality()
		// 2/3 of couples come from the same place
		// so a random dad probably comes from the same
		// place as mum does
		if other.Nationality != "" {
			p.Nationality = SampleStr(
				p.Nationality,
				other.Nationality,
				other.Nationality,
			)
		}
	}

	return p
}
