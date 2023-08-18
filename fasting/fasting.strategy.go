package fasting

type FASTING_STRATEGY int

const (
	SCRAPPING FASTING_STRATEGY = iota + 1
	CALCULATION
)

type FastingStrategy interface {
	GetTodaysHijriDate() (HijriDate, error)
}

func NewFastingStrategy(strategy FASTING_STRATEGY) FastingStrategy {
	switch strategy {
	case SCRAPPING:
		return NewFastingScraper()
	}
	return nil
}
