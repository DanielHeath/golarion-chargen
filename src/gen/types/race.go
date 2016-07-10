package types

import "math/rand"

func RandomRace() Race {
	return Races()[rand.Intn(len(Races()))]
}

type Race struct {
	Name       string
	MotherRace string
	FatherRace string
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
	r.MotherRace = n.MotherRace
	r.FatherRace = n.FatherRace
	return nil
}

func (r Race) Mother() Race {
	if r.MotherRace == "" {
		return r
	}
	return GetRace(r.MotherRace)
}

func (r Race) Father() Race {
	if r.FatherRace == "" {
		return r
	}
	return GetRace(r.FatherRace)
}

func (r Race) String() string {
	return r.Name
}

func halfbreeds(name string, base1 string, base2 string) []Race {
	return []Race{
		Race{Name: name, MotherRace: name, FatherRace: base1},
		Race{Name: name, MotherRace: name, FatherRace: base2},
		Race{Name: name, MotherRace: base1, FatherRace: name},
		Race{Name: name, MotherRace: base1, FatherRace: base2},
		Race{Name: name, MotherRace: base2, FatherRace: base1},
		Race{Name: name, MotherRace: base2, FatherRace: name},
	}
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

func GetRaceMatchingParents(name string, dad string, mum string) Race {
	for _, race := range Races() {
		if (name == "" || race.Name == name) &&
			(mum == "" || race.MotherRace == mum) &&
			(dad == "" || race.FatherRace == dad) {

			return race
		}
	}
	return Race{}
}

func Races() []Race {
	if len(races) == 0 {
		races = append([]Race{
			Race{Name: "Human"},
			Race{Name: "Dwarf"},
			Race{Name: "Elf"},
			Race{Name: "Orc"},
			Race{Name: "Halfling"},
			Race{Name: "Gnome"},
		},
			append(
				halfbreeds("Half-elf", "Human", "Elf"),
				halfbreeds("Half-orc", "Human", "Orc")...,
			)...,
		)
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
