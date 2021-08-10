package database

import (
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/weatherbitsvc"
	"github.com/flasherup/gradtage.de/weatherbitsvc/impl/parser"
)

type WeatherBitDB interface {
	CreateTable(name string) (err error)
	RemoveTable(name string) (err error)
	GetPeriod(stID, start string, end string) (temps []common.Temperature, err error)
	GetUpdateDate(stID string) (date string, err error)
	GetUpdateDateList(names []string) (temps map[string]string, err error)
	PushData(stID string, data *parser.WeatherBitData) (err error)
	GetWBData(name string, start string, end string) (wbd []weatherbitsvc.WBData, err error)
	PushWBData(stID string, wbd []weatherbitsvc.WBData) (err error)
	GetListOfTables() ([]string, error)
	Dispose()
}
