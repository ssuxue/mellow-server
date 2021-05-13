package initdata

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var DB *gorm.DB

func init() {
	var err error
	conStr := "root:wangqing@(127.0.0.1)/coco?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", conStr)
	if err != nil {
		log.Panicln("err:", err.Error())
	}
}
