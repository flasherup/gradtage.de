package common

import (
	"strings"
)

func ParseStations(str string) []string {
	firstComa := strings.Index(str,",")
	if firstComa == -1 {
		return []string{str}
	}

	return StringToSlice(str)
}

func StringToSlice(src string)[]string {
	res := strings.Split(src, ",")
	return res
}
