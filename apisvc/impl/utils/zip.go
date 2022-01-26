package utils

import (
	"fmt"
	"github.com/flasherup/gradtage.de/common"
	"time"
)

func GetZIPName(output string) string {
	d := time.Now().Format(common.TimeLayoutDay)
	return fmt.Sprintf("%s_Energy-Data_%s.zip",d ,output)
}
