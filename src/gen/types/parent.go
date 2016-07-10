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
