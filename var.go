package main

import "fmt"

func getUserInfo() (string, int) {
	return "zhangsan", 10
}

func main() {
	fmt.Println("hello")
	fmt.Print("A", "B", "C")
	fmt.Println()
	var a = 10
	fmt.Printf("%d", a)

	var name = "zhangsan1"
	var name2 string = "zhangsan2"
	name3 := "zhangsan3"

	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(name3)
	fmt.Printf("name1=%v name2=%v name3=%v \n", name, name2, name3)

	// var (
	// 	username string
	// 	age      int
	// 	sex      string
	// )
	// username = "zhangshan"
	// age = 10
	// sex = "男"

	// var (
	// 	username = "zhangsan"
	// 	age      = 10
	// 	sex      = "男"
	// )
	username, age, sex := "zhangsan", 10, "男"

	fmt.Println(username, age, sex)

	var username1, age1 = getUserInfo()
	fmt.Println(username1, age1)
	var username2, _ = getUserInfo()
	fmt.Println(username2)

}
