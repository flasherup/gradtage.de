package apisvc

import (
	"context"
	"github.com/moosh3/woogo"
	"net/http"
)

type Params struct {
	Key       string  `json:"key"`
	Station   string  `json:"station"`
	Start     string  `json:"start"`
	End       string  `json:"end"`
	TB        float64 `json:"tb"`
	TR        float64 `json:"tr"`
	Output    string  `json:"output"`
	Breakdown string  `json:"breakdown"`
	DayCalc   string  `json:"day_calc"`
}

type ParamsSourceData struct {
	Key     string `json:"key"`
	Station string `json:"station"`
	Start   string `json:"start"`
	End     string `json:"end"`
}

type ParamsSearch struct {
	Key  string `json:"key"`
	Text string `json:"text"`
}

type ParamsUser struct {
	Key    string            `json:"key"`
	Action string            `json:"action"`
	Params map[string]string `json:"params"`
}

type WCLineItems struct {
	ProductID int `json:"product_id"`
}

type WCDeleteEvent struct {
	ID int `json:"id"`
}

type WCUpdateEvent struct {
	ID              int           `json:"id"`
	ParentId        int           `json:"parent_id"`
	Status          string        `json:"status"`
	DateCreated     string        `json:"date_created"`
	DateCreatedGMT  string        `json:"date_created_gmt"`
	DateModified    string        `json:"date_modified"`
	DateModifiedGMT string        `json:"date_modified_gmt"`
	LineItems       []WCLineItems `json:"line_items"`
	Billing         woogo.Billing `json:"billing"`
}

type WoocommerceEvent struct {
	Type        string
	Signature   string
	Body        []byte
	Header      http.Header
	DeleteEvent WCDeleteEvent
	UpdateEvent WCUpdateEvent
}

type Service interface {
	GetHDD(ctx context.Context, params Params) (data [][]string, err error)
	GetHDDCSV(cts context.Context, params Params) (data [][]string, fileName string, err error)
	GetSourceData(ctx context.Context, params ParamsSourceData) (data [][]string, fileName string, err error)
	Search(ctx context.Context, params ParamsSearch) (data [][]string, err error)
	User(ctx context.Context, params ParamsUser) (data [][]string, err error)
	Woocommerce(ctx context.Context, event WoocommerceEvent) (json string, err error)
	Command(ctx context.Context, name string, params map[string]string) (json interface{}, err error)
}
