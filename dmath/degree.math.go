package dmath

import "math"

func ToRadian(d float64) float64 {
	return d * math.Pi / 180.0
}

func ToDegree(r float64) float64 {
	return r * 180.0 / math.Pi
}

func Sin(d float64) float64 {
	return math.Sin(ToRadian(d))
}

func Cos(d float64) float64 {
	return math.Cos(ToRadian(d))
}

func Tan(d float64) float64 {
	return math.Tan(ToRadian(d))
}

func ASin(x float64) float64 {
	return ToDegree(math.Asin(x))
}

func ACos(x float64) float64 {
	return ToDegree(math.Acos(x))
}

func ATan(x float64) float64 {
	return ToDegree(math.Atan(x))
}

func ACot(x float64) float64 {
	return ToDegree(math.Atan(1 / x))
}

func ATan2(y, x float64) float64 {
	return ToDegree(math.Atan2(y, x))
}

func FixAngle(a float64) float64 {
	return mod(a, 360)
}

func FixHour(a float64) float64 {
	return mod(a, 24)
}

// mod returns a%b
// Using the division algorithm a = q*b + r
// Where q is floor(a/b), r is the remainder
// Then r = a - q*b
func mod(a float64, b float64) float64 {
	q := math.Floor(a / b)
	r := a - q*b
	if r < 0 {
		return b + r
	}
	return r
}
