package fasting

import "time"

type FastingService struct {
	FastingStrategy FastingStrategy
}

func NewFastingService() *FastingService {
	return &FastingService{}
}

func (s *FastingService) SetFastingStrategy(strategy FastingStrategy) {
	s.FastingStrategy = strategy
}

func (s FastingService) GetCurrentMonthFastingDays() FastingDays {
	todayHijriDate, err := s.FastingStrategy.GetTodaysHijriDate()
	if err != nil {
		panic(err)
	}
	daysUntil13th := s.calcDaysUntilDay13OfCurrentMonth(todayHijriDate)

	today := getStartOfDate(time.Now())
	return FastingDays{
		THIRTEENTH: addDays(today, int(daysUntil13th)),
		FOURTEENTH: addDays(today, int(daysUntil13th+1)),
		FIFTEENTH:  addDays(today, int(daysUntil13th+2)),
	}
}

func (FastingService) calcDaysUntilDay13OfCurrentMonth(todaysHijriDate HijriDate) int64 {
	return 13 - todaysHijriDate.Day
}
