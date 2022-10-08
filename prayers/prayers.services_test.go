package prayers

import (
	"math"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/muhammad-magdi/islamic-calendar/astronomical"
)

const EPS = 1

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
		day   string
		times map[string]string
	}

	tests := []Test{
		{day: "Oct 02, 2022", times: map[string]string{"fajr": "4:22", "dhuhr": "11:44", "asr": "15:07", "maghrib": "17:40", "isha": "18:57"}},
		{day: "Oct 03, 2022", times: map[string]string{"fajr": "4:23", "dhuhr": "11:44", "asr": "15:06", "maghrib": "17:38", "isha": "18:55"}},
		{day: "Oct 04, 2022", times: map[string]string{"fajr": "4:23", "dhuhr": "11:44", "asr": "15:06", "maghrib": "17:37", "isha": "18:54"}},
		{day: "Oct 05, 2022", times: map[string]string{"fajr": "4:24", "dhuhr": "11:43", "asr": "15:04", "maghrib": "17:35", "isha": "18:53"}},
		{day: "Oct 06, 2022", times: map[string]string{"fajr": "4:25", "dhuhr": "11:43", "asr": "15:04", "maghrib": "17:35", "isha": "18:52"}},
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

		for prayerName, prayerTime := range tc.times {
			fPrayerTime := convertHH_MMToFloat(prayerTime)
			if math.Abs(res[prayerName]-fPrayerTime) > EPS {
				t.Fatalf("Test failed at day %v. Expected %s prayer at: %f found at %f", date, prayerName, fPrayerTime, res[prayerName])
			}
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

func convertHH_MMToFloat(tm string) float64 {
	components := strings.Split(tm, ":")
	hh, mm := components[0], components[1]

	fHH, _ := strconv.ParseFloat(hh, 64)
	fMM, _ := strconv.ParseFloat(mm, 64)

	return fHH + fMM/60
}
