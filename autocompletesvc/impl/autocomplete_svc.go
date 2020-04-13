package impl

import (
	"context"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/autocompletesvc"
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

func (ss AutocompleteSVC) GetAutocomplete(ctx context.Context, text string) (result map[string][]autocompletesvc.Source, err error) {
	level.Info(ss.logger).Log("msg", "GetAutocomplete", "text", text)
	result, err = ss.db.GetAutocomplete(text)
	if err != nil {
		level.Error(ss.logger).Log("msg", "GetAutocomplete DB error", "err", err)
		ss.sendAlert(NewErrorAlert(err))
		return nil,err
	}
	return result, nil
}

func (ss AutocompleteSVC) AddSources(ctx context.Context, sources []autocompletesvc.Source) (err error) {
	level.Info(ss.logger).Log("msg", "AddSource", "length", len(sources))
	err = ss.db.AddSources(sources)
	if err != nil {
		level.Error(ss.logger).Log("msg", "AddSources DB error", "err", err)
		ss.sendAlert(NewErrorAlert(err))
		return err
	}
	return nil
}

func (ss AutocompleteSVC)sendAlert(alert alertsvc.Alert) {
	err := ss.alert.SendAlert(alert)
	if err != nil {
		level.Error(ss.logger).Log("msg", "Send Alert Error", "err", err)
	}
}