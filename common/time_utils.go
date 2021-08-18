package common

import (
	"github.com/zsefvlol/timezonemapper"
	"time"
)

func GetDatesFromNow(daysNumber int, timeTemplate string) (string, end string) {
	e := time.Now()
	s := e.AddDate(0,0, -daysNumber)

	return s.Format(timeTemplate), e.Format(timeTemplate)
}

func GetTimezoneFormLatLon(lat, lon float64) (string, error) {
	timezone := timezonemapper.LatLngToTimezoneString(lat,lon)
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return "", err
	}
	now := time.Now()
	return now.In(loc).Format("MST"),nil
}
