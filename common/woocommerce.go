package common

type WCMetaData struct {
	ID    int         `json:"id"`
	Key   *string     `json:"key"`
	Value interface{} `json:"value"`
}