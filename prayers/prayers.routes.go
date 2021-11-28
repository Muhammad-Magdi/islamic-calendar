package prayers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/muhammad-magdi/islamic-calendar/astronomical"
)

type PrayersRouter struct {
}

func NewPrayersRouter() PrayersRouter {
	return PrayersRouter{}
}

func (PrayersRouter) GetPrayerTimes(c *gin.Context) {
	layoutISO := "2006-01-02"

	params := PrayerTimesQueryParams{}
	err := c.BindQuery(&params)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"message": "TODO", "error": err.Error()})
		return
	}

	prayerTimes := PrayerTimesResponse{}
	for day := params.DateFrom; !day.After(params.DateTo); day = day.Add(24 * time.Hour) {
		calculator := NewPrayerTimesCalculator(astronomical.Spacetime{
			Lng:      params.Long,
			Lat:      params.Lat,
			Timezone: params.Timezone,
			Date:     day,
		})
		times := calculator.GetPrayerTimes()
		prayerTimes.Prayers = append(prayerTimes.Prayers, DayPrayerTimesResponse{
			Date:       day.Format(layoutISO),
			Fajr:       times[DAY_TIME_FAJR],
			Ishraq:     times[DAY_TIME_ISHRAQ],
			Dhuhr:      times[DAY_TIME_DHUHR],
			Asr:        times[DAY_TIME_ASR],
			Ghoroub:    times[DAY_TIME_GHOROUB],
			Maghrib:    times[DAY_TIME_MAGHRIB],
			Isha:       times[DAY_TIME_ISHA],
			Nesfullail: times[DAY_TIME_NESFULLAIL],
		})
	}

	c.JSON(http.StatusOK, prayerTimes)
}

func (PrayersRouter) GetCalculationMethods(c *gin.Context) {
	methods := GetCalculationMethodsMap()

	c.JSON(http.StatusOK, methods)
}
