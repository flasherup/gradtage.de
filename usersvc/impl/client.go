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
	return &UsersSVCClient{
		logger:logger,
		host: host,
	}
}


func (us UsersSVCClient) CreateUser(userName string, plan string, email bool) (string, error) {
	conn, err := common.OpenGRPCConnection(us.host)
	if err != nil {
		return common.ErrorNilString,err
	}
	defer conn.Close()


	client := grpcusr.NewUserSVCClient(conn)
	resp, err := client.CreateUser(context.Background(), &grpcusr.CreateUserRequest{
		UserName: userName,
		Plan: plan,
		Email: email,
	})

	if err != nil {
		level.Error(us.logger).Log("msg", "Failed to create user", "err", err)
		return "", err
	}

	if resp.Err != common.ErrorNilString {
		err = errors.New(resp.Err)
		return resp.Key, err
	}
	return resp.Key, err
}

func (us UsersSVCClient) UpdateUser(user usersvc.User, email bool) (string, error) {
	conn, err := common.OpenGRPCConnection(us.host)
	if err != nil {
		return common.ErrorNilString,err
	}
	defer conn.Close()

	client := grpcusr.NewUserSVCClient(conn)
	u := usersvc.EncodeUser(&user)
	resp, err := client.UpdateUser(context.Background(), &grpcusr.UpdateUserRequest{
		User:  u,
		Email: email,
	})

	if err != nil {
		level.Error(us.logger).Log("msg", "Failed to create user", "err", err)
	}else if resp.Err != common.ErrorNilString {
		err = errors.New(resp.Err)
	}

	return resp.Key, err
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
		level.Error(us.logger).Log("msg", "Failed to create user", "err", err)
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
		level.Error(us.logger).Log("msg", "Failed to validate selection", "err", err)
	}else if resp.Err != common.ErrorNilString {
		err = errors.New(resp.Err)
	}

	return err
}

//ValidateKey(key string) (Parameters, error)
func (us UsersSVCClient) ValidateKey(key string) (usersvc.Parameters, error) {
	conn, err := common.OpenGRPCConnection(us.host)
	if err != nil {
		return usersvc.Parameters{}, err
	}
	defer conn.Close()

	client := grpcusr.NewUserSVCClient(conn)
	resp, err := client.ValidateKey(context.Background(), &grpcusr.ValidateKeyRequest{
		Key:  key,
	})

	if err != nil {
		level.Error(us.logger).Log("msg", "Failed to validate selection", "err", err)
		return usersvc.Parameters{}, err
	}else if resp.Err != common.ErrorNilString {
		err = errors.New(resp.Err)
		return usersvc.Parameters{}, err
	}

	p, err := usersvc.DecodeParameters(resp.Parameters)
	return *p, err
}


//ValidateName(name string) (Parameters, error)
func (us UsersSVCClient) ValidateName(name string) (usersvc.Parameters, error) {
	conn, err := common.OpenGRPCConnection(us.host)
	if err != nil {
		return usersvc.Parameters{}, err
	}
	defer conn.Close()

	client := grpcusr.NewUserSVCClient(conn)
	resp, err := client.ValidateName(context.Background(), &grpcusr.ValidateNameRequest{
		Name:  name,
	})

	if err != nil {
		level.Error(us.logger).Log("msg", "Failed to validate selection", "err", err)
		return usersvc.Parameters{},err
	}else if resp.Err != common.ErrorNilString {
		err = errors.New(resp.Err)
		return usersvc.Parameters{},err
	}

	p, err := usersvc.DecodeParameters(resp.Parameters)
	return *p, err
}

//ValidateName(name string) (Parameters, error)
func (us UsersSVCClient) ValidateStripe(stripe string) (usersvc.Parameters, error) {
	conn, err := common.OpenGRPCConnection(us.host)
	if err != nil {
		return usersvc.Parameters{}, err
	}
	defer conn.Close()

	client := grpcusr.NewUserSVCClient(conn)
	resp, err := client.ValidateStripe(context.Background(), &grpcusr.ValidateStripeRequest{
		Stripe:  stripe,
	})

	if err != nil {
		level.Error(us.logger).Log("msg", "Failed to validate stripe", "err", err)
		return  usersvc.Parameters{}, err
	}else if resp.Err != common.ErrorNilString {
		err = errors.New(resp.Err)
		return  usersvc.Parameters{}, err
	}

	p, err := usersvc.DecodeParameters(resp.Parameters)
	return *p, err
}