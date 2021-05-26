package impl

import (
	"errors"
	"github.com/flasherup/gradtage.de/apisvc"
	"github.com/flasherup/gradtage.de/apisvc/impl/utils"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/usersvc"
)

const (
	CrateAction 	= "create"
	AutoCrateAction = "auto_create"
	SetPLanAction 	= "set_plan"
	RenewAction 	= "renew"
)

var productMapToPlan  = map[string]string{
	"0": usersvc.PlanStarter,
	"1010": usersvc.PlanBasic,
	"1": usersvc.PlanAdvanced,
	"2": usersvc.PlanProfessional,
	"3": usersvc.PlanEnterprise,
}

func CreateWoocommerceUser(client usersvc.Client, email, key, planId string) error {
	plan, ok := productMapToPlan[planId]
	if !ok {
		plan = usersvc.PlanTrial
	}

	_, err := client.CreateUser(email, plan, key ,false)
	if err != nil {
		return err
	}
	return nil
}

func UpdateWoocommerceUser(client usersvc.Client, status, email, planId string, user usersvc.User) error {
	plan, ok := productMapToPlan[planId]
	if !ok {
		plan = usersvc.PlanTrial
	}

	if status == common.WCStatusTrash {
		client.DeleteUser(user)
	} else {
		user.Plan = plan
		_, err := client.UpdateUser(user, false)
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateUser(client usersvc.Client, req apisvc.ParamsUser, email bool) ([][]string, error){
	namep, ok := req.Params["name"]
	if !ok || len(namep) < 1 {
		err := errors.New("user name not set")
		return utils.CSVError(err), err
	}
	name := namep

	plan := usersvc.PlanTrial
	planp, ok := req.Params["plan"]
	if ok && len(planp) > 0 {
		plan = planp
	}

	key, err := client.CreateUser(name, plan, req.Key ,email)
	if err != nil {
		return utils.CSVError(err), err
	}

	if plan == usersvc.PlanTrial {
		params, err := client.ValidateKey(key)
		if err != nil {
			return utils.CSVError(err), err
		}

		params.User.Stations = []string{"WMO10142"}

		key, err = client.UpdateUser(params.User, false)
		if err != nil {
			return utils.CSVError(err), err
		}

	}

	return [][]string{
		{"status", "key"},
		{"ok", key},
	}, nil
}

func UpdateUser(client usersvc.Client, user usersvc.User) ([][]string, error){
	key, err := client.UpdateUser(user, true)
	if err != nil {
		return utils.CSVError(err), err
	}

	return [][]string{
		{"status", "key"},
		{"ok", key},
	}, nil
}

func SetUserPlan(client usersvc.Client, req apisvc.ParamsUser) ([][]string, error){
	namep, ok := req.Params["name"]
	if !ok || len(namep) < 1 {
		err := errors.New("user name is required")
		return utils.CSVError(err), err
	}
	name := namep

	planp, ok := req.Params["plan"]
	if !ok || len(planp) < 1 {
		err := errors.New("user plan is required")
		return utils.CSVError(err), err
	}
	plan := planp

	user, err := client.ValidateName(name)
	if err != nil {
		return utils.CSVError(err), err
	}

	user.User.Plan = plan

	key, err := client.UpdateUser(user.User, true)
	if err != nil {
		return utils.CSVError(err), err
	}

	return [][]string{
		{"status", "key"},
		{"ok", key},
	}, nil
}

func RenewUser(client usersvc.Client, req apisvc.ParamsUser) ([][]string, error){
	namep, ok := req.Params["name"]
	if !ok || len(namep) < 1 {
		err := errors.New("user name is required")
		return utils.CSVError(err), err
	}
	name := namep

	user, err := client.ValidateName(name)
	if err != nil {
		return utils.CSVError(err), err
	}

	key, err := client.UpdateUser(user.User, true)
	if err != nil {
		return utils.CSVError(err), err
	}

	return [][]string{
		{"status", "key"},
		{"ok", key},
	}, nil
}
