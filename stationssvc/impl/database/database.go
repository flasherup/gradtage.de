package database

import "github.com/flasherup/gradtage.de/stationssvc"

type StationsDB interface {
	CreateTable() (err error)
	RemoveTable() (err error)
	AddStation(station stationssvc.Station) error
	AddStations(stations []stationssvc.Station) error
	DeleteStation(id string) error
	GetStations(ids []string) ([]stationssvc.Station,error)
	GetAllStations() ([]stationssvc.Station,error)
	GetStationsBySrcType(types []string) ([]stationssvc.Station,error)
	GetCount() (int, error)
	Dispose()
}