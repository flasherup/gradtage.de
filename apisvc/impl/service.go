package impl

import (
	"errors"
	"github.com/flasherup/gradtage.de/autocompletesvc/acrpc"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/metricssvc/mtrgrpc"
	"strings"
	"time"
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

const (
	StationStatusOperating = "operating"
)

type MetricsCustomer struct {
	Name          string  `json:"Name"`
	Country       string  `json:"Country"`
	StationId     string  `json:"StationId"`
	Latitude      float64 `json:"Latitude"`
	Longitude     float64 `json:"Longitude"`
	Elevation     float64 `json:"Elevation"`
	StationStatus string  `json:"StationStatus"`
}

type MetricsInternal struct {
	Name         string  `json:"Name"`
	Country      string  `json:"Country"`
	StationId    string  `json:"StationId"`
	Latitude     float64 `json:"Latitude"`
	Longitude    float64 `json:"Longitude"`
	LastUpdate   string  `json:"LastUpdate"`
	FirstUpdate  string  `json:"FirstUpdate"`
	RecordsAll   int32   `json:"RecordsAll"`
	RecordsClean int32   `json:"RecordsClean"`
}

func processMetrics(apiService *APISVC, params map[string]string) (interface{}, error) {
	stations, errStations := apiService.autocomplete.GetAllStations()
	if errStations != nil {
		return nil, errStations
	}

	ids := make([]string, len(stations))
	i := 0
	for k, _ := range stations {
		ids[i] = strings.ToLower(k)
		i++
	}

	metrics, errMetrics := apiService.metrics.GetMetrics(ids)
	if errMetrics != nil {
		return nil, errMetrics
	}

	use, exist :=  params["use"]
	if exist && use == "internal" {
		return makeInternalMetrics(stations, metrics), nil
	}

	return makeCustomerMetrics(stations, metrics), nil
}

func makeCustomerMetrics(stations map[string]*acrpc.Source, metrics map[string]*mtrgrpc.Metrics) []MetricsCustomer {
	response := make([]MetricsCustomer, len(stations))
	now := time.Now()
	i := 0
	for k, st := range stations {
		r := MetricsCustomer{}
		m, ok := metrics[strings.ToLower(k)]
		if ok {
			lu, err := time.Parse(common.TimeLayout, m.LastUpdate)
			if err == nil && now.Sub(lu) > time.Hour * 168 {
				r.StationStatus = m.LastUpdate
			} else {
				r.StationStatus = StationStatusOperating
			}
		}
		r.Name = st.CityNameEnglish
		r.Country = st.CountryNameEnglish
		r.StationId = st.ID
		r.Elevation = st.Elevation
		r.Latitude = st.Latitude
		r.Longitude = st.Longitude
		response[i] = r
		i++
	}

	return response
}

func makeInternalMetrics(stations map[string]*acrpc.Source, metrics map[string]*mtrgrpc.Metrics) []MetricsInternal {
	res := make([]MetricsInternal, len(stations))
	i := 0
	for k, st := range stations {
		r := MetricsInternal{}
		m, ok := metrics[strings.ToLower(k)]
		if ok {
			r.LastUpdate = m.LastUpdate
			r.FirstUpdate = m.FirstUpdate
			r.RecordsAll = m.RecordsAll
			r.RecordsClean = m.RecordsClean
		}
		r.Name = st.CityNameEnglish
		r.Country = st.CountryNameEnglish
		r.StationId = st.ID
		r.Latitude = st.Latitude
		r.Longitude = st.Longitude
		res[i] = r
		i++
	}

	return res
}