package impl

import (
	"fmt"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/usersvc"
	"time"
)

func ValidateStart( start time.Time, plan usersvc.Plan) (bool, error) {
	if start.Sub(plan.Start) < 0 {
		return false, fmt.Errorf("strat date is invalid")
	}
	return true, nil
}

func ValidateEnd( end time.Time, plan usersvc.Plan) (bool, error) {
	if plan.End.Sub(end) < 0 {
		return false, fmt.Errorf("end date is invalid")
	}
	return true, nil
}

func ValidateRequestsAvailable(order *usersvc.Order, plan *usersvc.Plan) (int, error) {
	current := time.Now().UTC()
	count := order.Requests
	dif := order.RequestDate.Sub(current)
	if dif < time.Hour {
		count++
	} else {
		count = 1
	}

	if count > plan.Limitation {
		return count,  fmt.Errorf("requests limit is exceeded")
	}

	return count, nil
}
func ValidateStationId(stationId string, order *usersvc.Order, plan *usersvc.Plan) ([]string, error) {
	stationsList := order.Stations
	if  !isIDExist(stationsList, stationId) {
		if plan.Stations <= len(stationsList) {
			return order.Stations, fmt.Errorf("station list is full")
		}
		stationsList = append(order.Stations, stationId)
	}

	return stationsList, nil
}

func ValidateOutput(output string, plan *usersvc.Plan)  error {
	if output == common.HDDType && plan.HDD  {
		return nil
	}

	if output == common.DDType && plan.DD  {
		return nil
	}

	if output == common.CDDType && plan.CDD  {
		return nil
	}

	return fmt.Errorf("output: %s is not ollowed in this product", output)
}

func isIDExist(ids []string, stationId string) bool {
	for i, _ := range ids {
		if ids[i] == stationId {
			return true
		}
	}
	return false
}