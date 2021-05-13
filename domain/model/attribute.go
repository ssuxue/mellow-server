package model

import (
	"github.com/ssuxue/mellow-server/internal/initdata"
)

type Attribute struct {
	ID        int    `gorm:"AUTO_INCREMENT" json:"id"`
	Name      string `json:"name"`
	InputList string `json:"input_list"`
	Type      int8   `json:"type"`
}

func (attribute Attribute) TableName() string {
	return "pms_product_attribute"
}

func (attribute Attribute) Insert() int {
	create := initdata.DB.Create(&attribute)
	if create.Error != nil {
		return 0
	}
	return 1
}

func (attribute Attribute) FindAll() []Attribute {
	var attributes []Attribute
	initdata.DB.Find(&attributes)
	return attributes
}

func (attribute Attribute) FindById() Attribute {
	initdata.DB.First(&attribute, attribute.ID)
	return attribute
}

func (attribute Attribute) DeleteOne() {
	initdata.DB.Delete(attribute)
}
