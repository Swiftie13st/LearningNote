package main

import "fmt"

func main() {
	// 指针取值
	var c = 20
	// 得到c的地址，赋值给d
	var d = &c
	// 打印d的值，也就是c的地址
	fmt.Println(d)
	// 取出d指针所对应的值
	fmt.Println(*d)
	// c对应地址的值，改成30
	*d = 30
	// c已经变成30了
	fmt.Println(c)

	x := 5
	fn4(x)
	fmt.Println(x)
	fn5(&x)
	fmt.Println(x)

	// var userInfo = make(map[string]string)
	// userInfo["userName"] = "zhangsan"
	// fmt.Println(userInfo)

	// var array = make([]int, 4, 4)
	// array[0] = 1
	// fmt.Println(array)
	// // 指针变量初始化
	// var a *int
	// *a = 100
	// fmt.Println(a)

	// 使用new关键字创建指针
	aPoint := new(int)
	bPoint := new(bool)
	fmt.Printf("%T \n", aPoint)
	fmt.Printf("%T \n", bPoint)
	fmt.Println(*aPoint)
	fmt.Println(*bPoint)
}

// 这个类似于值传递
func fn4(x int) {
	x = 10
}

// 这个类似于引用数据类型
func fn5(x *int) {
	*x = 20
}
