package utils

import "time"

func SleepCheck(numberOfRequests, counter int, startTime time.Time, duration time.Duration) (int, time.Time) {
	now := time.Now()
	dif := now.Sub(startTime)
	if dif >= duration {
		return 0, now
	}

	counter++
	if counter >= numberOfRequests {
		time.Sleep(duration - dif)
		return 0, time.Now()
	}
	return counter, startTime
}

func DaysCheck(start, end time.Time, daysCount int) bool {
	dif := end.Sub(start)
	return int(dif.Hours()/24) >= daysCount
}

func DaysDifference(start, end time.Time) int {
	dif := end.Sub(start)
	return int(dif.Hours()/24)
}
