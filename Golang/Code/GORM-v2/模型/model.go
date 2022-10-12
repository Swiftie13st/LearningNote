package main

import (
	"database/sql"
	"time"
)

type Model struct {
	UUID uint      `gorm:"primaryKey"` // 主键
	Time time.Time `gorm:"column:my_time"`
}

type TestUser struct {
	Model        Model   `gorm:"embedded;embeddedPrefix:em_"`
	Name         string  `gorm:"default:Bruce"`
	Email        *string `gorm:"not null"`
	Age          uint8   `gorm:"comment:年龄"`
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
}

func TestUserCreate() {
	GLOBAL_DB.AutoMigrate(&TestUser{})
}
