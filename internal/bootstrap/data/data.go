package data

import "github.com/kinyokun/OpenList/cmd/flags"

func InitData() {
	initUser()
	initSettings()
	initTasks()
	if flags.Dev {
		initDevData()
		initDevDo()
	}
}
