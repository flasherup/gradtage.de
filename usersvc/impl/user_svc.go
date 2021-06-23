package impl

import (
	"context"
	"errors"
	"fmt"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/usersvc"
	"github.com/flasherup/gradtage.de/usersvc/config"
	"github.com/flasherup/gradtage.de/usersvc/impl/database"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	ktprom "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)


type UserSVC struct {
	logger  	log.Logger
	alert 		alertsvc.Client
	db 			database.UserDB
	counter 	*ktprom.Gauge
	config 		config.UsersConfig
}

func NewUserSVC(logger log.Logger, db database.UserDB, alert alertsvc.Client, usersConfig config.UsersConfig) (*UserSVC, error) {
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
		config: usersConfig,
	}
	return &st,nil
}


func (us UserSVC) CreateOrder(ctx context.Context, orderId int, email, plan, key string) (string, error) {
	level.Info(us.logger).Log("msg", "Create Order", "id", orderId, "email", email, "plan", plan, "key", key)
	var err error
	if key == "" {
		key, err = common.GenerateRandomString(database.KeyLength)
		if err != nil {
			level.Error(us.logger).Log("msg", "Key generation error", "err", err)
			us.sendAlert(NewErrorAlert(err))
		}
	}

	o, _, err := us.ValidateOrder(ctx, orderId)
	if err == nil {
		return o.Key, errors.New("order already exist")
	}

	order := usersvc.Order {
		OrderId:		orderId,
		Key: 			key,
		Email:			email,
		Plan:			plan,
		Stations: 		us.getDefaultStations(plan),
		RequestDate:	time.Now(),
		Requests: 		0,
		Admin:			false,
	}

	err = us.db.SetOrder(order)
	if err != nil {
		level.Error(us.logger).Log("msg", "Create order error", "err", err)
		us.sendAlert(NewErrorAlert(err))
	}
	return key,err
}

func (us UserSVC) UpdateOrder(ctx context.Context, order usersvc.Order,) (string, error) {
	level.Info(us.logger).Log("msg", "Update order", "id", order.OrderId,)

	err := us.db.SetOrder(order)
	if err != nil {
		level.Error(us.logger).Log("msg", "Update order error", "err", err)
		us.sendAlert(NewErrorAlert(err))
	}

	return order.Key,err
}

func (us UserSVC)  DeleteOrder(ctx context.Context, orderId int) error {
	err := us.db.DeleteOrders([]int{orderId})
	if err != nil {
		level.Error(us.logger).Log("msg", "Delete order error", "err", err)
		us.sendAlert(NewErrorAlert(err))
		return err
	}

	return nil
}

func (us UserSVC) AddPlan(ctx context.Context, plan usersvc.Plan) error {
	level.Info(us.logger).Log("msg", "Update Plan", "plan", plan.Name)

	err := us.db.SetPlan(plan)
	if err != nil {
		level.Error(us.logger).Log("msg", "Update Plan Error", "err", err)
		us.sendAlert(NewErrorAlert(err))
	}
	return err
}

func (us UserSVC) ValidateSelection(ctx context.Context, selection usersvc.Selection) error {
	level.Info(us.logger).Log("msg", "Validate Selection", "key", selection.Key)
	order, err := us.db.GetOrderByKey(selection.Key)
	if err != nil {
		level.Error(us.logger).Log("msg", "Validate Key Error", "err", err.Error())
		us.sendAlert(NewErrorAlert(err))
		return err
	}

	if order.Admin {
		return nil
	}

	plan, err := us.getPlan(order.Plan)
	if err != nil {
		level.Error(us.logger).Log("msg", "Validate Key Error", "err", err.Error())
		us.sendAlert(NewErrorAlert(err))
		return err
	}

	isStartValid, err := ValidateStart(selection.Start, plan)
	if !isStartValid {
		return err
	}

	isEndValid, err := ValidateEnd(selection.End, plan)
	if !isEndValid {
		return err
	}

	stationsList, err := ValidateStationId(selection.StationID, &order, &plan)
	if err != nil {
		return err
	}
	err = ValidateOutput(selection.Method, &plan)
	if err != nil {
		return err
	}

	order.Stations = stationsList

	err = us.validateUserParameters(&order, &plan)
	if err != nil {
		return err
	}

	return err
}

func (us UserSVC) ValidateKey(ctx context.Context, key string) (usersvc.Order, usersvc.Plan, error) {
	level.Info(us.logger).Log("msg", "Validate Key", "key", key)
	order, err := us.db.GetOrderByKey(key)
	if err != nil {
		level.Error(us.logger).Log("msg", "Validate Key Error", "err", err)
		us.sendAlert(NewErrorAlert(err))
		return order, usersvc.Plan{}, err
	}

	plan, err := us.getPlan(order.Plan)
	if err != nil {
		level.Error(us.logger).Log("msg", "Validate Key Error", "err", err)
		us.sendAlert(NewErrorAlert(err))
		return order, usersvc.Plan{}, err
	}

	return order, plan, us.validateUserParameters(&order, &plan)
}

func (us UserSVC) ValidateOrder(ctx context.Context, orderId int) (usersvc.Order, usersvc.Plan, error) {
	level.Info(us.logger).Log("msg", "Validate Order", "orderId", orderId)
	order, err := us.db.GetOrderById(orderId)
	if err != nil {
		level.Error(us.logger).Log("msg", "Validate order error", "err", err)
		us.sendAlert(NewErrorAlert(err))
		return order, usersvc.Plan{}, err
	}

	plan, err := us.getPlan(order.Plan)
	if err != nil {
		level.Error(us.logger).Log("msg", "Validate order Error", "err", err)
		us.sendAlert(NewErrorAlert(err))
		return order, usersvc.Plan{}, err
	}

	return order, plan, us.validateUserParameters(&order, &plan)
}

func (us UserSVC)validateUserParameters(order *usersvc.Order, plan *usersvc.Plan) error {
	if order.Admin {
		return nil
	}

	if plan.Name == usersvc.PlanCanceled {
		return errors.New("plan is canceled")
	}

	requests, err := ValidateRequestsAvailable(order, plan)
	if err != nil {
		return err
	}

	//Update user request time nad count
	order.RequestDate = time.Now().UTC()
	order.Requests = requests
	err = us.db.SetOrder(*order)
	if err != nil {
		level.Error(us.logger).Log("msg", "Update order request time and count error", "err", err)
		us.sendAlert(NewErrorAlert(err))
	}
	return nil
}

func (us UserSVC)sendAlert(alert alertsvc.Alert) {
	err := us.alert.SendAlert(alert)
	if err != nil {
		level.Error(us.logger).Log("msg", "SendAlert Alert Error", "err", err)
	}
}

func (us UserSVC)getDefaultStations (sType string) []string {
	if sType == usersvc.PlanTrial {
		return []string{us.config.Plans.FreeDefault}
	}
	return []string{}
}

func (us UserSVC) getPlan(name string) (usersvc.Plan, error) {
	plans, err := us.db.GetPlans([]string{name})
	if err != nil {
		return usersvc.Plan{}, err
	}

	if len(plans) != 1 {
		return usersvc.Plan{}, fmt.Errorf("plan:%s not find", name)
	}

	plan := plans[0]

	return plan, nil
}