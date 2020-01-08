package average

import (
	"errors"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/dailysvc"
	"github.com/flasherup/gradtage.de/dailysvc/impl/database"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"math"
	"time"
)


const averagePrefix = "_avg"
const hoursInYear = 8760


type Average struct {
	logger  	log.Logger
	db 			database.Postgres
}


func NewAverage(logger log.Logger, db database.Postgres) *Average {
	avg := Average{
		logger: logger,
		db:		db,
	}
	return &avg
}


func (avg Average)CalculateAndSaveYearlyAverage(stId string) error{
	tn := time.Now()
	start := tn.Add(-time.Hour * hoursInYear * 10)

	err := avg.db.CreateTable(stId + averagePrefix)
	if err != nil {
		level.Error(avg.logger).Log("msg", "Station creation error", "err", err)
		return err
	}

	for i:=1; i<=366; i++ {
			temp, err := avg.db.GetDOYPeriod(stId, i, start.Format(common.TimeLayout), tn.Format(common.TimeLayout))
			if err != nil {
				level.Error(avg.logger).Log("msg", "CalculateAndSaveYearlyAverage error", "err", err)
				continue
			}

			aTemp, err := calculateAverageForDOY(temp)
			if err != nil {
				level.Error(avg.logger).Log("msg", "CalculateAndSaveYearlyAverage error", "err", err)
				continue

			}
			avg.db.PushPeriod(stId + averagePrefix, []dailysvc.Temperature{ aTemp })
	}
	return nil
}

func (avg Average)CalculateAndSaveDOYAverage(stId string, doy int) error{
	err := avg.db.CreateTable(stId + averagePrefix)
	if err != nil {
		level.Error(avg.logger).Log("msg", "Station creation error", "err", err)
		return err
	}

	tn := time.Now()
	start := tn.Add(-time.Hour * hoursInYear * 10)
	temp, err := avg.db.GetDOYPeriod(stId, doy, start.Format(common.TimeLayout), tn.Format(common.TimeLayout))
	if err != nil {
		level.Error(avg.logger).Log("msg", "CalculateAndSaveYearlyAverage error", "err", err)
		return err
	}

	aTemp, err := calculateAverageForDOY(temp)
	if err != nil {
		level.Error(avg.logger).Log("msg", "CalculateAndSaveYearlyAverage error", "err", err)
		return err

	}
	avg.db.PushPeriod(stId + averagePrefix, []dailysvc.Temperature{ aTemp })
	return nil
}

func ToAverageDate( src string) (string, error) {
	d, err := time.Parse(common.TimeLayout, src)
	if err != nil {
		return src, err
	}
	date := time.Date(3000, d.Month(), d.Day(), 0, 0, 0, 0, d.Location() )
	return date.Format(common.TimeLayout),nil
}

func (avg Average)GetAll(name string) (temps map[int]dailysvc.Temperature, err error ) {
	t, err := avg.db.GetAll(name + averagePrefix)
	if err != nil {
		level.Error(avg.logger).Log("msg", "GetAvg error", "err", err)
		return nil,err
	}
	temps = make(map[int]dailysvc.Temperature)
	for _,v := range t {
		d, err := time.Parse(common.TimeLayout, v.Date)
		if err != nil {
			level.Error(avg.logger).Log("msg", "GetAvg error", "err", err)
			return nil,err
		}

		temps[d.YearDay()] = v
	}
	return temps,err
}

func calculateAverageForDOY(temperatures []dailysvc.Temperature) (dailysvc.Temperature, error) {
	res := dailysvc.Temperature{}
	if len(temperatures) == 0 {
		return res, errors.New("source temperature is empty")
	}
	var latest []dailysvc.Temperature
	if len(temperatures) > 10 {
		latest = temperatures[:len(temperatures)-10]
	} else {
		latest = temperatures
	}

	sum := 0.0
	count := 0.0
	for _,v := range latest {
		sum += v.Temperature
		count += 1.0
	}

	date, err := ToAverageDate(temperatures[0].Date)
	if err != nil {
		return res,err
	}
	value := sum/count
	return dailysvc.Temperature{ Date:date, Temperature:math.Round(value*10)/10 }, nil
}