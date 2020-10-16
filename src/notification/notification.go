package notification

import (
	"ExpiredReminder/src/wechat"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func SendMsg(msg string) {
	notificationMode := beego.AppConfig.String("notificationMode")

	if notificationMode == "wechat" {
		err := wechat.SendMsg(msg)
		if err != nil {
			logs.Info("Send Message failed.Err: %v", err)
		}
	}
}
