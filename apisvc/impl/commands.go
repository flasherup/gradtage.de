package impl

import (
	"github.com/flasherup/gradtage.de/stationssvc/stsgrpc"
	"github.com/go-kit/kit/log/level"
	"strings"
)

var CommandUpdateAverage = "update_average"


func ParseCommand(apiService APISVC, name string, params map[string]string) (interface{}, error){
	switch(name) {
	case CommandUpdateAverage:
		return updateAverage(apiService, params);
	}


	resp := struct {
		Command string `json:"command"`
		Error string `json:"error"`
	}{
		name,
		"not found",
	}
	return resp, nil
}

func updateAverage(apiService APISVC,  params map[string]string) (interface{}, error) {
	var stsList []string
	statiosIds, ok := params["stations"];
	if (ok) {
		stsList = strings.Split(statiosIds, ",");
	}

	var sts map[string]*stsgrpc.Station

	level.Info(apiService.logger).Log("msg", "Get Stations list strings", "stsList", len(stsList))
	if stsList!= nil && len(stsList) > 0 {
		resp, err := apiService.stations.GetStations(stsList)
		if err != nil {
			return "", err
		}
		sts = resp.Sts
	} else {
		resp, err := apiService.stations.GetAllStations()
		if err != nil {
			return "", err
		}

		sts = resp.Sts
	}

	level.Info(apiService.logger).Log("msg", "Get Stations list", "sts", len(sts))

	count := 0
	for k,_ := range sts {
		level.Info(apiService.logger).Log("msg", "Update Average", "sts", k)
		resp, err := apiService.daily.UpdateAvgForYear(k)
		if err != nil {
			level.Error(apiService.logger).Log("msg", "Average Update Error", "id", k, "err", err)
		} else if resp.Err != "nil" {
			level.Error(apiService.logger).Log("msg", "Average Update Error", "id", k, "err", resp.Err)
		} else {
			count++
		}
	}
	level.Info(apiService.logger).Log("msg", "Start Average Update Complete", "count", count)


	return sts, nil
}