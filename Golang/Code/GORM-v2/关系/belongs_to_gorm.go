package main

import (
	"gorm.io/gorm"
)

// belongsTo

type Dog1 struct {
	gorm.Model
	Name      string
	GoddessID uint
	Goddess   Goddess1
}

type Goddess1 struct {
	gorm.Model
	Name string
}

func BelongsTo() {

	g := Goddess1{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "G1",
	}

	d := Dog1{
		Model: gorm.Model{
			ID: 1,
		},
		Name:    "dog1",
		Goddess: g,
	}

	_ = GLOBAL_DB.AutoMigrate(&Dog1{})

	GLOBAL_DB.Create(&d)

}
