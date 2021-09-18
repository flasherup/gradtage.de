package common

import "strings"

func FixSingleQuote(src string) string {
	return strings.ReplaceAll(src, "'", "''")
}
