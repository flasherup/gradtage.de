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
	sub := current.Sub(params.User.RenewDate)
	period := sub.Hours()/24
	if int(period) >= params.Plan.Period {
		return fmt.Errorf("the key: '%s' for user: '%s' is expired", params.User.Key,  params.User.Name)
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
func ValidateStationId(stationId string, params *usersvc.Parameters) error {
	if  !isIDExist(params.User.Stations, stationId) {
		return fmt.Errorf("invalid station request")
	}

	return nil
}

func ValidateStationsCount(stationId string, params *usersvc.Parameters) error {
	if params.User.Stations != nil &&
		params.Plan.Stations == len(params.User.Stations) {
		return fmt.Errorf("all stations are seted")
	}
	return nil
}

func isIDExist(ids []string, stationId string) bool {
	for i, _ := range ids {
		if ids[i] == stationId {
			return true
		}
	}
	return false
}