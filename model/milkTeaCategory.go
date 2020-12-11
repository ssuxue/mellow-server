package model

import "memo/initDB"

type MilkTeaCategory struct {
	ID int `gorm:"AUTO_INCREMENT" json:"id"`
	Name string `json:"name"`
}

func (MilkTeaCategory) TableName() string {
	return "milk_type"
}

func (milkTeaCategory MilkTeaCategory) FindCategory() []MilkTeaCategory {
	var categories []MilkTeaCategory
	initDB.DB.Find(&categories)
	return categories
}
