package main

import "fmt"

func main() {
	// 类型断言
	var a interface{}
	a = "132"
	value, isString := a.(string)
	if isString {
		fmt.Println("是String类型, 值为：", value)
	} else {
		fmt.Println("断言失败")
	}
	Print(2147483647)
	Print2(true)

	var dog = &Dog{
		"小黑",
	}
	// // 同时实现两个接口
	// var d1 Animal = dog
	// var d2 Animal2 = dog
	// d1.SetName("小鸡")
	// fmt.Println(d2.GetName())

	// 同时实现两个接口
	var d Animal = dog
	d.SetName("小鸡")
	fmt.Println(d.GetName())

	// golang中空接口和类型断言
	var userInfo = make(map[string]interface{})
	userInfo["userName"] = "zhangsan"
	userInfo["age"] = 10
	userInfo["hobby"] = []string{"吃饭", "睡觉"}
	fmt.Println(userInfo["userName"])
	fmt.Println(userInfo["age"])
	fmt.Println(userInfo["hobby"])
	// 但是我们空接口如何获取数组中的值？发现 userInfo["hobby"][0]  这样做不行
	// fmt.Println(userInfo["hobby"][0])

	// 这个时候我们就可以使用类型断言了
	hobbyValue, ok := userInfo["hobby"].([]string)
	if ok {
		fmt.Println(hobbyValue[0])
	}
}

// 定义一个方法，可以传入任意数据类型，然后根据不同类型实现不同的功能
func Print(x interface{}) {
	if _, ok := x.(string); ok {
		fmt.Println("传入参数是string类型")
	} else if _, ok := x.(int); ok {
		fmt.Println("传入参数是int类型")
	} else {
		fmt.Println("传入其它类型")
	}
}

func Print2(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int类型")
	case string:
		fmt.Println("string类型")
	case bool:
		fmt.Println("bool类型")
	default:
		fmt.Println("其它类型")
	}
}

// 定义一个Animal的接口，Animal中定义了两个方法，分别是setName 和 getName，分别让DOg结构体和Cat结构体实现
type Animal1 interface {
	SetName(string)
}

// 接口2
type Animal2 interface {
	GetName() string
}

type Animal interface {
	Animal1
	Animal2
}

type Dog struct {
	Name string
}

func (d *Dog) SetName(name string) {
	d.Name = name
}
func (d Dog) GetName() string {
	return d.Name
}
