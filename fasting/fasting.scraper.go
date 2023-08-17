package fasting

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

const DAR_AL_IFTA_WEBSITE = "https://www.dar-alifta.org/Foreign/default.aspx"

type FastingScraper struct {
	scraper *colly.Collector
}

func NewFastingScraper() FastingScraper {
	return FastingScraper{scraper: colly.NewCollector()}
}

func (fastingScraper FastingScraper) GetTodaysHijriDate() (HijriDate, error) {
	return fastingScraper.GetDarAlIftaCurrentDate()
}

func (fastingScraper FastingScraper) GetDarAlIftaCurrentDate() (HijriDate, error) {
	ch := make(chan string)
	errCh := make(chan error)

	fastingScraper.scraper.OnHTML(`#header`, func(el *colly.HTMLElement) {
		el.ForEach("div.he-text", func(i int, el *colly.HTMLElement) {
			if strings.Contains(el.Text, "Hijri") {
				ch <- el.Text
			} else {
				errCh <- errors.New("failed to find Hijri date div: " + el.Text)
			}
		})
	})

	go func() {
		err := fastingScraper.scraper.Visit(DAR_AL_IFTA_WEBSITE)
		if err != nil {
			errCh <- errors.New("failed to visit website: " + err.Error())
		}

		close(ch)
	}()

	select {
	case err := <-errCh:
		return HijriDate{}, err

	case hijriDateDivContent := <-ch:
		if hijriDateDivContent == "" {
			return HijriDate{}, errors.New("failed to scrap for Hijri date: empty div content")
		}

		splitDivContent := strings.Split(hijriDateDivContent, "\"")
		if len(splitDivContent) != 3 {
			return HijriDate{}, errors.New("failed to scrap for Hijri date: unexpected div content: " + hijriDateDivContent)
		}

		hijriDate, err := parseHijriDate(splitDivContent[1])
		if err != nil {
			return HijriDate{}, errors.New("failed to scrap for Hijri date: " + err.Error())
		}

		return hijriDate, nil
	}
}

func parseHijriDate(str string) (HijriDate, error) {
	splitDate := strings.Split(str, " ")
	day, err := strconv.ParseInt(splitDate[0], 10, 64)
	if err != nil {
		return HijriDate{}, errors.New("failed to parse hijri day: " + err.Error())
	}

	strMonth := splitDate[1]
	month, err := NewHijriMonth(strMonth)
	if err != nil {
		return HijriDate{}, errors.New("failed to parse hijri month: " + err.Error())
	}

	year, err := strconv.ParseInt(splitDate[2], 10, 64)
	if err != nil {
		return HijriDate{}, errors.New("failed to parse hijri year: " + err.Error())
	}

	return NewHijriDate(day, month, year), nil
}
