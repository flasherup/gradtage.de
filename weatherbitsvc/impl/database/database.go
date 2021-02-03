package database

import (
	"github.com/flasherup/gradtage.de/hourlysvc"
	"github.com/flasherup/gradtage.de/weatherbitsvc/impl/parser"
)

type WeatherBitDB interface {
	CreateTable(name string) (err error)
	RemoveTable(name string) (err error)
	GetPeriod(stID, start string, end string) (temps []hourlysvc.Temperature, err error)
	PushData(stID string, data *parser.WeatherBitData) (err error)
	Dispose()
}
