package task

import (
	"github.com/astaxie/beego/toolbox"
	"log"
)

func notificationEveryDay() error {
	log.Println("test task")
	return nil
}

func Run() {
	//notification := toolbox.NewTask("notification", "0 0 21 * * *", notificationEveryDay)
	notification := toolbox.NewTask("notification", "0/10 * * * * *", notificationEveryDay)

	//err := notification.Run()
	//if err != nil {
	//	log.Println(err)
	//}
	exit := make(chan struct{})
	toolbox.AddTask("notification", notification)
	toolbox.StartTask()
	for {
		select {
		case <- exit:
			toolbox.StopTask()
		}
	}
}