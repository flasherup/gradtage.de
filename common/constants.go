package common

const TimeLayout 		= "2006-01-02T15:04:05Z"
const TimeStartAll 		= "2008-01-01T00:00:00Z"
const TimeEndAll	 	= "3008-01-01T00:00:00Z"
const TimeLayoutWBH      = "2006-01-02"

//Data source types
const (
	SrcTypeDWD 			= "dwd"
	SrcTypeNOAA 		= "noaa"
	SrcTypeCheckWX 		= "cwx"
	SrcTypeMeteostat 	= "mtst"
	SrcTypeWeatherBit 	= "wbio"
)

//Woocommerce event types
const (
	WCDeleteEvent = "deleted"
	WCUpdateEvent = "updated"
	WCCreatedEvent = "created"
	WCUndefinedEvent = "undefined"
)