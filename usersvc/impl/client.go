package impl

import (
	"context"
	"errors"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/usersvc"
	"github.com/flasherup/gradtage.de/usersvc/grpcusr"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type UsersSVCClient struct{
	logger     	log.Logger
	host 		string
}

func NewUsersSCVClient(host string, logger log.Logger) *UsersSVCClient {
	logger = log.With(logger,
		"client", "users",
	)
	return &UsersSVCClient{
		logger:logger,
		host: host,
	}
}

func (us UsersSVCClient) CreateOrder(orderId int, email, plan, key string) (string, error) {
	conn, err := common.OpenGRPCConnection(us.host)
	if err != nil {
		return common.ErrorNilString,err
	}
	defer conn.Close()


	client := grpcusr.NewUserSVCClient(conn)
	resp, err := client.CreateOrder(context.Background(), &grpcusr.CreateOrderRequest{
		OrderId: int32(orderId),
		Email: email,
		Plan: plan,
		Key: key,
	})

	if err != nil {
		level.Error(us.logger).Log("msg", "Failed to create Order", "err", err.Error())
		return "", err
	}

	if resp.Err != common.ErrorNilString {
		err = errors.New(resp.Err)
		return resp.Key, err
	}
	return resp.Key, err
}

func (us UsersSVCClient) UpdateOrder(Order usersvc.Order) (string, error) {
	conn, err := common.OpenGRPCConnection(us.host)
	if err != nil {
		return common.ErrorNilString,err
	}
	defer conn.Close()

	client := grpcusr.NewUserSVCClient(conn)
	u := usersvc.EncodeOrder(&Order)
	resp, err := client.UpdateOrder(context.Background(), &grpcusr.UpdateOrderRequest{
		Order:  u,
	})

	if err != nil {
		level.Error(us.logger).Log("msg", "Failed to create Order", "err", err.Error())
	}else if resp.Err != common.ErrorNilString {
		err = errors.New(resp.Err)
	}

	return resp.Key, err
}

func (us UsersSVCClient) DeleteOrder(OrderId int) error {
	conn, err := common.OpenGRPCConnection(us.host)
	if err != nil {
		return err
	}
	defer conn.Close()

	client := grpcusr.NewUserSVCClient(conn)
	resp, err := client.DeleteOrder(context.Background(), &grpcusr.DeleteOrderRequest{
		OrderId:  int32(OrderId),
	})

	if err != nil {
		level.Error(us.logger).Log("msg", "Failed to delete Order", "err", err.Error())
	}else if resp.Err != common.ErrorNilString {
		err = errors.New(resp.Err)
	}

	return err
}

//AddPlan(plan Plan) error
func (us UsersSVCClient) AddPlan(plan usersvc.Plan) error {
	conn, err := common.OpenGRPCConnection(us.host)
	if err != nil {
		return err
	}
	defer conn.Close()

	client := grpcusr.NewUserSVCClient(conn)
	p := usersvc.EncodePlan(&plan)
	resp, err := client.AddPlan(context.Background(), &grpcusr.AddPlanRequest{
		Plan:  p,
	})

	if err != nil {
		level.Error(us.logger).Log("msg", "Failed to create user", "err", err.Error())
	}else if resp.Err != common.ErrorNilString {
		err = errors.New(resp.Err)
	}

	return err
}

//ValidateSelection(selection Selection) (bool, error)
func (us UsersSVCClient) ValidateSelection(selection usersvc.Selection) error {
	conn, err := common.OpenGRPCConnection(us.host)
	if err != nil {
		return err
	}
	defer conn.Close()

	client := grpcusr.NewUserSVCClient(conn)
	p := usersvc.EncodeSelection(&selection)
	resp, err := client.ValidateSelection(context.Background(), &grpcusr.ValidateSelectionRequest{
		Selection:  p,
	})

	if err != nil {
		level.Error(us.logger).Log("msg", "Failed to validate selection", "err", err.Error())
	}else if resp.Err != common.ErrorNilString {
		err = errors.New(resp.Err)
	}

	return err
}

//ValidateKey(key string) (Parameters, error)
func (us UsersSVCClient) ValidateKey(key string) (usersvc.Order, usersvc.Plan, error) {
	conn, err := common.OpenGRPCConnection(us.host)
	if err != nil {
		return usersvc.Order{}, usersvc.Plan{}, err
	}
	defer conn.Close()

	client := grpcusr.NewUserSVCClient(conn)
	resp, validateError := client.ValidateKey(context.Background(), &grpcusr.ValidateKeyRequest{
		Key:  key,
	})

	if validateError != nil {
		return usersvc.Order{}, usersvc.Plan{}, validateError
	}

	order, decodeOrderError := usersvc.DecodeOrder(resp.Order)
	if decodeOrderError != nil {
		level.Error(us.logger).Log("msg", "Failed to validate key", "err", err)
		return *order, usersvc.Plan{}, decodeOrderError
	}

	plan, decodePlanError := usersvc.DecodePlan(resp.Plan)
	if decodePlanError != nil {
		level.Error(us.logger).Log("msg", "Failed to validate key", "err", err)
		return *order, *plan, decodePlanError
	}

	if resp.Err != common.ErrorNilString {
		return *order, *plan, errors.New(resp.Err)
	}


	return *order, *plan, nil
}


//ValidateName(name string) (Parameters, error)
func (us UsersSVCClient) ValidateOrder(orderId int) (usersvc.Order, usersvc.Plan, error) {
	conn, err := common.OpenGRPCConnection(us.host)
	if err != nil {
		return usersvc.Order{}, usersvc.Plan{}, err
	}
	defer conn.Close()

	client := grpcusr.NewUserSVCClient(conn)
	resp, validateError := client.ValidateOrder(context.Background(), &grpcusr.ValidateOrderRequest{
		OrderId:  int32(orderId),
	})

	if validateError != nil {
		return usersvc.Order{}, usersvc.Plan{}, validateError
	}

	order, decodeOrderError := usersvc.DecodeOrder(resp.Order)
	if decodeOrderError != nil {
		level.Error(us.logger).Log("msg", "Failed to validate key", "err", err)
		return *order, usersvc.Plan{}, decodeOrderError
	}

	plan, decodePlanError := usersvc.DecodePlan(resp.Plan)
	if decodePlanError != nil {
		level.Error(us.logger).Log("msg", "Failed to validate key", "err", err)
		return *order, *plan, decodePlanError
	}

	if resp.Err != common.ErrorNilString {
		return *order, *plan, errors.New(resp.Err)
	}


	return *order, *plan, nil
}