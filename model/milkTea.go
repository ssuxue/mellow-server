package model

import (
	"memo/initDB"
)

type MilkTea struct {
	ID int `gorm:"AUTO_INCREMENT" json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Type int `json:"type"`
}

func (MilkTea) TableName() string {
	return "milk_tea"
}

func (milkTea MilkTea) Insert() int {
	create := initDB.DB.Create(&milkTea)
	if create.Error != nil {
		return 0
	}
	return 1
}

func (milkTea MilkTea) FindAll() []MilkTea {
	var articles []MilkTea
	initDB.DB.Find(&articles)
	return articles
}

func (milkTea MilkTea) FindById() MilkTea {
	initDB.DB.First(&milkTea, milkTea.ID)
	return milkTea
}

func (milkTea MilkTea) DeleteOne() {
	initDB.DB.Delete(milkTea)
}
