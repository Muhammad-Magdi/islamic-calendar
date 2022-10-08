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
	type Test struct {
		day     string
		fajr    float64
		dhuhr   float64
		asr     float64
		maghrib float64
		isha    float64
	}

	// These times are taken from www.islamicfinder.org
	tests := []Test{
		{day: "Jan 01, 2021", fajr: 5.31666666666667, dhuhr: 11.9833333333333, asr: 14.7833333333333, maghrib: 17.1, isha: 18.4833333333333},
		{day: "Jan 02, 2021", fajr: 5.31666666666667, dhuhr: 11.9833333333333, asr: 14.8, maghrib: 17.1166666666667, isha: 18.4833333333333},
		{day: "Jan 03, 2021", fajr: 5.31666666666667, dhuhr: 11.9833333333333, asr: 14.8166666666667, maghrib: 17.1166666666667, isha: 18.5},
		{day: "Jan 04, 2021", fajr: 5.31666666666667, dhuhr: 12, asr: 14.8166666666667, maghrib: 17.1333333333333, isha: 18.5166666666667},
		{day: "Jan 05, 2021", fajr: 5.33333333333333, dhuhr: 12, asr: 14.8333333333333, maghrib: 17.15, isha: 18.5166666666667},
		{day: "Jan 06, 2021", fajr: 5.33333333333333, dhuhr: 12.0166666666667, asr: 14.85, maghrib: 17.1666666666667, isha: 18.5333333333333},
		{day: "Jan 07, 2021", fajr: 5.33333333333333, dhuhr: 12.0166666666667, asr: 14.8666666666667, maghrib: 17.1666666666667, isha: 18.55},
		{day: "Jan 08, 2021", fajr: 5.33333333333333, dhuhr: 12.0333333333333, asr: 14.8666666666667, maghrib: 17.1833333333333, isha: 18.5666666666667},
		{day: "Jan 09, 2021", fajr: 5.33333333333333, dhuhr: 12.0333333333333, asr: 14.8833333333333, maghrib: 17.2, isha: 18.5666666666667},
		{day: "Jan 10, 2021", fajr: 5.35, dhuhr: 12.05, asr: 14.9, maghrib: 17.2166666666667, isha: 18.5833333333333},
		{day: "Jan 11, 2021", fajr: 5.35, dhuhr: 12.05, asr: 14.9166666666667, maghrib: 17.2333333333333, isha: 18.6},
		{day: "Jan 12, 2021", fajr: 5.35, dhuhr: 12.05, asr: 14.9166666666667, maghrib: 17.2333333333333, isha: 18.6},
		{day: "Jan 13, 2021", fajr: 5.35, dhuhr: 12.0666666666667, asr: 14.9333333333333, maghrib: 17.25, isha: 18.6166666666667},
		{day: "Jan 14, 2021", fajr: 5.35, dhuhr: 12.0666666666667, asr: 14.95, maghrib: 17.2666666666667, isha: 18.6333333333333},
		{day: "Jan 15, 2021", fajr: 5.35, dhuhr: 12.0666666666667, asr: 14.9666666666667, maghrib: 17.2833333333333, isha: 18.65},
		{day: "Jan 16, 2021", fajr: 5.35, dhuhr: 12.0833333333333, asr: 14.9666666666667, maghrib: 17.3, isha: 18.65},
		{day: "Jan 17, 2021", fajr: 5.35, dhuhr: 12.0833333333333, asr: 14.9833333333333, maghrib: 17.3166666666667, isha: 18.6666666666667},
		{day: "Jan 18, 2021", fajr: 5.35, dhuhr: 12.1, asr: 15, maghrib: 17.3166666666667, isha: 18.6833333333333},
		{day: "Jan 19, 2021", fajr: 5.33333333333333, dhuhr: 12.1, asr: 15.0166666666667, maghrib: 17.3333333333333, isha: 18.7},
		{day: "Jan 20, 2021", fajr: 5.33333333333333, dhuhr: 12.1, asr: 15.0166666666667, maghrib: 17.35, isha: 18.7},
		{day: "Jan 21, 2021", fajr: 5.33333333333333, dhuhr: 12.1, asr: 15.0333333333333, maghrib: 17.3666666666667, isha: 18.7166666666667},
		{day: "Jan 22, 2021", fajr: 5.33333333333333, dhuhr: 12.1166666666667, asr: 15.05, maghrib: 17.3833333333333, isha: 18.7333333333333},
		{day: "Jan 23, 2021", fajr: 5.33333333333333, dhuhr: 12.1166666666667, asr: 15.0666666666667, maghrib: 17.4, isha: 18.75},
		{day: "Jan 24, 2021", fajr: 5.33333333333333, dhuhr: 12.1166666666667, asr: 15.0666666666667, maghrib: 17.4166666666667, isha: 18.75},
		{day: "Jan 25, 2021", fajr: 5.31666666666667, dhuhr: 12.1166666666667, asr: 15.0833333333333, maghrib: 17.4333333333333, isha: 18.7666666666667},
		{day: "Jan 26, 2021", fajr: 5.31666666666667, dhuhr: 12.1333333333333, asr: 15.1, maghrib: 17.4333333333333, isha: 18.7833333333333},
		{day: "Jan 27, 2021", fajr: 5.31666666666667, dhuhr: 12.1333333333333, asr: 15.1166666666667, maghrib: 17.45, isha: 18.8},
		{day: "Jan 28, 2021", fajr: 5.3, dhuhr: 12.1333333333333, asr: 15.1166666666667, maghrib: 17.4666666666667, isha: 18.8166666666667},
		{day: "Jan 29, 2021", fajr: 5.3, dhuhr: 12.1333333333333, asr: 15.1333333333333, maghrib: 17.4833333333333, isha: 18.8166666666667},
		{day: "Jan 30, 2021", fajr: 5.3, dhuhr: 12.15, asr: 15.15, maghrib: 17.5, isha: 18.8333333333333},
		{day: "Jan 31, 2021", fajr: 5.28333333333333, dhuhr: 12.15, asr: 15.1666666666667, maghrib: 17.5166666666667, isha: 18.85},
	}

	cairoLng, cairoLat, cairoTimezone := 31.24967, 30.06263, 2

	layout := "Jan 02, 2006"
	for _, tc := range tests {
		date, _ := time.Parse(layout, tc.day)
		config := astronomical.Spacetime{
			Lng:      cairoLng,
			Lat:      cairoLat,
			Timezone: cairoTimezone,
			Date:     date,
		}

		calculator := NewPrayerTimesCalculator(
			GetCalculationMethod("EGSA"),
			config,
		)

		res := calculator.GetPrayerTimes()

		// t.Fatalf("%f %f\n", getMinutes(res[DAY_TIME_DHUHR]), res[DAY_TIME_DHUHR])

		if math.Abs(nearestMinute(res[DAY_TIME_FAJR])-tc.fajr) > EPS {
			t.Fatalf("Test failed at day %v. Expected DHUHR prayer at: %f found at %f", date, tc.fajr, nearestMinute(res[DAY_TIME_FAJR]))
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

func nearestMinute(tm float64) float64 {
	return math.Round(tm*60) / 60
}
