package model

import "memo/initDB"

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
	create := initDB.DB.Create(&attribute)
	if create.Error != nil {
		return 0
	}
	return 1
}

func (attribute Attribute) FindAll() []Attribute {
	var attributes []Attribute
	initDB.DB.Find(&attributes)
	return attributes
}

func (attribute Attribute) FindById() Attribute {
	initDB.DB.First(&attribute, attribute.ID)
	return attribute
}

func (attribute Attribute) DeleteOne() {
	initDB.DB.Delete(attribute)
}
