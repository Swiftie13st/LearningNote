package main

import "fmt"

func CreatedTest() {
	dbres := GLOBAL_DB.Create(&[]TestUser{
		{Name: "", Age: 18},
		{Age: 19},
	})
	fmt.Println(dbres.Error, dbres.RowsAffected)

	if dbres.Error != nil {
		fmt.Println("创建失败")
	} else {
		fmt.Println("创建成功")
	}
}
