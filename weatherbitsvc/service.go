package weatherbitsvc

import (
	"context"
	"github.com/flasherup/gradtage.de/hourlysvc"
)

type WBData struct {
	Date        string  `json:"date"`
	Temp        float64 `json:"temp"`
	Pod         string  `json:"pod"`
	Pres        float64 `json:"pres"`
	Timezone    string  `json:"timezone"`
	CountryCode string  `json:"country_code"`
	Clouds      float64 `json:"clouds"`
	Vis         float64 `json:"vis"`
	SolarRad    float64 `json:"solar_rad"`
	WindSpd     float64 `json:"wind_spd"`
	StateCode   string  `json:"state_code"`
	CityName    string  `json:"city_name"`
	AppTemp     float64 `json:"app_temp"`
	Uv          float64 `json:"uv"`
	Lon         float64 `json:"lon"`
	Slp         float64 `json:"slp"`
	HAngle      float64 `json:"h_angle"`
	Dewpt       float64 `json:"dewpt"`
	Snow        float64 `json:"snow"`
	Aqi         float64 `json:"aqi"`
	WindDir     float64 `json:"wind_dir"`
	ElevAngle   float64 `json:"elev_angle"`
	Ghi         float64 `json:"ghi"`
	Lat         float64 `json:"lat"`
	Precip      float64 `json:"precip"`
	Sunset      string  `json:"sunset"`
	Station     string  `json:"station"`
	Dni         float64 `json:"dni"`
	Sunrise     string  `json:"sunrise"`
}

type Service interface {
	GetPeriod(ctx context.Context, ids []string, start string, end string) (temps map[string][]hourlysvc.Temperature, err error)
	GetWBPeriod(ctx context.Context, id string, start string, end string) (temps []WBData, err error)
	GetUpdateDate(ctx context.Context, ids []string) (dates map[string]string, err error)
}
