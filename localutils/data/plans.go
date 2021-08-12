package data

import (
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/usersvc"
	"time"
)

var startTime, _ = time.Parse("2006" , "2009")
var endTime, _ = time.Parse("2006" , "2019")


var startTimeAll, _ = time.Parse(common.TimeLayout , common.TimeStartAll)
var endTimeAll, _ = time.Parse(common.TimeLayout , common.TimeEndAll)

var Plans = []usersvc.Plan{
	usersvc.Plan{ "trial", 			1, 		100, 	true, true, true, 	startTime, 	endTime, 		14, false },
	usersvc.Plan{ "starter", 		1, 		100, 	true, true, false, 	startTimeAll, endTimeAll, 	29, false },
	usersvc.Plan{ "basic", 			3, 		100, 	true, true, false, 	startTimeAll, endTimeAll, 	29, false },
	usersvc.Plan{ "advanced", 		10, 		100, 	true, true, false, 	startTimeAll, endTimeAll, 	29, false },
	usersvc.Plan{ "professional", 	20, 		1000, true, true, true, 	startTimeAll, endTimeAll, 	29, false },
	usersvc.Plan{ "enterprise", 		100000, 	4000, true, true, true, 	startTimeAll, endTimeAll, 	29, false },
	usersvc.Plan{ "admin",	 		-1, 		-1, 	true, true, true, 	startTimeAll, endTimeAll, 	-1, true },
}

var UserKeys = map[string]string {
	"admin": "sZ4XsY2NzAVhqn1SYOE0",
	"starter": "mRtXC3TZVv5OHO3KwdeR",
	"trial": "p09w8hugMlg5HLcUu3bL",
}
