package controllers

import (
	"ExpiredReminder/models"
	"github.com/Echosong/beego_blog/util"
	"strings"
)

type UserController struct {
	baseController
}

func (c *UserController) Login() {
	if c.Ctx.Request.Method == "POST" {
		username := c.GetString("username")
		password := c.GetString("password")
		user := models.User{Username: username}

		c.o.Read(&user, "username")
		if user.Password == "" {
			c.History("帐号不存在", "")
		}

		if util.Md5(password) != strings.Trim(user.Password, " ") {
			c.History("密码错误", "")
		}

		user.LastIp = c.getClientIp()
		user.LoginCount = user.LoginCount +1
		if _, err := c.o.Update(&user); err != nil {
			c.History("登录异常", "")
		} else {
			c.History("登录成功", "/food/logging.html")
		}
		c.SetSession("user", user)
	}
	c.TplName = c.controllerName + "/login.html"
}