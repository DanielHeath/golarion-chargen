package types

func RandomNationality() string {
	return SampleStr(Nationalities...)
}

var Nationalities = []string{
	// Garund
	"Alkenstar",
	"Geb",
	"Jalmeray",
	"Katapesh",
	"The Mwangi Expanse",
	"Nex",
	"Osirion",
	"The Shackles",
	"The Sodden Lands",
	"Thuvia",
	// Avistan
	"Absalom",
	"Andoran",
	"Cheliax",
	"Galt",
	"Hermea",
	"The Hold of Belkzen",
	"The Lands of the Linnorm Kings",
	"Mendev",
	"Qadira",
	"The Realm of the Mammoth Lords",
	"The River Kingdoms",
	"Taldor",
	"Thassilon",
	"Ustalav",
	"Varisia",
	"The Worldwound",
}
