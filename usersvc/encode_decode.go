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
		Key: res.Key,
		Err: errorToString(res.Err),
	}, nil
}

func DecodeCreateUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*grpcusr.CreateUserRequest)
	return CreateUserRequest{req.UserName, req.Plan, req.Email}, nil
}

func EncodeUpdateUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(UpdateUserResponse)
	return &grpcusr.UpdateUserResponse {
		Key: res.Key,
		Err: errorToString(res.Err),
	}, nil
}

func DecodeUpdateUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*grpcusr.UpdateUserRequest)
	user, err := DecodeUser(req.User)
	if err != nil {
		return UpdateUserRequest{}, err
	}
	return UpdateUserRequest{*user, req.Email}, nil
}

func EncodeAddPlanResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(AddPlanResponse)
	return &grpcusr.AddPlanResponse {
		Err: errorToString(res.Err),
	}, nil
}

func DecodeAddPlanRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*grpcusr.AddPlanRequest)
	plan, err := DecodePlan(req.Plan)
	return AddPlanRequest{*plan}, err
}

func EncodeValidateSelectionResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(ValidateSelectionResponse)
	return &grpcusr.ValidateSelectionResponse {
		IsValid: res.IsValid,
		Err: errorToString(res.Err),
	}, nil
}

func DecodeValidateSelectionRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*grpcusr.ValidateSelectionRequest)
	selection, err := DecodeSelection(req.Selection)
	return ValidateSelectionRequest{*selection}, err
}

func EncodeValidateKeyResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(ValidateKeyResponse)
	params:= EncodeParameters(&res.Parameters)
	return &grpcusr.ValidateKeyResponse {
		Parameters: params,
		Err: errorToString(res.Err),
	}, nil
}

func DecodeValidateKeyRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*grpcusr.ValidateKeyRequest)
	return ValidateKeyRequest{req.Key}, nil
}

func EncodeValidateNameResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(ValidateNameResponse)
	params:= EncodeParameters(&res.Parameters)
	return &grpcusr.ValidateNameResponse {
		Parameters: params,
		Err: errorToString(res.Err),
	}, nil
}

func DecodeValidateNameRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*grpcusr.ValidateNameRequest)
	return ValidateNameRequest{req.Name}, nil
}

func DecodePlan(src *grpcusr.Plan) (*Plan, error) {
	start, err := time.Parse(common.TimeLayout, src.Start)
	if err != nil {
		return nil, err
	}
	end, err := time.Parse(common.TimeLayout, src.End)
	if err != nil {
		return nil, err
	}
	return &Plan{
		Name:  		src.Name,
		Stations: 	int(src.Stations),
		Limitation: int(src.Limitation),
		HDD: 		src.Hdd,
		DD: 		src.Dd,
		CDD: 		src.Cdd,
		Start: 		start,
		End:   		end,
		Period: 	int(src.Period),
		Admin: src.Admin,
	}, nil
}

func EncodePlan(src *Plan) (*grpcusr.Plan) {
	start := src.Start.Format(common.TimeLayout)
	end := src.End.Format(common.TimeLayout)
	return &grpcusr.Plan{
		Name:  		src.Name,
		Stations: 	int32(src.Stations),
		Limitation: int32(src.Limitation),
		Hdd: 		src.HDD,
		Dd: 		src.DD,
		Cdd: 		src.CDD,
		Start: 		start,
		End:   		end,
		Period: 	int32(src.Period),
		Admin: src.Admin,
	}
}

func DecodeUser(src *grpcusr.User) (*User, error) {
	renew, err := time.Parse(common.TimeLayout, src.RenewDate)
	if err != nil {
		return nil, err
	}
	requests, err := time.Parse(common.TimeLayout, src.RequestDate)
	if err != nil {
		return nil, err
	}
	return &User{
		Name:  			src.Name,
		Key: 			src.Key,
		RequestDate: 	renew,
		RenewDate: 		requests,
		Requests: 		int(src.Requests),
		Plan: 			src.Plan,
		Stations: 		src.Stations,
	}, nil
}

func EncodeUser(src *User) *grpcusr.User {
	renew := src.RenewDate.Format(common.TimeLayout)
	requests := src.RequestDate.Format(common.TimeLayout)
	return &grpcusr.User{
		Name:  			src.Name,
		Key: 			src.Key,
		RequestDate: 	renew,
		RenewDate: 		requests,
		Requests: 		int32(src.Requests),
		Plan: 			src.Plan,
		Stations: 		src.Stations,
	}
}

func DecodeParameters(src *grpcusr.Parameters) (*Parameters, error) {
	user, err := DecodeUser(src.User)
	if err != nil {
		return nil, err
	}
	plan, err := DecodePlan(src.Plan)
	if err != nil {
		return nil, err
	}
	return &Parameters{
		User: *user,
		Plan: *plan,
	}, nil
}

func EncodeParameters(src *Parameters) *grpcusr.Parameters {
	user := EncodeUser(&src.User)
	plan := EncodePlan(&src.Plan)
	return &grpcusr.Parameters{
		User: user,
		Plan: plan,
	}
}

func DecodeSelection(src *grpcusr.Selection) (*Selection, error) {
	start, err := time.Parse(common.TimeLayout, src.Start)
	if err != nil {
		return nil, err
	}
	end, err := time.Parse(common.TimeLayout, src.End)
	if err != nil {
		return nil, err
	}
	return &Selection{
		Key: 		src.Key,
		StationID: 	src.StationID,
		Method: 	src.Method,
		Start: 		start,
		End: 		end,
	}, nil
}

func EncodeSelection(src *Selection) *grpcusr.Selection {
	start := src.Start.Format(common.TimeLayout)
	end := src.End.Format(common.TimeLayout)
	return &grpcusr.Selection{
		Key: 		src.Key,
		StationID: 	src.StationID,
		Method: 	src.Method,
		Start: 		start,
		End: 		end,
	}
}

func errorToString(err error) string{
	if err == nil {
		return "nil"
	}

	return err.Error()
}


