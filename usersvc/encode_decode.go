package usersvc

import (
	"context"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/usersvc/grpcusr"
	"time"
)


func EncodeCreateOrderResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(CreateOrderResponse)
	return &grpcusr.CreateOrderResponse {
		Key: res.Key,
		Err: errorToString(res.Err),
	}, nil
}

func DecodeCreateOrderRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*grpcusr.CreateOrderRequest)
	return CreateOrderRequest {
		int(req.OrderId),
		req.Email,
		req.Plan,
		req.Key,
	}, nil
}

func EncodeUpdateOrderResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(UpdateOrderResponse)
	return &grpcusr.UpdateOrderResponse {
		Key: res.Key,
		Err: errorToString(res.Err),
	}, nil
}

func DecodeUpdateOrderRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*grpcusr.UpdateOrderRequest)
	Order, err := DecodeOrder(req.Order)
	if err != nil {
		return UpdateOrderRequest{}, err
	}
	return UpdateOrderRequest{*Order}, nil
}

func EncodeDeleteOrderResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(DeleteOrderResponse)
	return &grpcusr.DeleteOrderResponse {
		Err: errorToString(res.Err),
	}, nil
}

func DecodeDeleteOrderRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*grpcusr.DeleteOrderRequest)
	return DeleteOrderRequest{int(req.OrderId)}, nil
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
	order := EncodeOrder(&res.Order)
	plan := EncodePlan(&res.Plan)
	return &grpcusr.ValidateKeyResponse {
		Order:order,
		Plan:plan,
		Err: errorToString(res.Err),
	}, nil
}

func DecodeValidateKeyRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*grpcusr.ValidateKeyRequest)
	return ValidateKeyRequest{req.Key}, nil
}

func EncodeValidateOrderResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(ValidateOrderResponse)
	order := EncodeOrder(&res.Order)
	plan := EncodePlan(&res.Plan)
	return &grpcusr.ValidateOrderResponse {
		Order: order,
		Plan: plan,
		Err: errorToString(res.Err),
	}, nil
}

func DecodeValidateOrderRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*grpcusr.ValidateOrderRequest)
	return ValidateOrderRequest{int(req.OrderId)}, nil
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
	}
}

func DecodeOrder(src *grpcusr.Order) (*Order, error) {
	requests, err := time.Parse(common.TimeLayout, src.RequestDate)
	if err != nil {
		return nil, err
	}
	return &Order{
		OrderId:  		int(src.OrderId),
		Key: 			src.Key,
		Email: 			src.Email,
		Plan: 			src.Plan,
		Stations: 		src.Stations,
		RequestDate: 	requests,
		Requests: 		int(src.Requests),
		Admin:			src.Admin,
	}, nil
}

func EncodeOrder(src *Order) *grpcusr.Order {
	requests := src.RequestDate.Format(common.TimeLayout)
	return &grpcusr.Order{
		OrderId:  		int32(src.OrderId),
		Key: 			src.Key,
		Email: 			src.Email,
		Plan: 			src.Plan,
		Stations: 		src.Stations,
		RequestDate: 	requests,
		Requests: 		int32(src.Requests),
		Admin: 			src.Admin,
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


