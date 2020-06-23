package stripe

import (
	"encoding/json"
)

const (
	InvoiceFinalize 				= "invoice.finalized"
	SubscriptionScheduleCanceled 	= "subscription_schedule.canceled"
)

type InvoiceFinalizeObject struct {
	UserID 		string
	UserEmail 	string
}


func ParseInvoiceFinalize(data interface{}) (*InvoiceFinalizeObject, error){
	b, err := json.Marshal(data)
	if err != nil {
		return nil, nil
	}
	var src map[string]interface{}
	err = json.Unmarshal(b, &src)
	if err != nil {
		return nil, err
	}

	res := InvoiceFinalizeObject{
		UserID: 	src["customer"].(string),
		UserEmail: 	src["customer_email"].(string),
	}

	return &res, nil
}

type SubscriptionScheduleCanceledObject struct {
	UserID string
}

func ParseSubscriptionScheduleCanceled(data interface{}) (*SubscriptionScheduleCanceledObject, error){
	b, err := json.Marshal(data)
	if err != nil {
		return nil, nil
	}

	var src map[string]interface{}
	err = json.Unmarshal(b, &src)
	if err != nil {
		return nil, err
	}

	res := SubscriptionScheduleCanceledObject{
		UserID: 	src["customer"].(string),
	}
	return &res, nil
}