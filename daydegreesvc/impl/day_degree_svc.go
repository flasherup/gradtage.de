package impl

import (
	"context"
	"fmt"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/daydegreesvc"
	"github.com/flasherup/gradtage.de/daydegreesvc/config"
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
	level.Info(dd.logger).Log("msg", "GetDegree", "Station", params.Station, "Start", params.Start, "End", params.End)
	temps, err := dd.weatherBit.GetPeriod([]string{params.Station}, params.Start, params.End)
	if err != nil {
		level.Error(dd.logger).Log("msg", "GetPeriod error", "err", err)
		return []daydegreesvc.Degree{}, err
	}

	var degrees *[]common.Temperature
	t := (*temps)[params.Station]
	if params.Output == common.HDDType {
		degrees = common.CalculateHDDDegree(t, params.Tb, params.Breakdown, params.DayCalc, params.WeekStarts)
	} else if params.Output == common.DDType {
		degrees = common.CalculateDDegree(t, params.Tb, params.Tr, params.Breakdown, params.DayCalc, params.WeekStarts)
	} else if params.Output == common.CDDType {
		degrees = common.CalculateCDDegree(t, params.Tb, params.Breakdown, params.DayCalc, params.WeekStarts)
	}

	res := toDegree(degrees)
	return *res, nil
}

func (dd *DayDegreeSVC) GetAverageDegree(ctx context.Context, params daydegreesvc.Params, years int) ([]daydegreesvc.Degree, error) {
	if years < 1 {
		years = 1
	}else if years > 10 {
		years = 10
	}

	start, end, err := getSDates(years)
	level.Info(dd.logger).Log("msg", "GetDegree", "Station", params.Station, "Start", start, "End", end, "years", years);

	if err != nil{
		level.Error(dd.logger).Log("msg", "Get WB Data start average date error", "err", err)
		return []daydegreesvc.Degree{}, err
	}
	//temps, err := dd.weatherBit.GetAverage(params.Station, years, end)
	temps, err := dd.weatherBit.GetPeriod([]string{params.Station}, start, end)
	if err != nil {
		level.Error(dd.logger).Log("msg", "GetPeriod error", "err", err)
		return []daydegreesvc.Degree{}, err
	}

	var degrees *[]common.Temperature
	t := (*temps)[params.Station]
	if params.Output == common.HDDType {
		degrees = common.CalculateHDDDegree(t, params.Tb, params.Breakdown, params.DayCalc, params.WeekStarts)
	} else if params.Output == common.DDType {
		degrees = common.CalculateDDegree(t, params.Tb, params.Tr, params.Breakdown, params.DayCalc, params.WeekStarts)
	} else if params.Output == common.CDDType {
		degrees = common.CalculateCDDegree(t, params.Tb, params.Breakdown, params.DayCalc, params.WeekStarts)
	}

	days := make(map[string][]float64)
	keyFormat := "%d-%d-%d"

	for _, v := range *degrees {
		d, err := common.ParseTimeByBreakdown(v.Date, params.Breakdown)
		if err != nil{
			level.Error(dd.logger).Log("msg", "Time Parse error", "err", err)
			return []daydegreesvc.Degree{}, err
		}

		key := fmt.Sprintf(keyFormat, d.Month(), d.Day(), d.Hour())
		day, exist := days[key]
		if !exist {
			day = make([]float64, 0)
		}

		days[key] = append(day, v.Temp)
	}

	res := make([]common.Temperature, 0)

	initialDate, _ := time.Parse(common.TimeLayout, common.InitialDate)
	year := initialDate.Year()
	for initialDate.Year() == year {
		key := fmt.Sprintf(keyFormat, initialDate.Month(), initialDate.Day(), initialDate.Hour())
		day, exist := days[key]
		var temp = 0.0

		d := common.GetDateStringByBreakdown(initialDate, params.Breakdown)
		if exist {
			temp = common.GetAverageFloat64(day)
			temp = common.ToFixedFloat64(temp, 2)

		} else {
			temp = common.EmptyWeather

		}

		res = append(res, common.Temperature{
			Date: d,
			Temp: temp,
		})

		initialDate = addPeriod(initialDate, params.Breakdown)
	}

	return *toDegree(&res), nil
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

func getStartDate(end string, years int) (string, error) {
	t, err := time.Parse(common.TimeLayoutWBH, end)
	if err != nil{
		return end, err
	}
	layout := common.TimeLayoutWBH
	time.Parse(layout, end)
	start := t.AddDate(-years, 0, 0)
	return start.Format(common.TimeLayoutWBH), nil
}

func getSDates(years int) (string, string, error) {
	end := time.Now()
	start := end.AddDate(-years, 0, 0)
	return start.Format(common.TimeLayoutWBH), end.Format(common.TimeLayoutWBH), nil
}

func addPeriod(src time.Time, breakdown string) time.Time {
	if breakdown == common.BreakdownWeekly {
		return src.AddDate(0, 0, 7)
	}

	if breakdown == common.BreakdownWeeklyISO {
		return src.AddDate(0, 0, 7)
	}

	if breakdown == common.BreakdownMonthly {
		return src.AddDate(0, 1, 0)
	}

	if breakdown == common.BreakdownYearly {
		return src.AddDate(1, 0, 0)
	}

	return src.AddDate(0, 0, 1)
}
