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

	// extname := ".a"
	// switch extname {
	// case ".html":
	// 	{
	// 		fmt.Println(".html")
	// 		break
	// 	}
	// case ".doc":
	// 	{
	// 		fmt.Println(".doc")
	// 		break
	// 	}
	// case ".js":
	// 	{
	// 		fmt.Println(".js")
	// 	}
	// default:
	// 	{
	// 		fmt.Println("其它后缀")
	// 	}
	// }

	// switch extname := ".a"; extname {
	// case ".html":
	// 	{
	// 		fmt.Println(".html")
	// 		break
	// 	}
	// case ".doc":
	// 	{
	// 		fmt.Println(".doc")
	// 		break
	// 	}
	// case ".js":
	// 	{
	// 		fmt.Println(".js")
	// 	}
	// default:
	// 	{
	// 		fmt.Println("其它后缀")
	// 	}
	// }

	extname := ".txt"
	switch extname {
	case ".html":
		{
			fmt.Println(".html")
			break
		}
	case ".txt", ".doc":
		{
			fmt.Println("传递来的是文档")
			break
		}
	case ".js":
		{
			fmt.Println(".js")
		}
	default:
		{
			fmt.Println("其它后缀")
		}
	}

	var n = 20
	if n > 24 {
		fmt.Println("成年人")
	} else {
		goto lable3
	}

	fmt.Println("aaa")
	fmt.Println("bbb")
lable3:
	fmt.Println("ccc")
	fmt.Println("ddd")
}
