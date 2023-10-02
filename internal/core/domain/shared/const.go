package shared

var Location = map[string]string{
	"Abuja": "ABJ",
	"Lagos": "LAG",
}

var Denomination = map[int]string{
	5:    "F",
	10:   "T",
	20:   "TW",
	50:   "FT",
	100:  "HU",
	200:  "TH",
	500:  "FH",
	1000: "OT",
}

var PrefixType = map[string]string{
	"Single":          "A",
	"Double":          "AA",
	"Fraction Single": "A|01",
	"Fraction Double": "AA|01",
}



var DepartmentModule = map[string]string{
	"LAG BNT ICT":           "Box Manager,Packing Manager,Confirmation Manager,Evacuation Manager,Movement Manager,Invoice Manager,Generation Manager",
	"LAG BNT DOCUMENTATION": "Confirmation Manager",
	"LAG BNT FINISHING":     "Box Manager,Packing Manager",
	"LAG BNT INVOICE":       "Invoice Manager",
	"LAG BNT MATERIAL":      "Evacuation Manager,Movement Manager",
	"ABJ BNT ICT":           "Box Manager,Packing Manager,Confirmation Manager,Evacuation Manager,Movement Manager,Invoice Manager,Generation Manager",
	"ABJ BNT DOCUMENTATION": "Confirmation Manager",
	"ABJ BNT FINISHING":     "Box Manager, Packing Manager",
	"ABJ BNT INVOICE":       "Invoice Manager",
	"ABJ BNT MATERIAL":      "Evacuation Manager,Movement Manager",
}

const (
	Normal Status = iota
	Packed
	Printed
	Confirmed
	Evacuated
	Moved
)
const (
	Generation Activity = iota
	Packing
	Confirmation
	Evacuation
	Movement
	UnPacking
	UnConfirmation
	UnEvacuation
	Create
	Update
	Delete
)
