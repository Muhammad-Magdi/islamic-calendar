package prayers

import (
	"github.com/muhammad-magdi/islamic-calendar/astronomical"
	"github.com/muhammad-magdi/islamic-calendar/dmath"
)

type Shift struct {
	value float64
	from  string
}

type PrayerTimesCalculator struct {
	stConfig        astronomical.Spacetime
	prayerShift     map[string]Shift // +/- mins of prayer shift
	defaultAngles   map[string]float64
	defaultTimes    map[string]float64
	astroCalculator astronomical.TimeCalculator
}

func NewPrayerTimesCalculator(confing astronomical.Spacetime) PrayerTimesCalculator {
	timeCalculator := astronomical.NewTimeCalculator(confing)

	calculator := PrayerTimesCalculator{
		stConfig: confing,
		prayerShift: map[string]Shift{
			DAY_TIME_IMSAK:   {-10.0, DAY_TIME_FAJR},
			DAY_TIME_MAGHRIB: {0, DAY_TIME_GHOROUB},
		},
		defaultAngles: map[string]float64{
			DAY_TIME_IMSAK:   10,
			DAY_TIME_FAJR:    19.5,
			DAY_TIME_ISHRAQ:  0.833,
			DAY_TIME_GHOROUB: 0.833,
			DAY_TIME_MAGHRIB: 0,
			DAY_TIME_ISHA:    17.5,
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
	prayerTimes = calc.applyMethodShifts(prayerTimes)
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

	// prayerTimes[DAY_TIME_IMSAK] = calc.astroCalculator.GetSunAngleTime(dayPortions[DAY_TIME_IMSAK], calc.defaultAngles[DAY_TIME_IMSAK], astronomical.DIRECTION_CCW)
	prayerTimes[DAY_TIME_FAJR] = calc.astroCalculator.GetSunAngleTime(dayPortions[DAY_TIME_FAJR], calc.defaultAngles[DAY_TIME_FAJR], astronomical.DIRECTION_CCW)
	prayerTimes[DAY_TIME_ISHRAQ] = calc.astroCalculator.GetSunAngleTime(dayPortions[DAY_TIME_ISHRAQ], calc.defaultAngles[DAY_TIME_ISHRAQ], astronomical.DIRECTION_CCW)

	prayerTimes[DAY_TIME_DHUHR] = calc.astroCalculator.GetMidDayTime(dayPortions[DAY_TIME_DHUHR])
	asrFactor := 1.0 // In standard method
	prayerTimes[DAY_TIME_ASR] = calc.astroCalculator.GetAsrTime(dayPortions[DAY_TIME_ASR], asrFactor)

	prayerTimes[DAY_TIME_GHOROUB] = calc.astroCalculator.GetSunAngleTime(dayPortions[DAY_TIME_GHOROUB], calc.defaultAngles[DAY_TIME_GHOROUB], astronomical.DIRECTION_CW)
	prayerTimes[DAY_TIME_MAGHRIB] = calc.astroCalculator.GetSunAngleTime(dayPortions[DAY_TIME_MAGHRIB], calc.defaultAngles[DAY_TIME_MAGHRIB], astronomical.DIRECTION_CW)
	prayerTimes[DAY_TIME_ISHA] = calc.astroCalculator.GetSunAngleTime(dayPortions[DAY_TIME_ISHA], calc.defaultAngles[DAY_TIME_ISHA], astronomical.DIRECTION_CW)

	prayerTimes[DAY_TIME_NESFULLAIL] = prayerTimes[DAY_TIME_GHOROUB] + dmath.FixHour(prayerTimes[DAY_TIME_ISHRAQ]-prayerTimes[DAY_TIME_GHOROUB])/2

	return prayerTimes
}

func (calc PrayerTimesCalculator) applyMethodShifts(prayerTimes map[string]float64) map[string]float64 {
	for prayerName, shift := range calc.prayerShift {
		if _, exists := prayerTimes[shift.from]; !exists {
			panic("TODO: some error!")
		}
		prayerTimes[prayerName] = prayerTimes[shift.from] + shift.value/60
	}
	return prayerTimes
}

func (calc PrayerTimesCalculator) fixTimezone(prayerTimes map[string]float64) map[string]float64 {
	for prayerName := range prayerTimes {
		prayerTimes[prayerName] += float64(calc.stConfig.Timezone) - calc.stConfig.Lng/15
	}
	return prayerTimes
}
