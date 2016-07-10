package types

var Occupations = []string{
	"Priest(ess)",
	"Artist",
	"Mercenary",
	"Farmer",
	"Criminal",
	"Entertainer",
	"Rebel",
	"Bounty Hunter",
	"Merchant",
	"Noble",
	"Scholar",
	"Sailor",
	"Pirate",
	"Slave",
	"Soldier",
	"Mercenary",
	"Knight",
}

// Wealth level

//     1 - subsistence farmer, failed pirate
//     2 - failed merchant, successful peasant
//     3 - minor nobility, successful merchant
//     4 - church leader, noble, merchant prince(ss)
//     5 - major noble family, church head

func RandomOccupation() string {
	return SampleStr(Occupations...)
}
