package impl

import (
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/usersvc"
)

const (
	CrateAction 	= "create"
	AutoCrateAction = "auto_create"
	SetPLanAction 	= "set_plan"
	RenewAction 	= "renew"
)

var productMapToPlan  = map[string]string{
	"1010": usersvc.PlanLite,
	"25": usersvc.PlanLite,
	"1286": usersvc.PlanLite,
	"1299": usersvc.PlanProfessional,
	"32": usersvc.PlanProfessional,
	"1305": usersvc.PlanEnterprise,
	"394": usersvc.PlanEnterprise,
	"1554": usersvc.PlanTrial,
	"1567": usersvc.PlanTrial,
}

func CreateWoocommerceOrder(client usersvc.Client, orderId int, email, key, planId string) error {
	plan, ok := productMapToPlan[planId]
	if !ok {
		plan = usersvc.PlanTrial
	}

	_, err := client.CreateOrder(orderId, email, plan, key)
	if err != nil {
		return err
	}
	return nil
}

func UpdateWoocommerceOrder(client usersvc.Client, status, email, planId string, order usersvc.Order) error {
	plan, ok := productMapToPlan[planId]
	if !ok {
		plan = usersvc.PlanTrial
	}

	if status == common.WCStatusTrash {
		client.DeleteOrder(order.OrderId)
	} else if status == common.WCStatusProcess || status == common.WCStatusActive || status == common.WCStatusComplete {
		order.Plan = plan
		_, err := client.UpdateOrder(order)
		if err != nil {
			return err
		}
	} else {
		order.Plan = usersvc.PlanCanceled
		_, err := client.UpdateOrder(order)
		if err != nil {
			return err
		}
	}
	return nil
}
