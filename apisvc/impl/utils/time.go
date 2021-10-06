package utils

import (
	"github.com/flasherup/gradtage.de/common"
	"time"
)

const (
	wordToday = "today"
)

func WordToTime(word string) string {
	now := time.Now()
	if word == wordToday {
		return now.Format(common.TimeLayoutWBH)
	}

	return word
}
