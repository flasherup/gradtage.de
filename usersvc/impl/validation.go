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

func ValidatePlanExpiration(params *usersvc.Parameters) error {
	current := time.Now().UTC()
	period := params.User.RenewDate.Sub(current)/(time.Hour*24)
	fmt.Println("period", period)
	if int(period) >= params.Plan.Period {
		return fmt.Errorf("the key: '%s' is expired", params.User.Key)
	}
	return nil
}

func ValidateRequestsAvailable(params *usersvc.Parameters) (int, error) {
	current := time.Now().UTC()
	count := params.User.Requests
	dif := params.User.RequestDate.Sub(current)
	if dif < time.Hour {
		count++
	} else {
		count = 1
	}

	if count > params.Plan.Limitation {
		return count,  fmt.Errorf("requests limit is exceeded")
	}

	return count, nil
}