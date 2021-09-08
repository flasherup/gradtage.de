package common

import (
	"testing"
	"time"
)

func TestIsTheSamePeriod(t *testing.T) {
	type values struct {
		latest string
		current string
		period string
		res bool
	}
	testValues := []values {
		{"2020-01-02T15:04:05Z","2020-01-03T20:04:05Z", BreakdownDaily,false},
		{"2020-01-02T15:04:05Z","2020-01-02T20:04:05Z", BreakdownDaily,true},
		{"2020-01-02T15:04:05Z","2020-02-02T15:04:05Z", BreakdownMonthly,false},
		{"2020-02-02T15:04:05Z","2020-02-02T15:04:05Z", BreakdownMonthly,true},
		{"2020-01-02T15:04:05Z","2021-02-02T15:04:05Z", BreakdownYearly,false},
		{"2020-02-02T15:04:05Z","2020-02-02T15:04:05Z", BreakdownYearly,true},
	}

	for _,v := range testValues{
		last,err := time.Parse(TimeLayout, v.latest)
		if err != nil {
			t.Error(err)
		}

		current, err := time.Parse(TimeLayout, v.current)
		if err != nil {
			t.Error(err)
		}

		res := isTheSamePeriod(last, current, v.period)
		if res != v.res {
			t.Errorf(
				"isTheSamePeriod(%s, %s, %s) = %t; want %t",
				v.latest,
				v.current,
				v.period,
				res,
				v.res,
				)
		}
	}
}

func TestCalculateHDD(t *testing.T) {
	baseHDD := 10.0
	testValues := []float64{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 15, 16, 17, 18, 19, 20,
	}

	testAnswers := []float64{
		9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}

	for i,v := range testValues {
		res := calculateHDD(baseHDD, v)
		if res != testAnswers[i] {
			t.Errorf("calculateHDD(%g, %g) = %g; want %g",baseHDD, v, res, testAnswers[i])
		}
	}
}

func TestCalculateDD(t *testing.T) {
	baseHDD := 10.0
	baseDD := 15.0
	testValues := []float64{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 15, 16, 17, 18, 19, 20,
	}

	testAnswers := []float64{
		14, 13, 12, 11, 10, 9, 8, 7, 6, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}

	for i,v := range testValues {
		res := calculateDD(baseHDD, baseDD, v)
		if res != testAnswers[i] {
			t.Errorf("calculateDD(%g, %g, %g) = %g; want %g",baseHDD, baseDD, v, res, testAnswers[i])
		}
	}
}

func TestCalculateCDD(t *testing.T) {
	baseCDD := 18.0
	testValues := []float64{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 15, 16, 17, 18, 19, 20,
	}

	testAnswers := []float64{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2,
	}

	for i,v := range testValues {
		res := calculateCDD(baseCDD, v)
		if res != testAnswers[i] {
			t.Errorf("calculateCDD(%g, %g) = %g; want %g",baseCDD, v, res, testAnswers[i])
		}
	}
}
