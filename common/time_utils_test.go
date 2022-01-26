package common

import (
	"testing"
	"time"
)

/*func TestWeek(t *testing.T) {
	type test struct {
		date string
		day  time.Weekday
		res  int
	}

	values := []test{
		{
			"2020-01-01",
			time.Sunday,
			1,
		},{
			"2019-12-31",
			time.Sunday,
			1,
		}, {
			"2015-01-01",
			time.Sunday,
			12,
		}, {
			"2014-12-31",
			time.Sunday,
			12,
		},
	}

	for _, v := range values {
		d, e := time.Parse(TimeLayoutWBH, v.date)
		if e != nil {
			t.Errorf("week time parse error: %s", e.Error())
		}

		res := Week(d, v.day)

		if res != v.res {
			t.Errorf("week(%s, %d) = %d; want %d", v.date, v.day, res, v.res)
		}
	}
}*/

func TestCustomWeek(t *testing.T) {
	type test struct {
		date string
		day  time.Weekday
		week int
		year int
	}

	values := []test{
		{
			"2020-01-01",
			time.Monday,
			1,
			2020,
		}, {
			"2020-01-01",
			time.Tuesday,
			1,
			2020,
		}, {
			"2020-01-01",
			time.Wednesday,
			1,
			2020,
		}, {
			"2020-01-01",
			time.Thursday,
			52,
			2019,
		}, {
			"2020-01-01",
			time.Friday,
			52,
			2019,
		}, {
			"2020-01-01",
			time.Saturday,
			53,
			2019,
		}, {
			"2020-01-01",
			time.Sunday,
			1,
			2020,
		}, {
			"2019-12-31",
			time.Sunday,
			1,
			2020,
		}, {
			"2015-01-01",
			time.Sunday,
			53,
			2014,
		}, {
			"2014-12-31",
			time.Sunday,
			53,
			2014,
		},{
			"2015-01-01",
			time.Thursday,
			1,
			2015,
		}, {
			"2014-12-31",
			time.Thursday,
			52,
			2014,
		},
	}

	for _, v := range values {
		d, e := time.Parse(TimeLayoutWBH, v.date)
		if e != nil {
			t.Errorf("week time parse error: %s", e.Error())
		}
		y, w := CustomWeek(d, v.day)

		if w != v.week || y != v.year {
			t.Errorf("week(%s, %d) = year:%d, week:%d; want year:%d, week:%d", v.date, v.day, y, w, v.year, v.week)
		}
	}
}
