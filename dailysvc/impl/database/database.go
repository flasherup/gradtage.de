package database

import (
	"github.com/flasherup/gradtage.de/dailysvc"
)

type DailyDB interface {
	CreateTable(name string) (err error)
	RemoveTable(name string) (err error)
	GetPeriod(name string, start string, end string) (temps []dailysvc.Temperature, err error)
	GetDOYPeriod(name string, doy int, start string, end string) (temps []dailysvc.Temperature, err error)
	GetAll(name string) (temps []dailysvc.Temperature, err error)
	PushPeriod(name string, temps []dailysvc.Temperature) (err error)
	GetUpdateDate(name string) (date string, err error)
	GetUpdateDateList(names []string) (temps map[string]string, err error)
	GetTablesList() (map[string]bool, error)
	Dispose()
}
