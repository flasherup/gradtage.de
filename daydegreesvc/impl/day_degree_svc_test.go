package impl

import (
	"github.com/flasherup/gradtage.de/common"
	"testing"
	"time"
)

func TestGetDates(t *testing.T) {
	type test struct {
		initial   time.Time
		years     int
		breakdown string
		start     string
		end       string
		err       error
	}

	initial := time.Date(2000, 3, 5, 1, 1, 1, 1, time.Local)

	tests := map[string]test{
		"daily": {
			initial:   initial,
			years:     2,
			breakdown: common.BreakdownDaily,
			start:     time.Date(1998, 3, 4, 0, 0, 0, 0, time.Local).Format(common.TimeLayoutWBH),
			end:       time.Date(2000, 3, 4, 0, 0, 0, 0, time.Local).Format(common.TimeLayoutWBH),
			err:       nil,
		},
		"weeklyISO": {
			initial:   initial,
			years:     2,
			breakdown: common.BreakdownWeeklyISO,
			start:     time.Date(1998, 3, 1, 0, 0, 0, 0, time.Local).Format(common.TimeLayoutWBH),
			end:       time.Date(2000, 2, 29, 0, 0, 0, 0, time.Local).Format(common.TimeLayoutWBH),
			err:       nil,
		},
		"weekly": {
			initial:   initial,
			years:     2,
			breakdown: common.BreakdownWeekly,
			start:     time.Date(1998, 3, 1, 0, 0, 0, 0, time.Local).Format(common.TimeLayoutWBH),
			end:       time.Date(2000, 2, 29, 0, 0, 0, 0, time.Local).Format(common.TimeLayoutWBH),
			err:       nil,
		},
		"monthly": {
			initial:   initial,
			years:     2,
			breakdown: common.BreakdownMonthly,
			start:     time.Date(1998, 2, 1, 0, 0, 0, 0, time.Local).Format(common.TimeLayoutWBH),
			end:       time.Date(2000, 2, 1, 0, 0, 0, 0, time.Local).Format(common.TimeLayoutWBH),
			err:       nil,
		},
		"yearly": {
			initial:   initial,
			years:     2,
			breakdown: common.BreakdownYearly,
			start:     time.Date(1998, 1, 1, 0, 0, 0, 0, time.Local).Format(common.TimeLayoutWBH),
			end:       time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local).Format(common.TimeLayoutWBH),
			err:       nil,
		},
	}

	for k, v := range tests {
		start, end, err := getDates(v.initial, v.years, v.breakdown)
		if  start != v.start || end != v.end || err != v.err {
			t.Errorf("Test: %s fail, getEndDate(%s, %d, %s) = %s, %s, %v; want %s, %s, %v",
				k,
				v.initial.Format(common.TimeLayoutWBH),
				v.years,
				v.breakdown,
				start,
				end,
				err,
				v.start,
				v.end,
				v.err,
			)
		}
	}
}

func TestGetEndDate(t *testing.T) {
	type test struct {
		initial    time.Time
		breakdown  string
		resultDate time.Time
	}

	initial := time.Date(2000, 2, 2, 0, 0, 0, 0, time.Local)

	tests := map[string]test{
		"daily": {
			initial:    initial,
			breakdown:  common.BreakdownDaily,
			resultDate: time.Date(2000, 2, 1, 0, 0, 0, 0, time.Local),
		},
		"weeklyISO": {
			initial:    initial,
			breakdown:  common.BreakdownWeeklyISO,
			resultDate: time.Date(2000, 1, 28, 0, 0, 0, 0, time.Local),
		},
		"weekly": {
			initial:    initial,
			breakdown:  common.BreakdownWeekly,
			resultDate: time.Date(2000, 1, 28, 0, 0, 0, 0, time.Local),
		},
		"monthly": {
			initial:    initial,
			breakdown:  common.BreakdownMonthly,
			resultDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
		},
		"yearly": {
			initial:    initial,
			breakdown:  common.BreakdownYearly,
			resultDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
		},
	}

	for k, v := range tests {
		res := getEndDate(v.initial, v.breakdown)
		if res != v.resultDate {
			t.Errorf("Test: %s fail, getEndDate(%s, %s) = %s; want %s",
				k,
				v.initial.Format(common.TimeLayout),
				v.breakdown,
				res.Format(common.TimeLayout),
				v.resultDate.Format(common.TimeLayout))
		}
	}
}
