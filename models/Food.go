package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Food struct {
	Id int ``
	Name string
	Count int
	ProductionDate time.Time
	ExpiredDate time.Time
	RemainingTime float64
	Comment string
	Created time.Time
	Updated time.Time
	Status int
}

func (m *Food) TableName() string {
	return TableName("food")
}

func GetAllFoods() ([]*Food, int64) {
	o := orm.NewOrm()

	var foods []*Food
	nums, _ := o.QueryTable(Food{}).All(&foods)
	return foods, nums
}