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

// Query parameters:
// 1. lat, lng
// 2. timezone
// 3. from, to Date	YYYY-MM-DD
func (PrayersRouter) GetPrayerTimes(c *gin.Context) {
	layoutISO := "2006-01-02"

	params := PrayerTimesQueryParams{}
	err := c.BindQuery(&params)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"message": "TODO", "error": err.Error()})
		return
	}

	timesMap := make(map[string]map[string]float64)
	for day := params.DateFrom; !day.After(params.DateTo); day = day.Add(24 * time.Hour) {
		calculator := NewPrayerTimesCalculator(astronomical.Spacetime{
			Lng:      params.Long,
			Lat:      params.Lat,
			Timezone: params.Timezone,
			Date:     day,
		})
		times := calculator.GetPrayerTimes()
		timesMap[day.Format(layoutISO)] = times
	}

	c.JSON(http.StatusOK, timesMap)
}
