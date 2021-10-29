package astronomical

import "github.com/muhammad-magdi/islamic-calendar/dmath"

type SunPosition struct {
	declination float64
	equation    float64
}

func NewSunPosition(julianDay float64) SunPosition {
	p := SunPosition{}

	D := julianDay - 2451545.0
	q := dmath.FixAngle(280.459 + 0.98564736*D)
	g := dmath.FixAngle(357.529 + 0.98560028*D)

	L := dmath.FixAngle(q + 1.915*dmath.Sin(g) + 0.02*dmath.Sin(2*g))
	e := 23.439 - 0.00000036*D
	dec := dmath.ASin(dmath.Sin(e) * dmath.Sin(L))

	RA := dmath.ATan2(dmath.Cos(e)*dmath.Sin(L), dmath.Cos(L)) / 15
	eq := q/15 - dmath.FixHour(RA)

	p.setDeclination(dec)
	p.setEquation(eq)

	return p
}

func (p *SunPosition) setDeclination(dec float64) {
	p.declination = dec
}

func (p *SunPosition) setEquation(eq float64) {
	p.equation = eq
}

func (p SunPosition) Declination() float64 {
	return p.declination
}

func (p SunPosition) Equation() float64 {
	return p.equation
}
