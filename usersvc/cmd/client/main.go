package main

import (
	"github.com/flasherup/gradtage.de/localutils/data"
	"github.com/flasherup/gradtage.de/usersvc"
	"github.com/flasherup/gradtage.de/usersvc/impl"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "usersvc",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	//client := impl.NewUsersSCVClient("localhost:8110",logger)
	//client := impl.NewUsersSCVClient("82.165.18.228:8110",logger)//Old Server
	client := impl.NewUsersSCVClient("212.227.214.163:8110",logger)//New server

	level.Info(logger).Log("msg", "client started")
	defer level.Info(logger).Log("msg", "client ended")

	addPlans(client, logger, data.Plans)
	//checkUsers(client, logger)
	//createUser(client, logger)

}

func checkUsers(client *impl.UsersSVCClient, logger log.Logger) error {

	params, err := client.ValidateKey(data.UserKeys["admin"]);
	if err != nil {
		level.Error(logger).Log("msg", "Admin user check error", "err", err.Error())
	} else {
		level.Info(logger).Log("msg", "Admin user check ok", "Plan name", params.Plan.Name, "User", params.User.Name);
	}

	params, err = client.ValidateKey(data.UserKeys["starter"]);
	if err != nil {
		level.Info(logger).Log("msg", "Starter user check ok", "msg", "User should be expired")
		return err;
	} else {
		level.Error(logger).Log("msg", "Starter user check error", "err", "date is not expired", "date", params.User.RenewDate.String())

	}
	params, err = client.ValidateKey(data.UserKeys["trial"]);
	if err != nil {
		level.Info(logger).Log("msg", "Trial user check ok", "msg", "User id expired", "err", err.Error())
	} else {
		level.Error(logger).Log("msg", "Trial user check error", "err", "date is not expired", "date", params.User.RenewDate.String())
	}


	return nil
}

func addPlans(client *impl.UsersSVCClient, logger log.Logger, plans []usersvc.Plan) {
	for _, plan := range plans {
		err := client.AddPlan(plan)
		if err != nil {
			level.Error(logger).Log("msg", "Add Plan Error", "err", err)
			break
		}
	}
}

func createUser(client *impl.UsersSVCClient, logger log.Logger) {
	key, err := client.CreateUser("flasherup@gmail.com", "trial",  "85c08bd3-fe7f-4174-b294-efed0f1a2e52", false)
	if err != nil {
		level.Error(logger).Log("msg", "Create User Error", "err", err)
	}

	key, err = client.CreateUser("admin@gradtage.de", "trial", "", false)
	if err != nil {
		level.Error(logger).Log("msg", "Create User Error", "err", err)
	}
	level.Info(logger).Log("msg", "Users Created", "key", key)

	/*params, err := client.ValidateName("flasherup@gmail.com")
	if err != nil {
		level.Error(logger).Log("msg", "Validate User Error", "err", err)
	}
	level.Info(logger).Log("msg", "User Valid", "key", params.User.Key)
	*/

	params, err := client.ValidateName("admin@gradtage.de")
	if err != nil {
		level.Error(logger).Log("msg", "Validate User Error", "err", err)
	}
	level.Info(logger).Log("msg", "User Valid", "name", params.User.Name)

	user := params.User
	user.Plan = "admin"
	key, err = client.UpdateUser(user, false)
	if err != nil {
		level.Error(logger).Log("msg", "Update User Error", "err", err)
	}
	level.Info(logger).Log("msg", "User Updated", "name", key)
}
