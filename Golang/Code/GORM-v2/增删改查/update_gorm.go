package main

func TestUpdate() {
	// Update 只更新你选择的字段
	// Updates 更新所有字段 此时有两种形式 一种为Map 一种为结构体 结构体零值不参与更新
	// Save 无论如何都更新 所有的内容 包括0值

	// update
	//GLOBAL_DB.Model(&TestUser{}).Select("name").Where("name LIKE ?", "%ST%").Update("name", "ST_new")

	// save
	//var users []TestUser
	//
	//dbRes := GLOBAL_DB.Where("name LIKE ?", "%ST%").Find(&users)
	//for i := range users {
	//	users[i].Age = 18
	//}
	//dbRes.Save(&users)

	// updates
	// struct
	var users TestUser // 多条用切片

	GLOBAL_DB.First(&users).Updates(TestUser{Name: "ST111", Age: 0}) // age 为0值不会更新
	// map
	GLOBAL_DB.First(&users).Updates(map[string]interface{}{"Name": "ST1111", "Age": 0}) // map 0值可以更新
}
