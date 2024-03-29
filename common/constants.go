package common

const TimeLayout = "2006-01-02T15:04:05Z"
const TimeStartAll = "2000-01-01T00:00:00Z"
const TimeEndAll = "3008-01-01T00:00:00Z"
const TimeLayoutWBH = "2006-01-02"
const TimeVeryFirstWBH = "2006-01-02"
const InitialDate = "2400-01-01T00:00:00Z"

const TimeLayoutDay = "2006-01-02"
const TimeLayoutMonth = "2006-01"
const TimeLayoutYear = "2006"

const CutDateWBH = "2011-10-01"

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

//Output format
const (
	FormatCSV = "csv"
	FormatJSON = "json"
)

/*const (
	UnitCelsius = "celsius"
	UnitFahrenheit = "fahrenheit"
)*/

const (
	BreakdownDaily     = "daily"
	BreakdownWeeklyISO = "weeklyISO"
	BreakdownWeekly    = "weekly"
	BreakdownMonthly   = "monthly"
	BreakdownYearly    = "yearly"
)

const (
	Monday    = "monday"
	Tuesday   = "tuesday"
	Wednesday = "wednesday"
	Thursday  = "thursday"
	Friday    = "friday"
	Saturday  = "saturday"
	Sunday    = "sunday"
)

const (
	DayCalcInt  = "int"
	DayCalcMean = "mean"
	DayCalcMima = "mima"
)

//GRPC
const (
	MaxMessageSendSize    = 1024 * 1024 * 20
	MaxMessageReceiveSize = 1024 * 1024 * 20
)

//Services
const (
	ServiceMetrics = "metrics"
)

//Weather
const (
	EmptyWeather = -999.0
)
