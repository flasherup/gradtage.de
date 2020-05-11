package impl

import (
	"context"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/usersvc"
	"github.com/flasherup/gradtage.de/usersvc/impl/database"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	ktprom "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

const keyLength = 20

type UserSVC struct {
	logger  	log.Logger
	alert 		alertsvc.Client
	db 			database.UserDB
	counter 	*ktprom.Gauge
}

func NewUserSVC(logger log.Logger, db database.UserDB, alert alertsvc.Client) (*UserSVC, error) {
	options := prometheus.Opts{
		Name: "stations_count_total",
		Help: "The total number oh stations",
	}
	guage := ktprom.NewGaugeFrom(prometheus.GaugeOpts(options), []string{ "stations" })
	st := UserSVC{
		logger:  logger,
		alert:   alert,
		db:		 db,
		counter: guage,
	}
	return &st,nil
}

func (us UserSVC) CreateUser(ctx context.Context, userName string, plan string, email bool) (string, error) {
	level.Info(us.logger).Log("msg", "Create User", "user", userName, "plan", plan)
	key, err := common.GenerateRandomString(keyLength)
	if err != nil {
		level.Error(us.logger).Log("msg", "Key generation error", "err", err)
		us.sendAlert(NewErrorAlert(err))
	}

	//TODO send email

	user := usersvc.User{
		Name:        userName,
		Key: 		key,
		RenewDate:   time.Now(),
		RequestDate: time.Now(),
		Requests:    0,
		Plan:        plan,
		Stations:    []string{},
	}
	err = us.db.SetUser(user)
	if err != nil {
		level.Error(us.logger).Log("msg", "Create User Error", "err", err)
		us.sendAlert(NewErrorAlert(err))
	}
	return key,err
}

func (us UserSVC) UpdateUser(ctx context.Context, user usersvc.User, email bool) (string, error) {
	level.Info(us.logger).Log("msg", "Update User", "user", user.Name, "email", email)
	err := us.db.SetUser(user)
	if err != nil {
		level.Error(us.logger).Log("msg", "Update User Error", "err", err)
		us.sendAlert(NewErrorAlert(err))
	}

	//TODO email update

	return user.Key,err
}

func (us UserSVC) AddPlan(ctx context.Context, plan usersvc.Plan) error {
	level.Info(us.logger).Log("msg", "Update User", "plan", plan.Name)
	err := us.db.SetPlan(plan)
	if err != nil {
		level.Error(us.logger).Log("msg", "Update User Error", "err", err)
		us.sendAlert(NewErrorAlert(err))
	}
	return err
}

func (us UserSVC) ValidateKey(ctx context.Context, key string) (usersvc.Parameters, error) {
	level.Info(us.logger).Log("msg", "Validate Key", "key", key)
	params, err := us.db.GetUserDataByKey(key)
	if err != nil {
		level.Error(us.logger).Log("msg", "Validate Key Error", "err", err)
		us.sendAlert(NewErrorAlert(err))
	}
	return params, err
}

func (us UserSVC) ValidateName(ctx context.Context, name string) (usersvc.Parameters, error) {
	level.Info(us.logger).Log("msg", "Validate Name", "name", name)
	params, err := us.db.GetUserDataByName(name)
	if err != nil {
		level.Error(us.logger).Log("msg", "Validate Name Error", "err", err)
		us.sendAlert(NewErrorAlert(err))
	}
	return params, err
}

func (hs UserSVC)sendAlert(alert alertsvc.Alert) {
	err := hs.alert.SendAlert(alert)
	if err != nil {
		level.Error(hs.logger).Log("msg", "Send Alert Error", "err", err)
	}
}