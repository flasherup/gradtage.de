package impl

import (
	"errors"
	"github.com/flasherup/gradtage.de/apisvc"
	"github.com/flasherup/gradtage.de/apisvc/impl/utils"
	"github.com/flasherup/gradtage.de/usersvc"
)

const (
	CrateAction 	= "create"
	AutoCrateAction = "auto_create"
	SetPLanAction 	= "set_plan"
	RenewAction 	= "renew"
)

func CreateUser(client usersvc.Client, req apisvc.ParamsUser, email bool) ([][]string, error){
	namep, ok := req.Params["name"]
	if !ok || len(namep) < 1 {
		err := errors.New("user name not set")
		return utils.CSVError(err), err
	}
	name := namep[0]

	plan := usersvc.PlanTrial
	planp, ok := req.Params["plan"]
	if ok && len(namep) > 0 {
		plan = planp[0]
	}

	key, err := client.CreateUser(name, plan,email)
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
	name := namep[0]

	planp, ok := req.Params["plan"]
	if !ok || len(planp) < 1 {
		err := errors.New("user plan is required")
		return utils.CSVError(err), err
	}
	plan := planp[0]

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
	name := namep[0]

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
