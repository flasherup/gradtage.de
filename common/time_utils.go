package common

import "time"

func GetDatesFromNow(daysNumber int, timeTemplate string) (string, end string) {
	e := time.Now()
	s := e.AddDate(0,0, -daysNumber)

	return s.Format(timeTemplate), e.Format(timeTemplate)
}
