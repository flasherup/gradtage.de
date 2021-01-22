package database

import (
	"github.com/flasherup/gradtage.de/hourlysvc"
)

type WeatherBitDB interface {
	CreateTable(name string) (err error)
	RemoveTable(name string) (err error)
	GetPeriod(stID, start string, end string) (temps []hourlysvc.Temperature, err error)
	PushPeriod(stID string, temps []hourlysvc.Temperature) (err error)
	Dispose()
}
