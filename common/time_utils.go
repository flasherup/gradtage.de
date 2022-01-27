package common

import (
	"github.com/zsefvlol/timezonemapper"
	"strings"
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

func ParseTimeByBreakdown(date string, breakdown string) (time.Time, error) {
	timeLayout := TimeLayoutDay
	if breakdown == BreakdownWeekly {
		timeLayout = TimeLayoutDay
	} else if breakdown == BreakdownWeeklyISO {
		timeLayout = TimeLayoutDay
	} else if breakdown == BreakdownMonthly {
		timeLayout = TimeLayoutMonth
	} else if breakdown == BreakdownYearly {
		timeLayout = TimeLayoutYear
	}
	return time.Parse(timeLayout, date)
}

func StrDayToWeekday(day string) time.Weekday {
	d := strings.ToLower(day)
	if d == Tuesday { return time.Tuesday }
	if d == Wednesday { return time.Wednesday }
	if d == Thursday { return time.Thursday }
	if d == Friday { return time.Friday }
	if d == Saturday { return time.Saturday }
	if d == Sunday { return time.Sunday }
	return time.Monday
}

func LeapYearDay(date time.Time) int {
	daysShift := 0
	if !IsLeapYear(date.Year()) && date.Month() > 2 {
		daysShift = 1
	}
	return date.YearDay() + daysShift
}

func YearWeekISO(date time.Time) int {
	y,w := date.ISOWeek()
	return y+w
}

func WeekISO(date time.Time) int {
	_,w := date.ISOWeek()
	return w
}

func Week(date time.Time, day time.Weekday) int {
	_,w := CustomWeek(date, day)
	return w
}

func CustomWeek(date time.Time, day time.Weekday) (int, int) {
	delta := int((day + 6)%7 + 1) - 1
	shift :=  12 * delta
	d := time.Date(date.Year()-shift, date.Month(), date.Day(), date.Hour(), date.Minute(), date.Second(), date.Nanosecond(), date.Location())
	y,w := d.ISOWeek()
	return y + shift, w
}


func IsLeapYear(year int) bool {
	leapFlag := false
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				leapFlag = true
			} else {
				leapFlag = false
			}
		} else {
			leapFlag = true
		}
	} else {
		leapFlag = false
	}
	return leapFlag
}


