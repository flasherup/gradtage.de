package utils

import (
	"testing"
	"time"
)

func TestDaysCheck(t *testing.T) {
	type testValue struct {
		start time.Time
		end time.Time
		count int
		res bool
	}
	st := time.Now()

	testValues := []testValue{
		{
			start: st.Add(- time.Hour * 24),
			end: st,
			count: 1,
			res: true,
		},{
			start: st.Add(- time.Hour * 24 * 360),
			end: st,
			count: 260,
			res: true,
		},{
			start: st.Add(- time.Hour * 24 * 2),
			end: st,
			count: 3,
			res: false,
		},{
		start: st.Add(- time.Hour * 24 * 3),
		end: st,
		count: 2,
		res: true,
	},
	}

	for i,v := range testValues {
		res := DaysCheck(v.start, v.end, v.count)
		if res != v.res {
			t.Errorf("Test:%d DaysCheck(%s, %s, %d)", i,v.start.Format("2006-01-02"), v.end.Format("2006-01-02"), v.count)
		}
	}
}

func TestDaysDifference(t *testing.T) {
	type testValue struct {
		start time.Time
		end time.Time
		res int
	}
	st := time.Now()

	testValues := []testValue{
		{
			start: st.Add(- time.Hour * 24),
			end: st,
			res: 1,
		},{
			start: st.Add(- time.Hour * 24 * 360),
			end: st,
			res: 360,
		},{
			start: st.Add(- time.Hour * 24 * 2),
			end: st,
			res: 2,
		},{
			start: st.Add(- time.Hour * 24 * 3),
			end: st,
			res: 3,
		},
	}

	for i,v := range testValues {
		res := DaysDifference(v.start, v.end,)
		if res != v.res {
			t.Errorf("Test:%d DaysCheck(%s, %s) = %d, wants %d", i,v.start.Format("2006-01-02"), v.end.Format("2006-01-02"), res, v.res)
		}
	}
}
