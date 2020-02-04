package parser

import (
	"github.com/flasherup/gradtage.de/hourlysvc"
)

type ParsedData struct {
	Success bool
	StationID string
	Temps []hourlysvc.Temperature
	Error error
}
