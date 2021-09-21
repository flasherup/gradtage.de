package impl

import (
	"errors"
	"github.com/flasherup/gradtage.de/common"
	"strings"
)

var CommandUpdateAverage = "update_average"

func RunService(apiService *APISVC, name string, params map[string]string) (interface{}, error) {
	switch name {
	case common.ServiceMetrics:
		return processMetrics(apiService, params)
	}
	err := errors.New("service not found")
	return nil, err
}

func processMetrics(apiService *APISVC, params map[string]string) (interface{}, error) {
	stations, errStations := apiService.stations.GetAllStations()
	if errStations != nil {
		return nil, errStations
	}

	ids := make([]string, len(stations.Sts))
	i := 0
	for k, _ := range stations.Sts {
		ids[i] = strings.ToLower(k)
		i++
	}

	metrics, errMetrics := apiService.metrics.GetMetrics(ids)
	if errMetrics != nil {
		return nil, errMetrics
	}

	type MetricsJson struct {
		Name               string  `json:"Name"`
		Country            string  `json:"Country"`
		StationId          string  `json:"StationId"`
		Latitude           float64 `json:"Latitude"`
		Longitude          float64 `json:"Longitude"`
		Elevation          int     `json:"Elevation"`
		CityNameEnglish    string  `json:"CityNameEnglish"`
		CountryNameEnglish string  `json:"CountryNameEnglish"`
		ISO2CountryCode    string  `json:"ISO2CountryCode"`
		ICAO               string  `json:"ICAO"`
		WMO                string  `json:"WMO"`
		CWOP               string  `json:"CWOP"`
		Maslib             string  `json:"Maslib"`
		NationalID         string  `json:"NationalID"`
		IATA               string  `json:"IATA"`
		USAFWBAN           string  `json:"USAFWBAN"`
		GHCN               string  `json:"GHCN"`
		NWSLI              string  `json:"NWSLI" `
		LastUpdate         string  `json:"LastUpdate"`
		FirstUpdate        string  `json:"FirstUpdate"`
		RecordsAll         int32   `json:"RecordsAll"`
		RecordsClean       int32   `json:"RecordsClean"`
	}

	response := make([]MetricsJson, len(stations.Sts))

	i = 0
	for k, st := range stations.Sts {
		r := MetricsJson{}
		m, ok := metrics[strings.ToLower(k)]
		if ok {
			r.LastUpdate = m.LastUpdate
			r.FirstUpdate = m.FirstUpdate
			r.RecordsAll = m.RecordsAll
			r.RecordsClean = m.RecordsClean
		}
		r.StationId = st.Id
		response[i] = r
		i++
	}

	/*b, errMarshal := json.Marshal(response)
	if errMarshal != nil {
		return nil, errMarshal
	}
	fmt.Println(string(b))
	return string(b), nil*/
	return response, nil
}
