package stripe

import (
	"encoding/json"
)

const (
	InvoiceFinalize 				= "invoice.finalized"
	SubscriptionScheduleCanceled 	= "subscription_schedule.canceled"
)

type StatusTransitions struct {
	FinalizedAt 				string 		`json:"finalized_at"`
	MarkedUnCollectibleAt 		string 		`json:"marked_uncollectible_at"`
	PaidAt 						string 		`json:"paid_at"`
	VoidedAt 					string 		`json:"voided_at"`
}

type Recurring struct {
	AggregateUsage 			string 		`json:"aggregate_usage"`
	Interval 				string 		`json:"interval"`
	IntervalCount 			int 		`json:"interval_count"`
	UsageType 				string 		`json:"usage_type"`
}

type Price struct {
	ID 					string 		`json:"id"`
	Object 				string 		`json:"object"`
	Active 				bool 		`json:"active"`
	BillingScheme 		string 		`json:"billing_scheme"`
	Created 			int 		`json:"created"`
	Currency 			string 		`json:"currency"`
	LiveMode 			bool 		`json:"livemode"`
	LookupKey 			string 		`json:"lookup_key"`
	Metadata 			interface{} `json:"metadata"`
	Nickname 			string		`json:"nickname"`
	Product 			string		`json:"product"`
	Recurring 			Recurring	`json:"recurring"`
	TiersMode 			string		`json:"tiers_mode"`
	TransformQuantity 	string		`json:"transform_quantity"`
	Type 				string		`json:"type"`
	UnitAmount 			float64		`json:"unit_amount"`
	UnitAmountDecimal 	string		`json:"unit_amount_decimal"`
}

type PlanInvoice struct {
	ID 					string 		`json:"id"`
	Object 				string 		`json:"object"`
	Active 				bool 		`json:"active"`
	AggregateUsage 		string 		`json:"aggregate_usage"`
	Amount 				int 		`json:"amount"`
	AmountDecimal 		string 		`json:"amount_decimal"`
	BillingScheme 		string 		`json:"billing_scheme"`
	Created 			int 		`json:"created"`
	Currency 			string 		`json:"currency"`
	Interval 			string 		`json:"interval"`
	IntervalCount 		int 		`json:"interval_count"`
	LiveMode 			bool 		`json:"livemode"`
	Metadata 			interface{} `json:"metadata"`
	Nickname 			string		`json:"nickname"`
	Product 			string		`json:"product"`
	Tiers 				string		`json:"tiers"`
	TiersMode 			string		`json:"tiers_mode"`
	TransformUsage 		string		`json:"transform_usage"`
	TrialPeriodDays 	string		`json:"trial_period_days"`
	UsageType 			string		`json:"usage_type"`
}

type Period struct {
	End 		int 		`json:"end"`
	Start 		int 		`json:"start"`
}

type Data struct {
	ID               string      `json:"id"`
	Object           string      `json:"object"`
	Amount           int         `json:"amount"`
	Currency         string      `json:"currency"`
	Description      string      `json:"description"`
	Discountable     bool      	 `json:"discountable"`
	LiveMode         bool        `json:"livemode"`
	Metadata         interface{} `json:"metadata"`
	Period           Period      `json:"period"`
	Plan             PlanInvoice `json:"plan"`
	Price            Price       `json:"price"`
	Proration        bool        `json:"proration"`
	Quantity         int         `json:"quantity"`
	Subscription     string      `json:"subscription"`
	SubscriptionItem string      `json:"subscription_item"`
	TaxAmounts       []string    `json:"tax_amounts"`
	TaxRates         []string    `json:"tax_rates"`
	Type             string      `json:"type"`
}

type Lines struct {
	Data 		[]Data 		`json:"data"`
	HasMore 	bool 		`json:"has_more"`
	Object 		string 		`json:"object"`
	Url 		string 		`json:"url"`
}

type InvoiceFinalizeObject struct {
	ID 								string 				`json:"id"`
	Object 							string 				`json:"object"`
	AccountCountry 					string 				`json:"account_country"`
	AccountName 					string 				`json:"account_name"`
	AmountDue 						int 				`json:"amount_due"`
	AmountPaid 						int 				`json:"amount_paid"`
	AmountRemaining 				int 				`json:"amount_remaining"`
	ApplicationFeeAmount 			int 				`json:"application_fee_amount"`
	AttemptCount 					int 				`json:"attempt_count"`
	Attempted 						bool 				`json:"attempted"`
	AutoAdvance 					bool 				`json:"auto_advance"`
	BillingReason 					string 				`json:"billing_reason"`
	Charge 							string 				`json:"charge"`
	CollectionMethod 				string 				`json:"collection_method"`
	Created 						int 				`json:"created"`
	Currency 						string 				`json:"currency"`
	CustomFields 					string 				`json:"custom_fields"`
	Customer 						string 				`json:"customer"`
	CustomerAddress 				string 				`json:"customer_address"`
	CustomerEmail 					string 				`json:"customer_email"`
	CustomerName 					string 				`json:"customer_name"`
	CustomerPhone 					string 				`json:"customer_phone"`
	CustomerShipping 				string 				`json:"customer_shipping"`
	CustomerTaxExempt 				string 				`json:"customer_tax_exempt"`
	CustomerTaxIds 					[]string 			`json:"customer_tax_ids"`
	DefaultPaymentMethod 			string 				`json:"default_payment_method"`
	DefaultSource 					string 				`json:"default_source"`
	DefaultTaxRates 				[]string 			`json:"default_tax_rates"`
	Description 					string 				`json:"description"`
	Discount 						string 				`json:"discount"`
	DueDate 						string 				`json:"due_date"`
	EndingBalance 					string 				`json:"ending_balance"`
	Footer 							string 				`json:"footer"`
	HostedInvoiceUrl 				string 				`json:"hosted_invoice_url"`
	InvoicePdf 						string 				`json:"invoice_pdf"`
	Lines 							Lines 				`json:"lines"`
	LiveMode 						bool 				`json:"livemode"`
	Metadata 						interface{} 		`json:"metadata"`
	NextPaymentAttempt 				int 				`json:"next_payment_attempt"`
	Number 							string 				`json:"number"`
	Paid 							bool 				`json:"paid"`
	PaymentIntent 					string 				`json:"payment_intent"`
	PeriodEnd 						int 				`json:"period_end"`
	PeriodStart 					int 				`json:"period_start"`
	PostPaymentCreditNotesAmount 	int 				`json:"post_payment_credit_notes_amount"`
	PrePaymentCreditNotesAmount 	int 				`json:"pre_payment_credit_notes_amount"`
	ReceiptNumber 					string 				`json:"receipt_number"`
	StartingBalance 				int 				`json:"starting_balance"`
	StatementDescriptor 			string 				`json:"statement_descriptor"`
	Status 							string 				`json:"status"`
	StatusTransitions 				StatusTransitions 	`json:"status_transitions"`
	Subscription 					string 				`json:"subscription"`
	Subtotal 						int 				`json:"subtotal"`
	Tax 							string 				`json:"tax"`
	TaxPercent 						string 				`json:"tax_percent"`
	Total 							int 				`json:"total"`
	TotalTaxAmounts 				[]string 			`json:"total_tax_amounts"`
	TransferData 					string 				`json:"transfer_data"`
	WebHooksDeliveredAt 			string 				`json:"webhooks_delivered_at"`
}

func ParseInvoiceFinalize(data interface{}) (*InvoiceFinalizeObject, error){
	b, err := json.Marshal(data)
	if err != nil {
		return nil, nil
	}
	var res InvoiceFinalizeObject
	err = json.Unmarshal(b, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

type PlanSubscription struct {
	BillingThresholds 			string        	`json:"billing_thresholds"`
	Plan 						string        	`json:"plan"`
	Price 						string        	`json:"price"`
	Quantity 					int        		`json:"quantity"`
	TaxRates 					[]string        `json:"tax_rates"`
}

type Phase struct {
	AddInvoiceItems 			[]string        	`json:"add_invoice_items"`
	ApplicationFeePercent 		string        		`json:"application_fee_percent"`
	BillingThresholds 			string        		`json:"billing_thresholds"`
	CollectionMethod 			string         		`json:"collection_method"`
	Coupon 						string       		`json:"coupon"`
	DefaultPaymentMethod 		string         		`json:"default_payment_method"`
	DefaultTaxRates 			[]string        	`json:"default_tax_rates"`
	EndDate 					int             	`json:"end_date"`
	InvoiceSettings 			string          	`json:"invoice_settings"`
	Plans 						[]PlanSubscription 	`json:"plans"`
	Prorate 					bool				`json:"prorate"`
	ProrationBehavior 			string				`json:"proration_behavior"`
	StartDate 					int					`json:"start_date"`
	TaxPercent 					string				`json:"tax_percent"`
	TransferData 				string				`json:"transfer_data"`
	TrialEnd 					string				`json:"trial_end"`
}

type DefaultSettings struct {
	BillingThresholds 					string 				`json:"billing_thresholds"`
	CollectionMethod 					string 				`json:"collection_method"`
	DefaultPaymentMethod 				string 				`json:"default_payment_method"`
	InvoiceSettings 					string 				`json:"invoice_settings"`
	TransferData 						string 				`json:"transfer_data"`
}

type SubscriptionScheduleCanceledObject struct {
	ID 								string 				`json:"id"`
	Object 							string 				`json:"object"`
	CanceledAt 						string 				`json:"canceled_at"`
	CompletedAt 					string 				`json:"completed_at"`
	Created 						int 				`json:"created"`
	CurrentPhase 					string 				`json:"current_phase"`
	Customer 						string 				`json:"customer"`
	DefaultSettings 				DefaultSettings 	`json:"default_settings"`
	EndBehavior 					string 				`json:"end_behavior"`
	LiveMode 						bool 				`json:"livemode"`
	Metadata 						interface{} 		`json:"metadata"`
	Phases 							[]Phase	 			`json:"phases"`
	ReleasedAt 						string	 			`json:"released_at"`
	ReleasedSubscription 			string	 			`json:"released_subscription"`
	Status 							string	 			`json:"status"`
	Subscription 					string	 			`json:"subscription"`
}

func ParseSubscriptionScheduleCanceled(data interface{}) (*SubscriptionScheduleCanceledObject, error){
	b, err := json.Marshal(data)
	if err != nil {
		return nil, nil
	}
	var res SubscriptionScheduleCanceledObject
	err = json.Unmarshal(b, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}