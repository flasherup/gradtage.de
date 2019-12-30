package database

import (
	"github.com/flasherup/gradtage.de/hourlysvc"
)

type HourlyDB interface {
	CreateTable(name string) (err error)
	RemoveTable(name string) (err error)
	GetPeriod(stID, start string, end string) (temps []hourlysvc.Temperature, err error)
	PushPeriod(stID string, temps []hourlysvc.Temperature) (err error)
	GetUpdateDate(stID string) (date string, err error)
	Dispose()
}
