package main

import (
	"fmt"
	"gorm.io/gorm"
)

type Info struct {
	gorm.Model
	Money int
	DogID uint
}

type Dog struct {
	gorm.Model
	Name      string
	Info      Info
	Goddesses []Goddess `gorm:"many2many:dog_goddess"`
}

type Goddess struct {
	gorm.Model
	Name string
	Dogs []Dog `gorm:"many2many:dog_goddess"`
}

func Many2Many() {
	//i := Info{
	//	Money: 200,
	//}
	//g := Goddess{
	//	Model: gorm.Model{
	//		ID: 1,
	//	},
	//	Name: "G1",
	//}
	//g1 := Goddess{
	//	Model: gorm.Model{
	//		ID: 2,
	//	},
	//	Name: "G2",
	//}
	//
	d := Dog{
		Model: gorm.Model{
			ID: 1,
		},
	}
	//d2 := Dog{
	//	Model: gorm.Model{
	//		ID: 2,
	//	},
	//	Name: "Dog2",
	//}

	_ = GLOBAL_DB.AutoMigrate(&Dog{}, &Goddess{}, &Info{})

	//GLOBAL_DB.Create(&d)
	//GLOBAL_DB.Preload("Goddess").Find(&d)

	var girls []Goddess

	//GLOBAL_DB.Model(&d).Association("Goddesses").Find(&girls)
	//GLOBAL_DB.Model(&d).Preload("Dogs.Info").Association("Goddesses").Find(&girls)

	GLOBAL_DB.Model(&d).Preload("Dogs", func(db *gorm.DB) *gorm.DB {
		return db.Joins("Info").Where("money < ? ", 10000)
	}).Association("Goddesses").Find(&girls)
	fmt.Println(girls)

}
