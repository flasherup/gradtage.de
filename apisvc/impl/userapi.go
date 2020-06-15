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

func ProcessUpdateStripeUser(client usersvc.Client, name string, stripeId string, plan string) (string, error) {
	user, err := client.ValidateStripe(stripeId)
	if err != nil {
		user, err = client.ValidateName(name)
		if err != nil {
			key, err := client.CreateUser(name, plan, true)
			if err != nil {
				return "{\"status\":\"error\", \"error\":\"" + err.Error() + "\"}", err
			}

			user, err = client.ValidateKey(key)
			if err != nil {
				return "{\"status\":\"error\", \"error\":\"" + err.Error() + "\"}", err
			}
		}
	}

	user.User.Plan = plan
	user.User.Stripe = stripeId

	_, err = client.UpdateUser(user.User, true)
	if err != nil {
		return "{\"status\":\"error\", \"error\":\"" + err.Error() + "\"}", err
	}

	return "{\"status\":\"ok\"}", nil
}

func ProcessCancelStripeUser(client usersvc.Client, stripeId string) (string, error) {
	user, err := client.ValidateStripe(stripeId)
	if err != nil {
		return "{\"status\":\"error\", \"error\":\"" + err.Error() + "\"}", err
	}

	user.User.Plan = usersvc.PlanTrial
	user.User.Stripe = stripeId

	_, err = client.UpdateUser(user.User, true)
	if err != nil {
		return "{\"status\":\"error\", \"error\":\"" + err.Error() + "\"}", err
	}

	return "{\"status\":\"ok\"}", nil
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

	key, err := client.CreateUser(name, plan,email)
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
