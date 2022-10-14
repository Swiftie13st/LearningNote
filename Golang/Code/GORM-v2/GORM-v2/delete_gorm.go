package main

import "fmt"

func TestDelete() {
	var users []TestUser

	//GLOBAL_DB.Unscoped().Where("name = ?", "ST_new").Delete(&users)

	GLOBAL_DB.Unscoped().Where("age = 0").Find(&users)

	fmt.Println(users)
}
