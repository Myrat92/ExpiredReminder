package task

import (
	"ExpiredReminder/models"
	"ExpiredReminder/src/notification"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
	"math"
	"strconv"
	"time"
)

const MIN = 0.000001

func notificationEveryDay() error {
	foods, _ := models.GetAllFoods()
	expiredData := make(map[string]float64)

	for _, food := range foods {
		if math.Dim(food.RemainingTime, 365) < MIN {
			expiredData[food.Name] = 365
		}
		if math.Dim(food.RemainingTime, 180) < MIN {
			expiredData[food.Name] = 180
		}
		if math.Dim(food.RemainingTime, 90) < MIN {
			expiredData[food.Name] = 90
		}
		if math.Dim(food.RemainingTime, 30) < MIN {
			expiredData[food.Name] = 30
		}
		if math.Dim(food.RemainingTime, 15) < MIN {
			expiredData[food.Name] = 15
		}
		if math.Dim(food.RemainingTime, 7) < MIN {
			expiredData[food.Name] = 7
		}
		if math.Dim(food.RemainingTime, 3) < MIN {
			expiredData[food.Name] = 3
		}
		if math.Dim(food.RemainingTime, 1) < MIN {
			expiredData[food.Name] = 1
		}
	}

	msg := "有如下物品即将过期：\r\n"
	for k,v := range expiredData {
		remainingTime := strconv.Itoa(int(v))
		msg = msg + k + ":剩余" + remainingTime + "天.\r\n"
	}
	logs.Info(msg)
	notification.SendMsg(msg)
	return nil
}

func refreshRemainingTime() error {
	foods, _ := models.GetAllFoods()
	o := orm.NewOrm()

	for _, food := range foods {
		remainingTime := math.Ceil(food.ExpiredDate.Sub(time.Now()).Hours()/24)
		o.QueryTable(models.Food{}).Filter("ID", food.Id).Update(orm.Params{"RemainingTime": remainingTime})
	}
	logs.Info("Refresh remaining time successfully.")
	return nil
}

func Run() {
	notifications := toolbox.NewTask("notification", "0 0 21 * * *", notificationEveryDay)
	refreshRemaining := toolbox.NewTask("refreshRemaining", "1 0 0 * * *", refreshRemainingTime)

	//debug
	//notifications := toolbox.NewTask("notifications", "0/10 * * * * *", notificationEveryDay)
	//refreshRemaining := toolbox.NewTask("refreshRemaining", "0/10 * * * * *", refreshRemainingTime)

	exit := make(chan struct{})
	toolbox.AddTask("notifications", notifications)
	toolbox.AddTask("refreshRemaining", refreshRemaining)
	toolbox.StartTask()
	for {
		select {
		case <- exit:
			toolbox.StopTask()
		}
	}
}