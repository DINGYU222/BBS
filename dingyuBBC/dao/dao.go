package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func init() {
	db, _ = gorm.Open("mysql", "root:123456@/bbs?charset=utf8mb4&parseTime=True&loc=Local")
}

func GetDB() *gorm.DB {
	return db
}
