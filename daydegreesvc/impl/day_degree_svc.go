package impl

import (
	"context"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/daydegreesvc"
	"github.com/flasherup/gradtage.de/daydegreesvc/config"
	"github.com/flasherup/gradtage.de/weatherbitsvc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type DayDegreeSVC struct {
	weatherBit    weatherbitsvc.Client
	alert 		alertsvc.Client
	logger  	log.Logger
	conf		config.DayDegreeConfig
}



func NewDayDegreeSVC(
	logger 		log.Logger,
	weatherBit 	weatherbitsvc.Client,
	alert 		alertsvc.Client,
	conf 		config.DayDegreeConfig,
) (*DayDegreeSVC, error) {
	wb := DayDegreeSVC {
		weatherBit:weatherBit,
		alert:alert,
		logger:logger,
		conf:conf,
	}

	return &wb,nil
}

func (dd *DayDegreeSVC) GetDegree(ctx context.Context, params daydegreesvc.Params) ([]daydegreesvc.Degree, error) {
	level.Info(dd.logger).Log("msg", "GetDegree", "Station", params.Station, "Start", params.Start, "End", params.End)
	temps, err := dd.weatherBit.GetPeriod([]string{params.Station}, params.Start, params.End)
	if err != nil {
		level.Error(dd.logger).Log("msg", "GetPeriod error", "err", err)
		return []daydegreesvc.Degree{}, err
	}

	var degrees *[]common.Temperature
	t := (*temps)[params.Station]
	if params.Method == common.HDDType {
		degrees = common.CalculateHDDDegree(t, params.Tb, params.Breakdown, params.DayCalc)
	} else if params.Method == common.DDType {
		degrees = common.CalculateDDegree(t, params.Tb, params.Tr, params.Breakdown, params.DayCalc)
	} else if params.Method == common.CDDType {
		degrees = common.CalculateCDDegree(t, params.Tb, params.Breakdown, params.DayCalc)
	}

	res := toDegree(degrees)
	return *res, nil
}

func toDegree(temps *[]common.Temperature) *[]daydegreesvc.Degree {
	if temps == nil {
		return &[]daydegreesvc.Degree{}
	}
	res := make([]daydegreesvc.Degree, len(*temps))
	for i,v := range *temps {
		res[i] =  daydegreesvc.Degree{
			Date: v.Date,
			Temp: v.Temp,
		}
	}
	return &res
}
