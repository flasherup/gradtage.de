package dailysvc

import "github.com/flasherup/gradtage.de/dailysvc/dlygrpc"

type Client interface {
	GetPeriod(id string, start string, end string) 	(resp *dlygrpc.GetPeriodResponse, 		 err error)
	PushPeriod(id string, temps []Temperature) 		(resp *dlygrpc.PushPeriodResponse, 		 err error)
	GetUpdateDate(ids []string) 					(resp *dlygrpc.GetUpdateDateResponse, 	 err error)
	UpdateAvgForYear(id string) 					(resp *dlygrpc.UpdateAvgForYearResponse, err error)
	UpdateAvgForDOY(id string, doy int) 			(resp *dlygrpc.UpdateAvgForDOYResponse,  err error)
	GetAvg(id string) 								(resp *dlygrpc.GetAvgResponse, 			 err error)
}

