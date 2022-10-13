package main

import (
	"fmt"
	"gorm.io/gorm"
)

// has one

type Dog struct {
	gorm.Model
	Name      string
	GoddessID uint // 狗链子
}

type Goddess struct {
	gorm.Model
	Name string
	Dog  Dog
}

func Preload() {

	var girl Goddess

	GLOBAL_DB.Preload("Dog").First(&girl, 2)
	fmt.Println(girl)

}
