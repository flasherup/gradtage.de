package common

const TimeLayout = "2006-01-02T15:04:05Z"
const TimeStartAll = "2000-01-01T00:00:00Z"
const TimeEndAll = "3008-01-01T00:00:00Z"
const TimeLayoutWBH = "2006-01-02"

const TimeLayoutDay = "2006-01-02"
const TimeLayoutMonth = "2006-01"
const TimeLayoutYear = "2006"

//Data source types
const (
	SrcTypeDWD        = "dwd"
	SrcTypeNOAA       = "noaa"
	SrcTypeCheckWX    = "cwx"
	SrcTypeMeteostat  = "mtst"
	SrcTypeWeatherBit = "wbio"
)

//Woocommerce event types
const (
	WCDeleteEvent    = "deleted"
	WCUpdateEvent    = "updated"
	WCCreatedEvent   = "created"
	WCUndefinedEvent = "undefined"
)

const (
	WCStatusTrash    = "trash"
	WCStatusActive   = "active"
	WCStatusComplete = "completed"
	WCStatusProcess  = "process"
	WCStatusOnHold   = "on-hold"
)

const (
	CDDType = "cdd"
	HDDType = "hdd"
	DDType  = "dd"
)

const (
	BreakdownDaily   = "daily"
	BreakdownMonthly = "monthly"
	BreakdownYearly  = "yearly"
)

const (
	DayCalcInt = "int"
	DayCalcMean = "mean"
	DayCalcMima = "mima"
)

//GRPC
const (
	MaxMessageSendSize    = 1024 * 1024 * 20
	MaxMessageReceiveSize = 1024 * 1024 * 20
)
