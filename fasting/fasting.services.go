package fasting

import "time"

type FastingService struct {
	FastingStrategy FastingStrategy
}

func NewFastingService(strategy FastingStrategy) FastingService {
	return FastingService{FastingStrategy: strategy}
}

func (service FastingService) GetCurrentMonthFastingDays() FastingDays {
	todayHijriDate, err := service.FastingStrategy.GetTodaysHijriDate()
	if err != nil {
		panic(err)
	}
	daysUntil13th := service.calcDaysUntilDay13OfCurrentMonth(todayHijriDate)

	today := getStartOfDate(time.Now())
	return FastingDays{
		THIRTEENTH: addDays(today, int(daysUntil13th)),
		FOURTEENTH: addDays(today, int(daysUntil13th+1)),
		FIFTEENTH:  addDays(today, int(daysUntil13th+2)),
	}
}

func (service FastingService) calcDaysUntilDay13OfCurrentMonth(todaysHijriDate HijriDate) int64 {
	return 13 - todaysHijriDate.Day
}
