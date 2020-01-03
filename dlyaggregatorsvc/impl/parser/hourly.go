package parser

import "github.com/flasherup/gradtage.de/dailysvc"

type StationDaily struct {
	ID string
	Temps []dailysvc.Temperature
}