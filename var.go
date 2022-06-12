package main

import "fmt"

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
}
