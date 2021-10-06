package main

import (
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/localutils/data"
	"github.com/flasherup/gradtage.de/usersvc"
	"github.com/flasherup/gradtage.de/usersvc/impl"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"
	"time"
)

const testKey = "justtestkeynothingspecial"
const testOrderId = 102

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
	client := impl.NewUsersSCVClient("212.227.215.17:8110",logger)//New server

	level.Info(logger).Log("msg", "client started")
	defer level.Info(logger).Log("msg", "client ended")

	addPlans(client, logger, data.Plans)

	//createOrder(client, logger)
	//validateOrder(client, logger, 1277)
	//validateKey(client, logger, "fe52a98a-3b04-41d3-bd35-79a31ec3e9b2")
	/*for i := 0; i<10; {
		i++
		validateSelection(client, logger)
	}*/
	//validateSelection(client, logger)
	//updateOrder(client, logger)
	//deleteOrder(client, logger)

}

func createOrder(client *impl.UsersSVCClient, logger log.Logger) {
	//key, err := client.CreateOrder(testOrderId, "test@test.test", "trial", testKey)
	key, err := client.CreateOrder(10000010, "admin@gradtage.de", usersvc.PlanEnterprise, "hSlSfSlHmfJtpGqCmfLR")
	if err !=nil {
		level.Error(logger).Log("msg", "Order Create error", "err", err.Error())
	}
	level.Info(logger).Log("msg", "Order created successfully", "key", key)
}

func validateOrder(client *impl.UsersSVCClient, logger log.Logger, orderId int) {
	order, plan, err := client.ValidateOrder(orderId)
	if err != nil {
		level.Error(logger).Log("msg", "Order validation error", "err", err.Error())
	}

	level.Info(logger).Log(
		"msg", "Order validated successfully",
		"orderId", order.OrderId,
		"plan", plan.Name)
}

func validateKey(client *impl.UsersSVCClient, logger log.Logger, key string) {
	order, plan, err := client.ValidateKey(key)
	if err != nil {
		level.Error(logger).Log("msg", "Key validation error", "err", err.Error())
	}

	level.Info(logger).Log(
		"msg", "Order validated successfully",
		"orderId", order.OrderId,
		"plan", plan.Name)
}

func validateSelection(client *impl.UsersSVCClient, logger log.Logger) {
	start, err := time.Parse(common.TimeLayoutWBH, "2012-01-01")
	if err != nil {
		level.Error(logger).Log("msg", "Start time validation error", "err", err)
		return
	}

	end, err := time.Parse(common.TimeLayoutWBH, "2012-02-01")
	if err != nil {
		level.Error(logger).Log("msg", "End time validation error", "err", err)
		return
	}

	selection := usersvc.Selection{
		//Key:       testKey,
		Key: 		"fe52a98a-3b04-41d3-bd35-79a31ec3e9b2",
		StationID: 	"EDDH",
		Method:    	common.HDDType,
		Start:     	start,
		End:       	end,
	}

	err = client.ValidateSelection(selection)
	if err != nil {
		level.Error(logger).Log("msg", "Selection validation error", "err", err.Error())
		return
	}

	level.Info(logger).Log("msg", "Selection validated successfully")
}


func updateOrder(client *impl.UsersSVCClient, logger log.Logger) {
	order, _, err := client.ValidateOrder(10000000)
	if err != nil {
		level.Error(logger).Log("msg", "Order update error", "err", err.Error())
	}

	order.Admin = true
	order.Plan = usersvc.PlanAdmin
	key, err := client.UpdateOrder(order)
	if err != nil {
		level.Error(logger).Log("msg", "Order update error", "err", err.Error())
	}

	level.Info(logger).Log(
		"msg", "Order updated successfully",
		"key", key)
}

func deleteOrder(client *impl.UsersSVCClient, logger log.Logger) {
	err := client.DeleteOrder(0)
	if err !=nil {
		level.Error(logger).Log("msg", "Order delete error", "err", err.Error())
	}
	level.Info(logger).Log("msg", "Order delete successfully", "orderId", testOrderId)
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
