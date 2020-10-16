package wechat

import (
	"ExpiredReminder/util"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func GetAccessToken() interface{} {
	corpid := beego.AppConfig.String("corpid")
	expiredReminderSecret := beego.AppConfig.String("expiredReminderSecret")
	url := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=" + corpid + "&corpsecret=" + expiredReminderSecret

	data, err := util.HttpGetJson(url)
	if err != nil {
		logs.Error(err)
		return nil
	}

	errCode := data["errcode"].(float64)
	if errCode != 0 {
		logs.Error("Get Access Token failed. ErrCode:%d.", errCode)
		return nil
	}
	return data["access_token"]
}

func SendMsg(msg string) error {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", GetAccessToken())
	agentId := beego.AppConfig.String("agentId")
	req := map[string]interface{}{
		"touser":  "@all",
		"msgtype": "text",
		"agentid": agentId,
		"text": map[string]interface{}{
			"content": msg,
		},
	}

	_, err := util.HttpPostJson(url, req)
	if err != nil {
		return err
	}
	return nil
}