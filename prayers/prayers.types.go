package prayers

import "time"

type PrayerTimesQueryParams struct {
	Long     float64   `json:"long" form:"long" binding:"required"`
	Lat      float64   `json:"lat" form:"lat" binding:"required"`
	Timezone int       `json:"timezone" form:"timezone" binding:"required"` // Can be deduced given the long,lat
	DateFrom time.Time `json:"date_from" form:"date_from" time_format:"2006-01-02" binding:"required"`
	DateTo   time.Time `json:"date_to" form:"date_to" time_format:"2006-01-02" binding:"required"`
}

type DayPrayerTimesResponse struct {
	Date string `json:"date"`

	Fajr       float64 `json:"fajr"`
	Ishraq     float64 `json:"ishraq"`
	Dhuhr      float64 `json:"dhuhr"`
	Asr        float64 `json:"asr"`
	Ghoroub    float64 `json:"ghoroub"`
	Maghrib    float64 `json:"maghrib"`
	Isha       float64 `json:"isha"`
	Nesfullail float64 `json:"nesfullail"`
}

type PrayerTimesResponse struct {
	Prayers []DayPrayerTimesResponse `json:"prayers"`
}
