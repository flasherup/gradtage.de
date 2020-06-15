package apisvc

import "context"

type Params struct {
	Key     string  `json:"key"`
	Station string  `json:"station"`
	Start   string  `json:"start"`
	End     string  `json:"end"`
	TB      float64 `json:"tb"`
	TR      float64 `json:"tr"`
	Output  string  `json:"output"`
}

type ParamsSourceData struct {
	Key 	string  `json:"key"`
	Station string  `json:"station"`
	Start   string  `json:"start"`
	End     string  `json:"end"`
	Type	string	`json:"type"`
}

type ParamsSearch struct {
	Key 	string  `json:"key"`
	Text 	string  `json:"text"`
}

type ParamsUser struct {
	Key 	string  			`json:"key"`
	Action 	string  			`json:"action"`
	Params 	map[string]string 	`json:"params"`
}

type ParamsPlan struct {
	Key 	string  			`json:"key"`
	Action 	string  			`json:"action"`
	Params 	map[string]string `json:"params"`
}

type StripeData struct {
	Object 		interface{}  			`json:"object"`
}

type StripeEvent struct {
	Created 			int 		`json:"created"`
	LiveMode 			bool 		`json:"livemode"`
	ID 					string 		`json:"id"`
	Type 				string 		`json:"type"`
	Object 				string 		`json:"object"`
	Request 			string 		`json:"request"`
	PendingWebHooks 	int 		`json:"pending_webhooks"`
	ApiVersion 			int 		`json:"api_version"`
	Data 				StripeData 	`json:"data"`
}

type Service interface {
	GetHDD(ctx context.Context, params Params) (data [][]string, err error)
	GetHDDCSV(cts context.Context, params Params) (data [][]string, fileName string, err error)
	GetSourceData(ctx context.Context, params ParamsSourceData) (data [][]string, fileName string, err error)
	Search(ctx context.Context, params ParamsSearch) (data [][]string, err error)
	User(ctx context.Context, params ParamsUser) (data [][]string, err error)
	Plan(ctx context.Context, params ParamsPlan) (data [][]string, err error)
	Stripe(ctx context.Context, event StripeEvent) (json string, err error)
}