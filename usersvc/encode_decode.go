package usersvc

import (
	"context"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/usersvc/grpcusr"
	"time"
)


func EncodeCreateUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(CreateUserResponse)
	return &grpcusr.CreateUserResponse {
		Err: errorToString(res.Err),
	}, nil
}

func DecodeCreateUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*grpcusr.CreateUserRequest)
	plan, err := DecodePlan(req.Plan)
	return CreateUserRequest{req.UserName, *plan}, err
}

func EncodeCreateUserAutoResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(CreateUserAutoResponse)
	return &grpcusr.CreateUserAutoResponse {
		Err: errorToString(res.Err),
	}, nil
}

func DecodeCreateUserAutoRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*grpcusr.CreateUserAutoRequest)
	plan, err := DecodePlan(req.Plan)
	return CreateUserAutoRequest{req.UserName, *plan}, err
}

func EncodeSetPlanResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(SetPlanResponse)
	return &grpcusr.SetPlanResponse {
		Err: errorToString(res.Err),
	}, nil
}

func DecodeSetPlanRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*grpcusr.SetPlanRequest)
	plan, err := DecodePlan(req.Plan)
	return SetPlanRequest{req.UserName, *plan}, err
}

func EncodeSetStationsResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(SetStationsResponse)
	return &grpcusr.SetStationsResponse {
		Err: errorToString(res.Err),
	}, nil
}

func DecodeSetStationsRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*grpcusr.SetStationsRequest)
	return SetStationsRequest{req.UserName, req.Stations}, nil
}

func EncodeValidateKeyResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(ValidateKeyResponse)
	return &grpcusr.ValidateKeyResponse {
		Err: errorToString(res.Err),
	}, nil
}

func DecodeValidateKeyRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*grpcusr.ValidateKeyRequest)
	return ValidateKeyRequest{req.Key}, nil
}

func DecodePlan(src *grpcusr.Plan) (*Plan, error) {
	start, err := time.Parse(src.Start, common.TimeLayout)
	if err != nil {
		return nil, err
	}
	end, err := time.Parse(src.End, common.TimeLayout)
	if err != nil {
		return nil, err
	}
	return &Plan{
		Name:  src.Name,
		Start: start,
		End:   end,
		Admin: src.Admin,
	}, nil
}

func EncodePlan(src *Plan) (*grpcusr.Plan) {
	start := src.Start.Format(common.TimeLayout)
	end := src.End.Format(common.TimeLayout)
	return &grpcusr.Plan{
		Name:  src.Name,
		Start: start,
		End:   end,
		Admin: src.Admin,
	}
}

func errorToString(err error) string{
	if err == nil {
		return "nil"
	}

	return err.Error()
}


