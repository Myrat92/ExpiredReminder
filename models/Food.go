package models

import "time"

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
}

func (m *Food) TableName() string {
	return TableName("food")
}

//func (m *Food) GetAllFoods() ([]*Food, int) {
//
//}