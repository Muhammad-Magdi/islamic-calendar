package astronomical

import (
	"testing"
	"time"
)

func TestGetJulianDate(t *testing.T) {
	now, _ := time.Parse(time.RFC822, "29 Oct 21 14:56 CAI")
	expectedJulianDate := 2459516.50000

	returnedJulianDate := GetJulianDate(now)
	if returnedJulianDate != float64(expectedJulianDate) {
		t.Errorf("Wrong julian date: expected = %f, found = %f", expectedJulianDate, returnedJulianDate)
	}
}
