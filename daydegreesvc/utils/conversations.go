package utils

import (
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/daydegreesvc"
)

func ToDegree(temps []common.Temperature) []daydegreesvc.Degree {
	if temps == nil {
		return []daydegreesvc.Degree{}
	}
	res := make([]daydegreesvc.Degree, len(temps))
	for i,v := range temps {
		res[i] =  daydegreesvc.Degree{
			Date: v.Date,
			Temp: v.Temp,
		}
	}
	return res
}