package prayers

import (
	"github.com/muhammad-magdi/islamic-calendar/astronomical"
	"github.com/muhammad-magdi/islamic-calendar/dmath"
)

type PrayerTimesCalculator struct {
	stConfig        astronomical.Spacetime
	method          CalculationMethod
	defaultAngles   map[string]float64
	defaultTimes    map[string]float64
	astroCalculator astronomical.TimeCalculator
}

func NewPrayerTimesCalculator(method CalculationMethod, confing astronomical.Spacetime) PrayerTimesCalculator {
	timeCalculator := astronomical.NewTimeCalculator(confing)

	calculator := PrayerTimesCalculator{
		stConfig: confing,
		method:   method,

		defaultAngles: map[string]float64{
			DAY_TIME_ISHRAQ:  0.833,
			DAY_TIME_GHOROUB: 0.833,
		},
		defaultTimes: map[string]float64{
			DAY_TIME_IMSAK:   5,
			DAY_TIME_FAJR:    5,
			DAY_TIME_ISHRAQ:  6,
			DAY_TIME_DHUHR:   12,
			DAY_TIME_ASR:     13,
			DAY_TIME_GHOROUB: 18,
			DAY_TIME_MAGHRIB: 18,
			DAY_TIME_ISHA:    18,
		},
		astroCalculator: timeCalculator,
	}

	return calculator
}

func (calc PrayerTimesCalculator) GetPrayerTimes() map[string]float64 {

	dayPortions := calc.computeDayPortions()
	prayerTimes := calc.computePrayerTimes(dayPortions)
	prayerTimes = calc.fixTimezone(prayerTimes)
	// prayerTimes = adjustHighLats(prayerTimes)

	return prayerTimes
}

func (calc PrayerTimesCalculator) computeDayPortions() map[string]float64 {
	dayPortions := make(map[string]float64)
	for prayerName, defaultTime := range calc.defaultTimes {
		dayPortions[prayerName] = defaultTime / 24
	}

	return dayPortions
}

func (calc PrayerTimesCalculator) computePrayerTimes(dayPortions map[string]float64) map[string]float64 {
	prayerTimes := map[string]float64{}

	// prayerTimes[DAY_TIME_IMSAK] = calc.astroCalculator.GetSunAngleTime(dayPortions[DAY_TIME_IMSAK], calc.defaultAngles[DAY_TIME_IMSAK])
	if calc.method.FajrOffset.IsAngle() {
		prayerTimes[DAY_TIME_FAJR] = calc.astroCalculator.GetBySunAngle(dayPortions[DAY_TIME_FAJR], calc.method.FajrOffset.Angle())
	}
	prayerTimes[DAY_TIME_ISHRAQ] = calc.astroCalculator.GetBySunAngle(dayPortions[DAY_TIME_ISHRAQ], calc.defaultAngles[DAY_TIME_ISHRAQ])

	prayerTimes[DAY_TIME_DHUHR] = calc.astroCalculator.GetMidDayTime(dayPortions[DAY_TIME_DHUHR])

	prayerTimes[DAY_TIME_ASR] = calc.astroCalculator.GetByShadowRatio(dayPortions[DAY_TIME_ASR], float64(calc.method.AsrFactor))

	prayerTimes[DAY_TIME_GHOROUB] = calc.astroCalculator.GetBySunAngle(dayPortions[DAY_TIME_GHOROUB], calc.defaultAngles[DAY_TIME_GHOROUB])
	if calc.method.MaghribOffset.IsAngle() {
		prayerTimes[DAY_TIME_MAGHRIB] = calc.astroCalculator.GetBySunAngle(dayPortions[DAY_TIME_MAGHRIB], calc.method.MaghribOffset.Angle())
	} else {
		prayerTimes[DAY_TIME_MAGHRIB] = getShiftedTime(prayerTimes[calc.method.MaghribOffset.From], calc.method.MaghribOffset)
	}

	if calc.method.IshaOffset.IsAngle() {
		prayerTimes[DAY_TIME_ISHA] = calc.astroCalculator.GetBySunAngle(dayPortions[DAY_TIME_ISHA], calc.method.IshaOffset.Angle())
	} else {
		prayerTimes[DAY_TIME_ISHA] = getShiftedTime(prayerTimes[calc.method.IshaOffset.From], calc.method.IshaOffset)
	}

	// TODO: Fix bug, Refix hour
	prayerTimes[DAY_TIME_NESFULLAIL] = prayerTimes[DAY_TIME_GHOROUB] + dmath.FixHour(prayerTimes[DAY_TIME_ISHRAQ]-prayerTimes[DAY_TIME_GHOROUB])/2

	return prayerTimes
}

func getShiftedTime(value float64, shift MethodOffest) float64 {
	return value + shift.Value/60
}

func (calc PrayerTimesCalculator) fixTimezone(prayerTimes map[string]float64) map[string]float64 {
	for prayerName := range prayerTimes {
		prayerTimes[prayerName] += float64(calc.stConfig.Timezone) - calc.stConfig.Lng/15
	}
	return prayerTimes
}
