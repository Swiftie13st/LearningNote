package main

import "fmt"

func main() {
	// var num = 10
	// if num == 10 {
	// 	fmt.Println("hello == 10")
	// } else if num > 10 {
	// 	fmt.Println("hello > 10")
	// } else {
	// 	fmt.Println("hello < 10")
	// }

	// if num2 := 10; num2 >= 10 {
	// 	fmt.Println("hello >=10")
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("%v ", i+1)
	// }

	var str = "你好golang"
	for key, value := range str {
		fmt.Printf("%v - %c \n", key, value)
	}

	var array = []string{"php", "java", "node", "golang"}
	for index, value := range array {
		fmt.Printf("%v %s \n", index, value)
	}

	var i = 0
	for {
		if i == 10 {
			fmt.Println("跳出循环")
			break
		}
		i++
		fmt.Println(i)
	}
}
