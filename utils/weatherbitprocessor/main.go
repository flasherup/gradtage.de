package main

import (
	"flag"
	"fmt"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/utils/weatherbitprocessor/config"
	"github.com/flasherup/gradtage.de/utils/weatherbitprocessor/csv"
	"github.com/flasherup/gradtage.de/weatherbitsvc"
	"github.com/flasherup/gradtage.de/weatherbitsvc/impl"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"
	"time"
)

func main() {
	configFile := flag.String("config.file", "config.yml", "Config file name.")
	//startDate := flag.String("start", "2021-07-12", "Start Date.")
	//endDate := flag.String("end", "2021-07-18", "End Date.")
	//csvFile := flag.String("csv.file", "stationsR.csv", "CSV file name. ")
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger)
	}

	//Config
	conf, err := config.LoadConfig(*configFile)
	if err != nil {
		level.Error(logger).Log("msg", "config loading error", "exit", err.Error())
		return
	}

	client := impl.NewWeatherBitSVCClient(conf.Clients.WeatherBitAddr,logger)

	level.Info(logger).Log("msg", "client started")
	defer level.Info(logger).Log("msg", "client ended")

	//calculateEntries(client, logger, *csvFile, *startDate, *endDate)
	//calculateHourlyDegree(client, logger)
	//addDegreeDataFormCSV(client, "data/KDCA_C.csv", logger)
	addDailyDegreeDataFormCSV(client, "data/US_KDCA_daily.csv", logger)
}

func calculateEntries(client *impl.WeatherBitSVCClient, logger log.Logger, stations, startDate, endDate string) {
	//Just for test
	stationList, err := csv.CSVToMap(stations)
	if err != nil {
		level.Error(logger).Log("msg", "CSV loading error", "err", err.Error())
		return
	}

	for innerId, stationId := range *stationList {
		level.Info(logger).Log("Station:", stationId, "innerId", innerId)
		data, err := client.GetPeriod([]string{innerId}, startDate, endDate)
		if err != nil {
			level.Error(logger).Log("Station error:", stationId, "innerId", innerId, "error", err)
			continue
		}

		for _,temperatures := range *data {
			if len(temperatures) == 0 {
				level.Error(logger).Log("Error", "no entries", "Station:", stationId, "innerId", innerId)
				continue
			}

			counted := countEntriesPerDay(&temperatures)
			for k,v := range counted {
				//level.Info(logger).Log("date", k, "hours", v)
				if v < 24 {
					level.Error(logger).Log("Error", "less the 24 entries", "date", k, "hours", v, "Station:", stationId, "innerId", innerId)
				}
			}
		}
	}
}

func countEntriesPerDay(temps *[]common.Temperature) map[string]int {
	res := make(map[string]int)
	for _,v := range *temps {
		date,err := time.Parse(common.TimeLayout, v.Date)
		if  err !=  nil {
			continue
		}
		cattedDate := date.Format("2006-01-02")
		if val, ok := res[cattedDate]; ok {
			res[cattedDate] = val+1
		} else {
			res[cattedDate] = 1
		}
	}
	return res
}

func calculateHourlyDegree(client *impl.WeatherBitSVCClient, logger log.Logger) {
	stationId := "us_kheq";
	data, err := client.GetPeriod([]string{stationId}, "2020-01-01T00:00:00", "2021-01-01T00:00:00")
	if err != nil {
		level.Error(logger).Log("Station error:", stationId, "error", err)
		return
	}

	temps := (*data)[stationId]
	/*for _,temperatures :=  range temps {
		fmt.Println(temperatures.Date, temperatures.Temp)
	}*/

	//Hdd
	fmt.Println("Hdd Day")
	hdd := common.CalculateHDDDegree(temps, 15, common.BreakdownDaily, common.DayCalcMean, time.Monday)
	//fmt.Println(hdd)
	for _,temps :=  range *hdd {
		fmt.Println(temps.Date, temps.Temp)
	}

	fmt.Println("Hdd Week ISO")
	hdd = common.CalculateHDDDegree(temps, 15, common.BreakdownWeeklyISO, common.DayCalcMean, time.Monday)
	//fmt.Println(hdd)
	for _,temps :=  range *hdd {
		fmt.Println(temps.Date, temps.Temp)
	}

	fmt.Println("Hdd Month")
	hdd = common.CalculateHDDDegree(temps, 15, common.BreakdownMonthly, common.DayCalcMean, time.Monday)
	//fmt.Println(hdd)
	for _,temps :=  range *hdd {
		fmt.Println(temps.Date, temps.Temp)
	}

	fmt.Println("Hdd Year")
	hdd = common.CalculateHDDDegree(temps, 15, common.BreakdownYearly, common.DayCalcMean, time.Monday)
	//fmt.Println(hdd)
	for _,temps :=  range *hdd {
		fmt.Println(temps.Date, temps.Temp)
	}
}

func addDegreeDataFormCSV(client *impl.WeatherBitSVCClient, fileName string , logger log.Logger) {
	tempList, err := csv.CSVToTempsData(fileName)
	if err != nil {
		level.Error(logger).Log("msg", "CSV loading error", "err", err.Error())
		return
	}

	data := make([]weatherbitsvc.WBData, len(*tempList))
	for i,v := range *tempList {
		d, err := time.Parse("2006-01-02 04:05", v.Date)
		if err != nil {
			level.Error(logger).Log("msg", "Date parse error", "err", err.Error())
			continue
		}
		data[i] = weatherbitsvc.WBData{
			Date: d.Format(common.TimeLayout),
			Temp: v.Temp,
			Timezone: v.Timezone,
		}
	}

	stName := "US_KDCA"

	level.Info(logger).Log("msg", "Station: " + stName + " data received", "count", len(data))
	err = client.PushWBPeriod(stName, data)
	if err != nil{
		level.Error(logger).Log("msg", "Saving station data error","station", stName, "error", err.Error())
	}
}

func addDailyDegreeDataFormCSV(client *impl.WeatherBitSVCClient, fileName string , logger log.Logger) {
	tempList, err := csv.CSVDailyToHourlyTempsData(fileName)
	if err != nil {
		level.Error(logger).Log("msg", "CSV loading error", "err", err.Error())
		return
	}

	data := make([]weatherbitsvc.WBData, len(*tempList))
	for i,v := range *tempList {
		d, err := time.Parse("2006-01-02 15:04", v.Date)
		if err != nil {
			level.Error(logger).Log("msg", "Date parse error", "err", err.Error())
			continue
		}
		data[i] = weatherbitsvc.WBData{
			Date: d.Format(common.TimeLayout),
			Temp: v.Temp,
			Timezone: v.Timezone,
		}
	}

	stName := "US_KDCA"

	level.Info(logger).Log("msg", "Station: " + stName + " data received", "count", len(data))
	err = client.PushWBPeriod(stName, data)
	if err != nil{
		level.Error(logger).Log("msg", "Saving station data error","station", stName, "error", err.Error())
	}
}

