package main

import (
	"ExpiredReminder/models"
	_ "ExpiredReminder/routers"
	"ExpiredReminder/src/task"
	"github.com/astaxie/beego"
)

func init() {
	models.Init()
	beego.BConfig.WebConfig.Session.SessionOn = true
}

func main() {
	go task.Run()
	beego.Run()
}

