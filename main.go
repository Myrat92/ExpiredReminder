package main

import (
	"ExpiredReminder/models"
	_ "ExpiredReminder/routers"
	"github.com/astaxie/beego"
)

func init() {
	models.Init()
	beego.BConfig.WebConfig.Session.SessionOn = true
}

func main() {
	beego.Run()
}

