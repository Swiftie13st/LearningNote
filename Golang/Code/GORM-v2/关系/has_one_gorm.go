package main

import (
	"gorm.io/gorm"
)

// has one

type Dog2 struct {
	gorm.Model
	Name       string
	Goddess2ID uint // 狗链子
}

type Goddess2 struct {
	gorm.Model
	Name string
	Dog  Dog2
}

func HasOne() {
	d := Dog2{
		Name: "dog2",
	}

	g := Goddess2{
		Name: "G2",
		Dog:  d,
	}

	_ = GLOBAL_DB.AutoMigrate(&Dog2{}, &Goddess2{})

	GLOBAL_DB.Create(&g)

}
