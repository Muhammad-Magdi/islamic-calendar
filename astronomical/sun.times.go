package astronomical

import (
	"math"
	"time"

	"github.com/muhammad-magdi/islamic-calendar/dmath"
)

type Spacetime struct {
	Lng      float64
	Lat      float64
	Timezone float64
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

func (calc TimeCalculator) GetMidDayTime(dayPortion float64) float64 {
	jTime := calc.jDate + dayPortion

	eqt := NewSunPosition(jTime).Equation()
	midDay := dmath.FixHour(12 - eqt)

	return midDay
}

func (calc TimeCalculator) GetSunAngleTime(dayPortion float64, angle float64, direction string) float64 {
	jTime := calc.jDate + dayPortion

	decl := NewSunPosition(jTime).Declination()

	t := (1.0 / 15.0) * dmath.ACos((-dmath.Sin(angle)-dmath.Sin(decl)*dmath.Sin(calc.st.Lat))/(dmath.Cos(decl)*dmath.Cos(calc.st.Lat)))

	midDay := calc.GetMidDayTime(dayPortion)
	if direction == DIRECTION_CCW {
		return midDay - t
	}

	return midDay + t
}

func (calc TimeCalculator) GetAsrTime(dayPortion float64, factor float64) float64 {
	jTime := calc.jDate + dayPortion

	decl := NewSunPosition(jTime).Declination()
	angle := -dmath.ACot(factor + dmath.Tan(math.Abs(calc.st.Lat-decl)))

	sunAngle := calc.GetSunAngleTime(angle, dayPortion, DIRECTION_CW)

	return sunAngle
}
