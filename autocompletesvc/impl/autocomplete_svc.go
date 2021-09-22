package impl

import (
	"context"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/autocompletesvc"
	"github.com/flasherup/gradtage.de/autocompletesvc/acrpc"
	"github.com/flasherup/gradtage.de/autocompletesvc/impl/database"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type AutocompleteSVC struct {
	logger  	log.Logger
	alert 		alertsvc.Client
	db 			database.AutocompleteDB
}

func NewAutocompleteSVC(logger log.Logger, db database.AutocompleteDB, alert alertsvc.Client) (*AutocompleteSVC, error) {
	err := db.CreateTable()
	if err != nil {
		return nil, err
	}

	st := AutocompleteSVC{
		logger: logger,
		alert:  alert,
		db:		db,
	}
	return &st,nil
}

func (ss AutocompleteSVC) GetAutocomplete(ctx context.Context, text string) (result map[string][]autocompletesvc.Autocomplete, err error) {
	level.Info(ss.logger).Log("msg", "GetAutocomplete", "text", text)
	result, err = ss.db.GetStationId(text)
	if err != nil {
		level.Error(ss.logger).Log("msg", "GetAutocomplete DB error", "err", err)
		ss.sendAlert(NewErrorAlert(err))
		return nil,err
	}
	return result, nil
}

func (ss AutocompleteSVC) AddSources(ctx context.Context, sources []autocompletesvc.Autocomplete) (err error) {
	level.Info(ss.logger).Log("msg", "AddSource", "length", len(sources))

	sources = *ss.validateFields(&sources)

	err = ss.db.AddSources(sources)
	if err != nil {
		level.Error(ss.logger).Log("msg", "AddSources DB error", "err", err)
		ss.sendAlert(NewErrorAlert(err))
		return err
	}
	return nil
}

func (ss AutocompleteSVC) ResetSources(ctx context.Context, sources []autocompletesvc.Autocomplete) (err error) {
	level.Info(ss.logger).Log("msg", "ResetSource", "length", len(sources))
	err = ss.db.RemoveTable()
	if err != nil {
		level.Error(ss.logger).Log("msg", "Remove Table DB error", "err", err)
		return err
	}

	err = ss.db.CreateTable()
	if err != nil {
		level.Error(ss.logger).Log("msg", "Create Table DB error", "err", err)
		return err
	}

	sources = *ss.validateFields(&sources)

	err = ss.db.AddSources(sources)
	if err != nil {
		level.Error(ss.logger).Log("msg", "AddSources DB error", "err", err)
		return err
	}

	return nil
}

func (ss AutocompleteSVC) validateFields(sources *[]autocompletesvc.Autocomplete) *[]autocompletesvc.Autocomplete {
	namesLength := 70
	for i,v := range *sources {
		if len(v.CountryNameNative) > namesLength {
			level.Warn(ss.logger).Log("msg", "CountryNameNative too long", "index", i, "value", v.CountryNameNative)
			v.CountryNameNative = v.CountryNameNative[:namesLength]
		}

		if len(v.CountryNameEnglish) > namesLength {
			level.Warn(ss.logger).Log("msg", "CountryNameEnglish too long", "index", i, "value", v.CountryNameEnglish)
			v.CountryNameEnglish = v.CountryNameEnglish[:namesLength]
		}

		if len(v.CityNameNative) > namesLength {
			level.Warn(ss.logger).Log("msg", "CityNameNative too long", "index", i, "value", v.CityNameNative)
			v.CityNameNative = v.CityNameNative[:namesLength]
		}

		if len(v.CityNameEnglish) > namesLength {
			level.Warn(ss.logger).Log("msg", "CityNameEnglish too long", "index", i, "value", v.CityNameEnglish)
			v.CityNameEnglish = v.CityNameEnglish[:namesLength]
		}
	}

	return sources
}


func (ss AutocompleteSVC) GetAllStations(ctx context.Context) (map[string]*acrpc.Source, error) {
	level.Info(ss.logger).Log("msg", "GetAllStations")
	sts, err := ss.db.GetAllStations()
	if err != nil {
		level.Error(ss.logger).Log("msg", "Get All Stations Error", "err", err)
		return nil, err
	}
	return sts, nil
}

func (ss AutocompleteSVC)sendAlert(alert alertsvc.Alert) {
	err := ss.alert.SendAlert(alert)
	if err != nil {
		level.Error(ss.logger).Log("msg", "SendAlert Alert Error", "err", err)
	}
}