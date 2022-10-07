package astronomical

import (
	"testing"
	"time"
)

func TestGetMidDayTime(t *testing.T) {

	day := time.Date(2021, 1, 2, 0, 0, 0, 0, time.FixedZone("CAI", 2*60*60))
	for d := 1; d <= 365; d++ {
		c := NewTimeCalculator(Spacetime{
			Lat:      30.06263,
			Lng:      31.24967,
			Timezone: 2,
			Date:     day,
		})

		// if c.GetByShadowRatio(0.5, 0) != c.GetBySunAngle(0.5, 90, "CCW") {
		// 	t.Fatalf("%v %f != %f", day, c.GetByShadowRatio(0.5, 0), c.GetBySunAngle(0.5, 90, "CCW"))
		// }

		if c.GetByShadowRatio(0.5, 0) != c.GetMidDayTime(0.5) {
			t.Fatalf("%v %f != %f", day, c.GetByShadowRatio(0.5, 0), c.GetMidDayTime(0.5))
		}
		day = day.Add(24 * time.Hour)
	}
}
