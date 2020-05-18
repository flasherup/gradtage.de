package impl

import (
	"fmt"
	"github.com/flasherup/gradtage.de/usersvc"
	"time"
)

func ValidateStart( start time.Time, params usersvc.Parameters) (bool, error) {
	if params.Plan.Start.Sub(start) < 0 {
		return false, fmt.Errorf("strat is invalid")
	}
	return true, nil
}

func ValidateEnd( end time.Time, params usersvc.Parameters) (bool, error) {
	if params.Plan.End.Sub(end) < 0 {
		return false, fmt.Errorf("end is invalid")
	}
	return true, nil
}