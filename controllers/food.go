package controllers

import (
	"ExpiredReminder/models"
	"github.com/astaxie/beego/logs"
	"strconv"
	"time"
)

type FoodController struct {
	baseController
}

func (c *FoodController) Logging() {
	if c.Ctx.Request.Method == "POST" {
		name := c.GetString("name")
		count := c.GetString("count")
		productionDate := c.GetString("production-date")
		expiredDate := c.GetString("expired-date")
		comment := c.GetString("comment")



		countInt, err := strconv.Atoi(count)
		if err != nil {
			logs.Error("Trans food's count error: %v", err)
		}

		productionDateTime, err := time.Parse("2006-01-02", productionDate)
		expiredDateTime, err := time.Parse("2006-01-02", expiredDate)

		remainingTime := expiredDateTime.Sub(productionDateTime)

		c.o.Insert(&models.Food{Name: name, Count: countInt, ProductionDate: productionDateTime, ExpiredDate: expiredDateTime, Comment: comment, RemainingTime: remainingTime.Hours()/24})
	}

	c.TplName = c.controllerName + "/logging.html"
}

func (c *FoodController) List() {
	var foods []*models.Food
	nums, _ := c.o.QueryTable(models.Food{}).All(&foods)

	logs.Debug(foods)
	c.Data["json"] = map[string]interface{}{
		"code": 0,
		"msg": "",
		"count": nums,
		"data": foods,
	}
	c.ServeJSON()
}

func (c *FoodController) View() {
	c.TplName = c.controllerName + "/view.html"
}