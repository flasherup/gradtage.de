package utils

import "github.com/flasherup/gradtage.de/usersvc"

func CloneParameters(src *usersvc.Parameters) *usersvc.Parameters {
	res := usersvc.Parameters {
		User: *CloneUser(&src.User),
		Plan: *ClonePlan(&src.Plan),
	}
	return &res
}

func CloneUser(src *usersvc.User) *usersvc.User {
	res := usersvc.User {
		Name: 			src.Name,
		Key: 			src.Key,
		RenewDate: 		src.RenewDate,
		RequestDate: 	src.RequestDate,
		Requests: 		src.Requests,
		Plan: 			src.Plan,
		Stations: 		src.Stations,
	}
	return &res
}


func ClonePlan(src *usersvc.Plan) *usersvc.Plan {
	res := usersvc.Plan {
		Name:			src.Name,
		Stations:		src.Stations,
		Limitation:		src.Limitation,
		HDD:			src.HDD,
		DD:				src.DD,
		CDD:			src.CDD,
		Start:			src.Start,
		End:			src.End,
		Period:			src.Period,
		Admin:			src.Admin,
	}

	return &res
}