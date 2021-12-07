package prayers

import (
	"math"
	"testing"
	"time"

	"github.com/muhammad-magdi/islamic-calendar/astronomical"
)

const EPS = 1e-6

func getTestSetup() (PrayerTimesCalculator, map[string]float64) {
	cairoLng, cairoLat, cairoTimezone := 31.24967, 30.06263, 2
	config := astronomical.Spacetime{
		Lng:      cairoLng,
		Lat:      cairoLat,
		Timezone: cairoTimezone,
		Date:     time.Date(2021, 10, 31, 0, 0, 0, 0, time.FixedZone("CAI", 2*60*60)),
	}

	calculator := NewPrayerTimesCalculator(
		GetCalculationMethod("EGSA"),
		config,
	)

	expectedTimes := map[string]float64{
		DAY_TIME_FAJR:    04.670996,
		DAY_TIME_DHUHR:   11.643454,
		DAY_TIME_ASR:     14.759843,
		DAY_TIME_MAGHRIB: 17.143815,
		DAY_TIME_ISHA:    18.454795,
	}

	return calculator, expectedTimes
}

func TestComputePrayerTimes(t *testing.T) {
	calculator, expectedTimes := getTestSetup()

	times := calculator.GetPrayerTimes()
	for prayerName, expectedTime := range expectedTimes {
		if math.Abs(expectedTime-times[prayerName]) > EPS {
			t.Errorf("Error in %s: expected at %f, found at %f", prayerName, expectedTime, times[prayerName])
		}
	}

}

func BenchmarkComputePrayerTimes(b *testing.B) {
	calculator, expectedTimes := getTestSetup()

	for i := 0; i < b.N; i++ {
		times := calculator.GetPrayerTimes()
		for prayerName, expectedTime := range expectedTimes {
			if math.Abs(expectedTime-times[prayerName]) > EPS {
				b.Errorf("Error in %s: expected at %f, found at %f", prayerName, expectedTime, times[prayerName])
			}
		}
	}

}
