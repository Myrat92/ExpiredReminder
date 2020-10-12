package routers

import (
	"ExpiredReminder/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	beego.AutoRouter(&controllers.UserController{})

    beego.AutoRouter(&controllers.FoodController{})
}
