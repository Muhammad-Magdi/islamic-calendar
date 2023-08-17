package fasting

import (
	"errors"
	"time"
)

type FastingDaysQueryParams struct {
	Long float64 `json:"long" form:"long" binding:"required"`
	Lat  float64 `json:"lat" form:"lat" binding:"required"`
	// Country  string    // TODO: Replace this type with an enum or so
	DateFrom time.Time `json:"date_from" form:"date_from" time_format:"2006-01-02" binding:"required"`
	DateTo   time.Time `json:"date_to" form:"date_to" time_format:"2006-01-02" binding:"required"`
}

type HijriMonth int

type FastingDay string

const (
	THIRTEENTH FastingDay = "13"
	FOURTEENTH            = "14"
	FIFTEENTH             = "15"
)

type FastingDays map[FastingDay]time.Time

const (
	UNKNOWN HijriMonth = iota
	MOHARRAM
	SAFAR
	RABIE_ALAWAL
	RABIE_ALAKHAR
	JUMADA_ALAWAL
	JUMADA_ALAKHAR
	RAJAB
	SHAABAN
	RAMADAN
	SHAWWAL
	THULQEADA
	THULHEJJA
)

func NewHijriMonth(str string) (HijriMonth, error) {
	switch str {
	case "Muharram": // Confirmed
		return MOHARRAM, nil
	case "Safar": // Confirmed
		return SAFAR, nil
	case "Rabi' alAwwal":
		return RABIE_ALAWAL, nil
	case "Rabi' alAkhar":
		return RABIE_ALAKHAR, nil
	case "Jumada alAwwal":
		return JUMADA_ALAWAL, nil
	case "Jumada alAkhar":
		return JUMADA_ALAKHAR, nil
	case "Rajab":
		return RAJAB, nil
	case "Shaaban":
		return SHAABAN, nil
	case "Ramadan":
		return RAMADAN, nil
	case "Shawwal":
		return SHAWWAL, nil
	case "Thulqeada":
		return THULQEADA, nil
	case "Thulhejja":
		return THULHEJJA, nil
	default:
		return UNKNOWN, errors.New("unknown month: " + str)
	}
}

type HijriDate struct {
	Day   int64
	Month HijriMonth
	Year  int64
}

func NewHijriDate(day int64, month HijriMonth, year int64) HijriDate {
	return HijriDate{Day: day, Month: month, Year: year}
}
