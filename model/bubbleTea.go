package model

import (
	"memo/initDB"
)

type BubbleTea struct {
	ID    int     `gorm:"AUTO_INCREMENT" json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Type  int     `json:"type"`
}

func (BubbleTea) TableName() string {
	return "bubble_tea"
}

func (bubble BubbleTea) Insert() int {
	create := initDB.DB.Create(&bubble)
	if create.Error != nil {
		return 0
	}
	return 1
}

func (bubble BubbleTea) FindAll() []BubbleTea {
	var bubbleTeas []BubbleTea
	initDB.DB.Find(&bubble)
	return bubbleTeas
}

func (bubble BubbleTea) FindById() BubbleTea {
	initDB.DB.First(&bubble, bubble.ID)
	return bubble
}

func (bubble BubbleTea) DeleteOne() {
	initDB.DB.Delete(bubble)
}
