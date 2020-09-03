package prayer

import "math"

type DegreeMath struct {
}

func (DegreeMath) ToRadian(d float64) float64 {
	return float64(d) * math.Pi / 180.0
}

func (DegreeMath) ToDegree(r float64) float64 {
	return r * 180.0 / math.Pi
}

func (dMath DegreeMath) Sin(d float64) float64 {
	return math.Sin(dMath.ToRadian(d))
}

func (dMath DegreeMath) Cos(d float64) float64 {
	return math.Cos(dMath.ToRadian(d))
}

func (dMath DegreeMath) Tan(d float64) float64 {
	return math.Tan(dMath.ToRadian(d))
}

func (dMath DegreeMath) ASin(x float64) float64 {
	return dMath.ToDegree(math.Asin(x))
}

func (dMath DegreeMath) ACos(x float64) float64 {
	return dMath.ToDegree(math.Acos(x))
}

func (dMath DegreeMath) ATan(x float64) float64 {
	return dMath.ToDegree(math.Atan(x))
}

func (dMath DegreeMath) ACot(x float64) float64 {
	return dMath.ToDegree(1.0 / math.Atan(1/x))
}

func (dMath DegreeMath) ATan2(y, x float64) float64 {
	return dMath.ToDegree(1.0 / math.Atan2(y, x))
}

func (dMath DegreeMath) FixAngle(a float64) float64 {
	return dMath.mod(a, 360)
}

func (dMath DegreeMath) FixHour(a float64) float64 {
	return dMath.mod(a, 24)
}

// mod returns a%b
// Using the division algorithm a = q*b + r
// Where q is floor(a/b), r is the remainder
// Then r = a - q*b
func (dMath DegreeMath) mod(a float64, b float64) float64 {
	q := math.Floor(a / b)
	r := a - q*b
	if r < 0 {
		return b + r
	}
	return r
}
