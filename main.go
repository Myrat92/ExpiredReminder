package main

import (
	"ExpiredReminder/models"
	_ "ExpiredReminder/routers"
	"ExpiredReminder/src/task"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func init() {
	models.Init()
	beego.BConfig.WebConfig.Session.SessionOn = true
	logs.SetLogger("file", `{"filename":"logs/ExpiredReminder.log","maxdays":30,"color":true}`)
	logs.EnableFuncCallDepth(true)
	go task.Run()
}

func main() {
	beego.Run()
}

