package data

import (
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/usersvc"
	"time"
)

//var startTime, _ = time.Parse("2006", "2009")
var endTime, _ = time.Parse(common.TimeLayoutWBH, "2020-12-31")

var startTimeAll, _ = time.Parse(common.TimeLayout, common.TimeStartAll)
var endTimeAll, _ = time.Parse(common.TimeLayout, common.TimeEndAll)

var Plans = []usersvc.Plan{
	{usersvc.PlanCanceled, 		0, 		0, 	false,false,false, 	startTimeAll, 	startTimeAll, 1},
	{usersvc.PlanTrial, 			1, 		100, 	true, true, true, 	startTimeAll, 	endTimeAll, 	14},
	{usersvc.PlanLite, 			10, 		100, 	true, true, false, 	startTimeAll, 	endTimeAll,29},
	{usersvc.PlanProfessional, 	25, 		1000, 	true, true, true, 	startTimeAll, 	endTimeAll,29},
	{usersvc.PlanEnterprise, 		100000, 	4000, 	true, true, true, 	startTimeAll, 	endTimeAll,29},
	{usersvc.PlanAdmin, 			-1, 		-1, 	true, true, true, 	startTimeAll, 	endTimeAll,-1},
}

var UserKeys = map[string]string{
	"admin":   "sZ4XsY2NzAVhqn1SYOE0",
	"starter": "mRtXC3TZVv5OHO3KwdeR",
	"trial":   "p09w8hugMlg5HLcUu3bL",
}
