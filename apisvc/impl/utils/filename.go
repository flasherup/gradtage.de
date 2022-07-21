package utils

import (
	"fmt"
	"github.com/flasherup/gradtage.de/common"
	"time"
)

func GetCSVName(output, station string, tb, tr float64) string {
	if output == common.DDType {
		return fmt.Sprintf("%s_DD_%gC_%gC.csv",
			station,
			tb,
			tr)
	}

	if output == common.HDDType {
		return fmt.Sprintf("%s_HDD_%gC.csv",
			station,
			tb)
	}
	if output == common.CDDType {
		return fmt.Sprintf("%s_CDD_%gC.csv",
			station,
			tb)
	}

	return "unnamed.scv"
}

func GetJSONName(output, station string, tb, tr float64) string {
	if output == common.DDType {
		return fmt.Sprintf("%s_DD_%gC_%gC.json",
			station,
			tb,
			tr)
	}

	if output == common.HDDType {
		return fmt.Sprintf("%s_HDD_%gC.json",
			station,
			tb)
	}
	if output == common.CDDType {
		return fmt.Sprintf("%s_CDD_%gC.json",
			station,
			tb)
	}

	return "unnamed.scv"
}

func GetZIPName(output string) string {
	d := time.Now().Format(common.TimeLayoutDay)
	return fmt.Sprintf("%s_Energy-Data_%s.zip",d ,output)
}
