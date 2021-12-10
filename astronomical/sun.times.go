package astronomical

import (
	"math"
	"time"

	"github.com/muhammad-magdi/islamic-calendar/dmath"
)

type Spacetime struct {
	Lng      float64
	Lat      float64
	Timezone int
	Date     time.Time
}

type TimeCalculator struct {
	st    Spacetime
	jDate float64
}

func NewTimeCalculator(st Spacetime) TimeCalculator {
	timeCalculator := TimeCalculator{
		st:    st,
		jDate: GetJulianDate(st.Date) - st.Lng/(15.0*24.0),
	}

	return timeCalculator
}

// GetMidDayTime returns the time of the mid day i.e. when the sun becomes overhead
func (calc TimeCalculator) GetMidDayTime(dayPortion float64) float64 {
	jTime := calc.jDate + dayPortion

	eqt := NewSunPosition(jTime).Equation()
	midDay := dmath.FixHour(12 - eqt)

	return midDay
}

// GetBySunAngle returns the time of the day given the angle between the sun and the horizon
//
// It does NOT work when the angle is 90. Use GetMidDayTime instead
func (calc TimeCalculator) GetBySunAngle(dayPortion float64, angle float64, direction string) float64 {
	jTime := calc.jDate + dayPortion

	decl := NewSunPosition(jTime).Declination()

	t := (1.0 / 15.0) * dmath.ACos((-dmath.Sin(angle)-dmath.Sin(decl)*dmath.Sin(calc.st.Lat))/(dmath.Cos(decl)*dmath.Cos(calc.st.Lat)))

	midDay := calc.GetMidDayTime(dayPortion)
	if direction == DIRECTION_CCW {
		return midDay - t
	}

	return midDay + t
}

// GetByShadowRatio returns the time of the day given the ratio between a body and its shadow length
func (calc TimeCalculator) GetByShadowRatio(dayPortion float64, factor float64) float64 {
	jTime := calc.jDate + dayPortion

	decl := NewSunPosition(jTime).Declination()
	angle := -dmath.ACot(factor + dmath.Tan(math.Abs(calc.st.Lat-decl)))

	sunAngle := calc.GetBySunAngle(dayPortion, angle, DIRECTION_CW)

	return sunAngle
}
