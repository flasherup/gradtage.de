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


func (us UsersSVCClient) CreateUser(userName string, plan string, key string, email bool) (string, error) {
	conn, err := common.OpenGRPCConnection(us.host)
	if err != nil {
		return common.ErrorNilString,err
	}
	defer conn.Close()


	client := grpcusr.NewUserSVCClient(conn)
	resp, err := client.CreateUser(context.Background(), &grpcusr.CreateUserRequest{
		UserName: userName,
		Plan: plan,
		Key: key,
		Email: email,
	})

	if err != nil {
		level.Error(us.logger).Log("msg", "Failed to create user", "err", err.Error())
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
		level.Error(us.logger).Log("msg", "Failed to create user", "err", err.Error())
	}else if resp.Err != common.ErrorNilString {
		err = errors.New(resp.Err)
	}

	return resp.Key, err
}

func (us UsersSVCClient) DeleteUser(user usersvc.User) error {
	conn, err := common.OpenGRPCConnection(us.host)
	if err != nil {
		return err
	}
	defer conn.Close()

	client := grpcusr.NewUserSVCClient(conn)
	u := usersvc.EncodeUser(&user)
	resp, err := client.DeleteUser(context.Background(), &grpcusr.DeleteUserRequest{
		User:  u,
	})

	if err != nil {
		level.Error(us.logger).Log("msg", "Failed to delete user", "err", err.Error())
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
func (us UsersSVCClient) ValidateKey(key string) (usersvc.Parameters, error) {
	conn, err := common.OpenGRPCConnection(us.host)
	if err != nil {
		return usersvc.Parameters{}, err
	}
	defer conn.Close()

	client := grpcusr.NewUserSVCClient(conn)
	resp, validateError := client.ValidateKey(context.Background(), &grpcusr.ValidateKeyRequest{
		Key:  key,
	})

	if validateError != nil {
		return usersvc.Parameters{}, validateError
	}

	p, decodeError := usersvc.DecodeParameters(resp.Parameters)
	if decodeError != nil {
		level.Error(us.logger).Log("msg", "Failed to validate selection", "err", err)
		return *p, decodeError
	}

	if resp.Err != common.ErrorNilString {
		return *p, errors.New(resp.Err)
	}


	return *p, nil
}


//ValidateName(name string) (Parameters, error)
func (us UsersSVCClient) ValidateName(name string) (usersvc.Parameters, error) {
	conn, err := common.OpenGRPCConnection(us.host)
	if err != nil {
		return usersvc.Parameters{}, err
	}
	defer conn.Close()

	client := grpcusr.NewUserSVCClient(conn)
	resp, validateError := client.ValidateName(context.Background(), &grpcusr.ValidateNameRequest{
		Name:  name,
	})

	if validateError != nil {
		return usersvc.Parameters{}, validateError
	}

	p, decodeError := usersvc.DecodeParameters(resp.Parameters)
	if decodeError != nil {
		level.Error(us.logger).Log("msg", "Failed to validate selection", "err", err.Error())
		return *p, decodeError
	}

	if resp.Err != common.ErrorNilString {
		return *p,errors.New(resp.Err)
	}
	return *p, nil
}