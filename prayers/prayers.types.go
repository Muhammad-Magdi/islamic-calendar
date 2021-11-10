package prayers

import "time"

type PrayerTimesQueryParams struct {
	Long     float64   `form:"long" binding:"required"`
	Lat      float64   `form:"lat" binding:"required"`
	Timezone int       `form:"timezone" binding:"required"` // Can be deduced given the long,lat
	DateFrom time.Time `form:"date_from" time_format:"2006-01-02" binding:"required"`
	DateTo   time.Time `form:"date_to" time_format:"2006-01-02" binding:"required"`
}
