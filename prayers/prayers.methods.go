package prayers

type OffsetUnit int

const (
	INVALID_UNIT OffsetUnit = iota
	OFFSET_UNIT_MINUTES
	OFFSET_UNIT_DEGREES
)

type MethodOffest struct {
	Value float64
	Unit  OffsetUnit
	From  string
}

func newOffset(value float64, unit OffsetUnit) MethodOffest {
	return MethodOffest{Value: value, Unit: unit}
}

type Mathhab int

const (
	INVALID_MATHHAB Mathhab = iota
	MATHHAB_STANDARD
	MATHHAB_HANAFI
	MATHHAB_MALIKY
	MATHHAB_SHAFEY
	MATHHAB_HANBALI
)

type AstFactor int

const (
	InvalidAsrFactor AstFactor = iota
	ASR_FACTOR_ONE
	ASR_FACTOR_TWO
)

type CalculationMethod struct {
	Name string
	Key  string

	DhuhrOffset       MethodOffest
	AsrFactor         AstFactor
	MaghribOffset     MethodOffest
	IshaOffset        MethodOffest
	NesfullailMathhab Mathhab
	FajrOffset        MethodOffest
}

// TODO: Update Dhuhr and Maghrib
var (
	MWL_CALCULATION_METHOD CalculationMethod = CalculationMethod{
		Name: "Muslim World League",
		Key:  "MWL",

		DhuhrOffset:       newOffset(0, OFFSET_UNIT_MINUTES),
		AsrFactor:         ASR_FACTOR_ONE,
		MaghribOffset:     newOffset(0, OFFSET_UNIT_MINUTES),
		IshaOffset:        newOffset(17, OFFSET_UNIT_DEGREES),
		NesfullailMathhab: MATHHAB_STANDARD,
		FajrOffset:        newOffset(18, OFFSET_UNIT_DEGREES),
	}
	ISNA_CALCULATION_METHOD CalculationMethod = CalculationMethod{
		Name: "Islamic Society of North America",
		Key:  "ISNA",

		DhuhrOffset:       newOffset(0, OFFSET_UNIT_MINUTES),
		AsrFactor:         ASR_FACTOR_ONE,
		MaghribOffset:     newOffset(0, OFFSET_UNIT_MINUTES),
		IshaOffset:        newOffset(15, OFFSET_UNIT_DEGREES),
		NesfullailMathhab: MATHHAB_STANDARD,
		FajrOffset:        newOffset(15, OFFSET_UNIT_DEGREES),
	}
	EGSA_CALCULATION_METHOD CalculationMethod = CalculationMethod{
		Name: "Egyptian General Survey Authority",
		Key:  "EGSA",

		DhuhrOffset:       newOffset(0, OFFSET_UNIT_MINUTES),
		AsrFactor:         ASR_FACTOR_ONE,
		MaghribOffset:     newOffset(0, OFFSET_UNIT_MINUTES),
		IshaOffset:        newOffset(17.5, OFFSET_UNIT_DEGREES),
		NesfullailMathhab: MATHHAB_STANDARD,
		FajrOffset:        newOffset(19.5, OFFSET_UNIT_DEGREES),
	}
	UQU_CALCULATION_METHOD CalculationMethod = CalculationMethod{
		Name: "Umm Al-Qura University, Makkah",
		Key:  "UQU",

		DhuhrOffset:       newOffset(0, OFFSET_UNIT_MINUTES),
		AsrFactor:         ASR_FACTOR_ONE,
		MaghribOffset:     newOffset(0, OFFSET_UNIT_MINUTES),
		IshaOffset:        newOffset(90, OFFSET_UNIT_MINUTES),
		NesfullailMathhab: MATHHAB_STANDARD,
		FajrOffset:        newOffset(18.5, OFFSET_UNIT_DEGREES),
	}
	UOK_CALCULATION_METHOD CalculationMethod = CalculationMethod{
		Name: "University of Islamic Studies, Karachi",
		Key:  "UOK",

		DhuhrOffset:       newOffset(0, OFFSET_UNIT_MINUTES),
		AsrFactor:         ASR_FACTOR_ONE,
		MaghribOffset:     newOffset(0, OFFSET_UNIT_MINUTES),
		IshaOffset:        newOffset(18, OFFSET_UNIT_DEGREES),
		NesfullailMathhab: MATHHAB_STANDARD,
		FajrOffset:        newOffset(18, OFFSET_UNIT_DEGREES),
	}
)

var CALCULATION_METHODS = map[string]CalculationMethod{
	"MWL":  MWL_CALCULATION_METHOD,
	"ISNA": ISNA_CALCULATION_METHOD,
	"EGSA": EGSA_CALCULATION_METHOD,
	"UQU":  UQU_CALCULATION_METHOD,
	"UOK":  UOK_CALCULATION_METHOD,
}

func GetCalculationMethodsMap() map[string]CalculationMethod {
	return CALCULATION_METHODS
}

func GetCalculationMethod(methodKey string) CalculationMethod {
	methods := GetCalculationMethodsMap()

	return methods[methodKey]
}
