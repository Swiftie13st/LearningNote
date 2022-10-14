package main

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	UUID uint      `gorm:"primaryKey"` // 主键
	Time time.Time `gorm:"column:my_time"`
}

type TestUser struct {
	gorm.Model
	Name string `gorm:"default:Nobody"`
	Age  uint8  `gorm:"comment:年龄"`
}

func TestUserCreate() {
	GLOBAL_DB.AutoMigrate(&TestUser{})
}
