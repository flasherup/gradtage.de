package database

import (
	"github.com/flasherup/gradtage.de/dailysvc"
)

type HourlyDB interface {
	CreateTable(name string) (err error)
	RemoveTable(name string) (err error)
	GetPeriod(stID, start string, end string) (temps []dailysvc.Temperature, err error)
	PushPeriod(stID string, temps []dailysvc.Temperature) (err error)
	GetUpdateDate(stID string) (date string, err error)
	Dispose()
}
