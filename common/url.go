package common

import (
	"strings"
)

func StringToSlice(src string)[]string {
	res := strings.Split(src, ",")
	return res
}
