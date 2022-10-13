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
	Goddess   Goddess
}

type Goddess struct {
	gorm.Model
	Name string
}

func One2One() {
	d := Dog{
		Model: gorm.Model{
			ID: 1,
		},
	}
	d2 := Dog{
		Model: gorm.Model{
			ID: 2,
		},
	}

	g := Goddess{
		Model: gorm.Model{
			ID: 1,
		},
	}
	g2 := Goddess{
		Model: gorm.Model{
			ID: 2,
		},
	}
	GLOBAL_DB.Model(&d).Association("Goddess").Append(&g)       // 添加关联
	GLOBAL_DB.Model(&d2).Association("Goddess").Append(&g2)     // 添加关联
	GLOBAL_DB.Model(&d).Association("Goddess").Delete(&g)       // 删除关联
	GLOBAL_DB.Model(&d).Association("Goddess").Replace(&g, &g2) // 替换关联
	GLOBAL_DB.Model(&d).Association("Goddess").Clear()          // 清空关联
	res := GLOBAL_DB.Model(&d2).Association("Goddess").Count()  // 关联计数
	fmt.Println(res)
}
