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

// Angle returns the offset angle in degrees, if it's of type Angle.
//
// Otherwise, it returns 0
func (offset MethodOffest) Angle() float64 {
	if offset.IsAngle() {
		return offset.Value
	}
	return 0
}

// IsAngle checks whether the offset is of type angle or not
func (offset MethodOffest) IsAngle() bool {
	return offset.Unit == OFFSET_UNIT_DEGREES
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

		AsrFactor:     ASR_FACTOR_ONE,
		MaghribOffset: MethodOffest{Value: 0, Unit: OFFSET_UNIT_MINUTES, From: DAY_TIME_GHOROUB},
		IshaOffset:    MethodOffest{Value: 17, Unit: OFFSET_UNIT_DEGREES},
		FajrOffset:    MethodOffest{Value: 18, Unit: OFFSET_UNIT_DEGREES},
	}
	ISNA_CALCULATION_METHOD CalculationMethod = CalculationMethod{
		Name: "Islamic Society of North America",
		Key:  "ISNA",

		AsrFactor:     ASR_FACTOR_ONE,
		MaghribOffset: MethodOffest{Value: 0, Unit: OFFSET_UNIT_MINUTES, From: DAY_TIME_GHOROUB},
		IshaOffset:    MethodOffest{Value: 15, Unit: OFFSET_UNIT_DEGREES},
		FajrOffset:    MethodOffest{Value: 15, Unit: OFFSET_UNIT_DEGREES},
	}
	EGSA_CALCULATION_METHOD CalculationMethod = CalculationMethod{
		Name: "Egyptian General Survey Authority",
		Key:  "EGSA",

		AsrFactor:     ASR_FACTOR_ONE,
		MaghribOffset: MethodOffest{Value: 0, Unit: OFFSET_UNIT_MINUTES, From: DAY_TIME_GHOROUB},
		IshaOffset:    MethodOffest{Value: 17.5, Unit: OFFSET_UNIT_DEGREES},
		FajrOffset:    MethodOffest{Value: 19.5, Unit: OFFSET_UNIT_DEGREES},
	}
	UQU_CALCULATION_METHOD CalculationMethod = CalculationMethod{
		Name: "Umm Al-Qura University, Makkah",
		Key:  "UQU",

		AsrFactor:     ASR_FACTOR_ONE,
		MaghribOffset: MethodOffest{Value: 0, Unit: OFFSET_UNIT_MINUTES, From: DAY_TIME_GHOROUB},
		IshaOffset:    MethodOffest{Value: 90, Unit: OFFSET_UNIT_MINUTES, From: DAY_TIME_MAGHRIB},
		FajrOffset:    MethodOffest{Value: 18.5, Unit: OFFSET_UNIT_DEGREES},
	}
	UOK_CALCULATION_METHOD CalculationMethod = CalculationMethod{
		Name: "University of Islamic Studies, Karachi",
		Key:  "UOK",

		AsrFactor:     ASR_FACTOR_ONE,
		MaghribOffset: MethodOffest{Value: 0, Unit: OFFSET_UNIT_MINUTES, From: DAY_TIME_GHOROUB},
		IshaOffset:    MethodOffest{Value: 18, Unit: OFFSET_UNIT_DEGREES},
		FajrOffset:    MethodOffest{Value: 18, Unit: OFFSET_UNIT_DEGREES},
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
