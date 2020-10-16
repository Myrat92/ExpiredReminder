package controllers

import (
	"ExpiredReminder/models"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"log"
	"math"
	"time"
)

type FoodController struct {
	baseController
}

func (c *FoodController) Logging() {
	c.TplName = c.controllerName + "/logging.html"
}

func (c *FoodController) Update() {
	if c.Ctx.Request.Method == "POST" {
		name := c.GetString("name")
		count, _ := c.GetInt("count")
		productionDate := c.GetString("production-date")
		expiredDate := c.GetString("expired-date")
		comment := c.GetString("comment")

		productionDateTime, _ := time.Parse("2006-01-02", productionDate)
		expiredDateTime, _ := time.Parse("2006-01-02", expiredDate)

		num, err := c.o.QueryTable(models.Food{}).Filter("Name", name).Filter("ProductionDate", productionDateTime).
			Filter("ExpiredDate", expiredDateTime).Update(orm.Params{"Count": count+1, "Updated": time.Now()})
		if err != nil {
			logs.Error("Query same food err: %v", err)
		}

		if num == 0 {
			remainingTime := expiredDateTime.Sub(time.Now())
			c.o.Insert(&models.Food{Name: name, Count: count, ProductionDate: productionDateTime,
				ExpiredDate: expiredDateTime, Comment: comment, RemainingTime: math.Ceil(remainingTime.Hours()/24),
				Status: 0, Created: time.Now(), Updated: time.Now()})
		}

		c.Data["json"] = map[string]interface{}{
			"code": 0,
		}

		c.ServeJSON()
	}
}

func (c *FoodController) List() {
	var foods []*models.Food
	nums, _ := c.o.QueryTable(models.Food{}).Filter("Status", 0).Filter("count__gt", 0).All(&foods)

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

func (c *FoodController) Operate() {
	log.Println(c.Input())
	id := c.GetString("id")
	methodName := c.GetString("methodName")
	count, err := c.GetInt("count")
	if err != nil {
		logs.Error("count trans int error: %v", err)
	}

	if methodName == "del" {
		c.o.QueryTable(models.Food{}).Filter("ID", id).Update(orm.Params{"Status":1, "Updated":time.Now()})
	}

	if methodName == "edit" {
		c.o.QueryTable(models.Food{}).Filter("ID", id).Update(orm.Params{"Count":count, "Updated":time.Now()})
	}

	c.Data["json"] = map[string]interface{}{
		"status":1,
	}
	c.ServeJSON()
}