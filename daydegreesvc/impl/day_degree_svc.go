package impl

import (
	"context"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/daydegreesvc"
	"github.com/flasherup/gradtage.de/daydegreesvc/config"
	"github.com/flasherup/gradtage.de/daydegreesvc/utils"
	"github.com/flasherup/gradtage.de/weatherbitsvc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"time"
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
	level.Info(dd.logger).Log("msg", "Get Degree", "Station", params.Station, "Start", params.Start, "End", params.End)
	temps, err := dd.weatherBit.GetPeriod([]string{params.Station}, params.Start, params.End)
	if err != nil {
		level.Error(dd.logger).Log("msg", "GetPeriod error", "err", err)
		return []daydegreesvc.Degree{}, err
	}
	degrees := calculateDayDegree(*temps, params)
	res := utils.ToDegree(degrees)
	return res, nil
}

func (dd *DayDegreeSVC) GetAverageDegree(ctx context.Context, params daydegreesvc.Params, years int) ([]daydegreesvc.Degree, error) {
	if years < 1 {
		years = 1
	}else if years > 10 {
		years = 10
	}

	start, end, err := getSDates(years)
	level.Info(dd.logger).Log("msg", "Get Average Degree", "Station", params.Station, "Start", start, "End", end, "years", years);

	if err != nil{
		level.Error(dd.logger).Log("msg", "Get WB Data start average date error", "err", err)
		return []daydegreesvc.Degree{}, err
	}
	temps, err := dd.weatherBit.GetPeriod([]string{params.Station}, start, end)
	if err != nil {
		level.Error(dd.logger).Log("msg", "GetPeriod error", "err", err)
		return []daydegreesvc.Degree{}, err
	}

	degrees := calculateDayDegree(*temps, params)
	var res []daydegreesvc.Degree
	if params.Breakdown == common.BreakdownWeekly || params.Breakdown == common.BreakdownWeeklyISO {
		res, err = utils.WeeklyAverage(degrees, params)
	} else {
		res, err = utils.CommonAverage(degrees, params)
	}

	return res, err
}

func calculateDayDegree(temps map[string][]common.Temperature, params daydegreesvc.Params) []common.Temperature {
	var degrees *[]common.Temperature
	t := (temps)[params.Station]
	if params.Output == common.HDDType {
		degrees = common.CalculateHDDDegree(t, params.Tb, params.Breakdown, params.DayCalc, params.WeekStart)
	} else if params.Output == common.DDType {
		degrees = common.CalculateDDegree(t, params.Tb, params.Tr, params.Breakdown, params.DayCalc, params.WeekStart)
	} else if params.Output == common.CDDType {
		degrees = common.CalculateCDDegree(t, params.Tb, params.Breakdown, params.DayCalc, params.WeekStart)
	}

	if degrees == nil {
		return []common.Temperature{}
	}

	return *degrees
}

func getSDates(years int) (string, string, error) {
	end := time.Now()
	start := time.Date(end.Year()-years, end.Month(), end.Day(), end.Hour(), end.Minute(), end.Second(), end.Nanosecond(), end.Location())
	return start.Format(common.TimeLayoutWBH), end.Format(common.TimeLayoutWBH), nil
}
