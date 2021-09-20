package utils

import (
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/metricssvc/mtrgrpc"
	"github.com/flasherup/gradtage.de/weatherbitsvc"
	"time"
)

func GetWeatherbitMetrics(wbData *[]weatherbitsvc.WBData) (*mtrgrpc.Metrics, error) {
	date := time.Now()
	res := mtrgrpc.Metrics {
		Date: date.Format(common.TimeLayout),
	}


	cutDate, err := time.Parse(common.TimeLayoutWBH, common.CutDateWBH)
	if err != nil {
		return nil, err
	}

	var lastUpdate time.Time
	var firstUpdate time.Time
	recordsAll := int32(0)
	recordsClean := int32(0)

	for _,v := range *wbData {
		cDate, err := time.Parse(common.TimeLayout, v.Date)
		if err != nil {
			return nil, err
		}
		if lastUpdate.IsZero() || cDate.After(lastUpdate) {
			lastUpdate = cDate
		}

		if firstUpdate.IsZero() || cDate.Before(firstUpdate) {
			firstUpdate = cDate
		}

		recordsAll++

		if cDate.After(cutDate) {
			recordsClean++
		}
	}

	res.LastUpdate = lastUpdate.Format(common.TimeLayout)
	res.FirstUpdate = firstUpdate.Format(common.TimeLayout)
	res.RecordsAll = recordsAll
	res.RecordsClean = recordsClean

	return &res, nil
}
