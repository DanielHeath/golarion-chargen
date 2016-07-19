package types

import "math/rand"

func RandomRace() Race {
	return Races()[rand.Intn(len(Races()))]
}

type Race struct {
	Name string

	// For half-breeds
	Mix1 string
	Mix2 string

	// Adjustments to stats
	Stats Stats
}

func (r Race) MarshalText() (text []byte, err error) {
	return []byte(r.Name), nil
}

func (r Race) Halfbreed() bool {
	return r.Name == "Half-orc" || r.Name == "Half-elf"
}

func (r *Race) UnmarshalText(text []byte) (err error) {
	n := TryGetRace(string(text))
	r.Name = n.Name
	r.Mix1 = n.Mix1
	r.Mix2 = n.Mix2
	return nil
}

func (r Race) String() string {
	return r.Name
}

var races []Race

func GetRace(name string) Race {
	for _, race := range Races() {
		if race.Name == name {
			return race
		}
	}
	panic("Bad name for GetRace: " + name)
}

func TryGetRace(name string) Race {
	for _, race := range Races() {
		if race.Name == name {
			return race
		}
	}
	return Race{}
}

func Races() []Race {
	if len(races) == 0 {
		races = []Race{
			Race{Name: "Human"},                                // +2 to a stat of your choice
			Race{Name: "Half-elf", Mix1: "Elf", Mix2: "Human"}, // +2 to a stat of your choice
			Race{Name: "Half-orc", Mix1: "Orc", Mix2: "Human"}, // +2 to a stat of your choice
			Race{Name: "Dwarf", Stats: Stats{Constitution: 2, Wisdom: 2, Charisma: -2}},
			Race{Name: "Elf", Stats: Stats{Constitution: -2, Intelligence: 2, Dexterity: 2}},
			Race{Name: "Orc", Stats: Stats{Strength: 4, Intelligence: -2, Wisdom: -2, Charisma: -2}},
			Race{Name: "Halfling", Stats: Stats{Strength: -2, Dexterity: 2, Charisma: 2}},
			Race{Name: "Gnome", Stats: Stats{Constitution: 2, Strength: -2, Charisma: 2}},
		}
	}

	return races
}

func (r Race) RandomMaleName() string {
	switch r.Name {
	case "Human":
		return SampleStr(MaleHumanNames...)
	case "Dwarf":
		return SampleStr(MaleDwarvenNames...)
	case "Elf":
		return SampleStr(MaleElvenNames...)
	case "Orc":
		return SampleStr(MaleOrcNames...)
	case "Half-elf":
		return SampleStr(append(MaleElvenNames, MaleHumanNames...)...)
	case "Half-orc":
		return SampleStr(append(MaleOrcNames, MaleHumanNames...)...)
	case "Halfling":
		return SampleStr(MaleHalflingNames...)
	case "Gnome":
		return SampleStr(MaleGnomeNames...)
	default:
		return "Robert"
	}
}

func (r Race) RandomFemaleName() string {
	switch r.Name {
	case "Human":
		return SampleStr(FemaleHumanNames...)
	case "Dwarf":
		return SampleStr(FemaleDwarvenNames...)
	case "Elf":
		return SampleStr(FemaleElvenNames...)
	case "Orc":
		return SampleStr(FemaleOrcNames...)
	case "Half-elf":
		return SampleStr(append(FemaleElvenNames, FemaleHumanNames...)...)
	case "Half-orc":
		return SampleStr(append(FemaleOrcNames, FemaleHumanNames...)...)
	case "Halfling":
		return SampleStr(FemaleHalflingNames...)
	case "Gnome":
		return SampleStr(FemaleGnomeNames...)
	default:
		return "Janet"
	}
}
func (r Race) RandomSurname() string {
	switch r.Name {
	case "Human":
		return SampleStr(HumanSurnames...)
	case "Dwarf":
		return SampleStr(DwarvenSurnames...)
	case "Elf":
		return SampleStr(ElvenSurnames...)
	case "Orc":
		return SampleStr(OrcSurnames...)
	case "Halfling":
		return SampleStr(HalflingSurnames...)
	case "Gnome":
		return SampleStr(GnomeSurnames...)
	case "Half-elf":
		return SampleStr(append(ElvenSurnames, HumanSurnames...)...)
	case "Half-orc":
		return SampleStr(append(OrcSurnames, HumanSurnames...)...)
	default:
		return "Jones"
	}
}
