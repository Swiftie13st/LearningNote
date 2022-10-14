package main

import (
	"fmt"
	"gorm.io/gorm"
)

//type Info struct {
//	gorm.Model
//	Money int
//	DogID uint
//}
//
//type Dog struct {
//	gorm.Model
//	Name      string
//	GoddessID uint // 狗链子
//	Info      Info
//}
//
//type Goddess struct {
//	gorm.Model
//	Name string
//	Dogs []Dog
//}

func One2Many() {
	//d := Dog{
	//	Model: gorm.Model{
	//		ID: 1,
	//	},
	//	Name: "Dog1",
	//}
	//d2 := Dog{
	//	Model: gorm.Model{
	//		ID: 2,
	//	},
	//	Name: "Dog2",
	//}
	//
	//g := Goddess{
	//	Model: gorm.Model{
	//		ID: 1,
	//	},
	//	Name: "G1",
	//	Dogs: []Dog{d, d2},
	//}
	//
	//_ = GLOBAL_DB.AutoMigrate(&Dog{}, &Goddess{}, &Info{})
	//
	//GLOBAL_DB.Create(&g)

	var girl Goddess

	//GLOBAL_DB.Preload("Dogs.Info", "money > 100").Preload("Dogs", "name = ? ", "Dog1").First(&girl)

	GLOBAL_DB.Preload("Dogs", func(db *gorm.DB) *gorm.DB {
		return db.Joins("Info").Where("money > 100")
	}).First(&girl)
	fmt.Println(girl)

}
