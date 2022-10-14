package main

import "fmt"

type UserInfo struct {
	Name string
	Age  uint8
}

func TestSelect() {
	// 接收返回值
	// 用map接收
	//var result map[string]interface{} // map需用make初始化才可使用，but新版本可以
	//var result = make(map[string]interface{})
	//result := map[string]interface{}{}
	//GLOBAL_DB.Model(&TestUser{}).First(&result)
	//GLOBAL_DB.Model(&TestUser{}).Find(&result)
	//fmt.Println(result)
	//for i, j := range result {
	//	fmt.Println(i, j)
	//}
	// 用结构体接收
	//User := TestUser{}
	//GLOBAL_DB.Model(&TestUser{}).Last(&User)
	//fmt.Println(User)

	// 条件
	//User := TestUser{}
	////GLOBAL_DB.Where("name = ?", "Nobody").Last(&User)
	////GLOBAL_DB.Where("name = ? AND age = 21", "Nobody").Last(&User)
	////GLOBAL_DB.Where(TestUser{Name: "Nobody"}).First(&User)
	//GLOBAL_DB.Where(map[string]interface{}{
	//	"name": "Nobody",
	//}).First(&User)
	//fmt.Println(User)

	//User := TestUser{}
	//GLOBAL_DB.Model(&TestUser{}).First(&User, 3) // 主键
	//GLOBAL_DB.Model(&TestUser{}).First(&User, "name = ?", "Nobody") // String
	//GLOBAL_DB.Model(&TestUser{}).First(&User, map[string]interface{}{
	//	"name": "ST",
	//}) // map
	//GLOBAL_DB.Model(&TestUser{}).First(&User, TestUser{Name: "ST"}) // struct
	//fmt.Println(User)

	var User []TestUser
	GLOBAL_DB.Select("name").Where("name LIKE ?", "%ST%").Find(&User)
	fmt.Println(User)
	fmt.Println("*****************************************")
	var u []UserInfo
	GLOBAL_DB.Model(&TestUser{}).Select("name").Where("name LIKE ?", "%ST%").Find(&u)
	fmt.Println(u)
}
