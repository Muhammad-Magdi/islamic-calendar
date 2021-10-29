package astronomical

import (
	"math"
	"time"
)

func GetJulianDate(date time.Time) float64 {
	year, month, day := date.Year(), date.Month(), date.Day()
	if month <= 2 {
		year--
		month += 12
	}
	A := math.Floor(float64(year) / 100)
	B := 2 - A + math.Floor(A/4)

	JD := math.Floor(365.25*(float64(year)+4716)) + math.Floor(30.6001*float64(month+1)) + float64(day) + B - 1524.5
	return JD
}
