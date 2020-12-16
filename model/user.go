package model

import (
	"memo/initDB"
	"reflect"
	"time"
)

type User struct {
	ID         int       `json:"id" gorm:"AUTO_INCREMENT"`
	Username   string    `json:"username" form:"username"`
	Password   string    `json:"password" form:"password"`
	Nickname   string    `json:"nickname" form:"nickname"`
	Phone      string    `json:"phone" form:"phone"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"create_time"`
	Icon       string    `json:"icon" form:"icon"`
	Gender     int       `json:"gender"`
	Birthday   time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
	//Birthday   string    `json:"birthday"`
	Credit int `json:"credit"`
}

func (User) TableName() string {
	return "ums_member"
}

func (user User) Insert(member User) int {
	create := initDB.DB.Create(&member)
	if create.Error != nil {
		return 0
	}
	return 1
}

func (user User) FindAll() []User {
	var users []User
	initDB.DB.Find(&users)
	return users
}

func (user User) FindById(id int) User {
	initDB.DB.Where("id = ?", id).Find(&user)
	return user
}

func (user User) FindByUsername(username string) User {
	initDB.DB.Debug().Where("username = ?", username).Find(&user)
	return user
}

func (user User) FindByPhone(phone string) User {
	initDB.DB.Where("phone = ?", phone).Find(&user)
	return user
}

func (user User) DeleteOne(id int) {
	initDB.DB.Where("id = ?", id).Delete(user)
}

func (user User) IsEmpty() bool {

	return reflect.DeepEqual(user, User{})

}
