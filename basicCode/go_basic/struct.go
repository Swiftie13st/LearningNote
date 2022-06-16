package main

import (
	"encoding/json"
	"fmt"
)

// type myInt int

// type myFn func(int, int) int

// func fun(x int, y int) int {
// 	return x + y
// }

// func (m myInt) PrintInfo() {
// 	fmt.Println("我是自定义类型里面的自定义方法")
// }

// // 定义一个人结构体
// type Person struct {
// 	name string
// 	age  int
// 	sex  string
// }

// // 定义一个结构体方法
// func (p Person) PrintInfo() {
// 	fmt.Print(" 姓名: ", p.name)
// 	fmt.Print(" 年龄: ", p.age)
// 	fmt.Print(" 性别: ", p.sex)
// 	fmt.Println()
// }
// func (p *Person) SetInfo(name string, age int, sex string) {
// 	p.name = name
// 	p.age = age
// 	p.sex = sex
// }

// func main() {
// 	var a myInt = 10
// 	fmt.Printf("%v %T \n", a, a)
// 	a.PrintInfo()

// 	var fn myFn = fun
// 	fmt.Println(fn(1, 2))
// 	fmt.Printf("%T - %T \n", fun, fn)

// 	// 实例化结构体
// 	var person Person
// 	person.name = "张三"
// 	person.age = 20
// 	person.sex = "男"
// 	fmt.Printf("%#v\n", person)
// 	// main.Person{name:"张三", age:20, sex:"男"}

// 	// 第二种方式实例化
// 	var person2 = new(Person)
// 	person2.name = "李四"
// 	person2.age = 30
// 	person2.sex = "女"
// 	fmt.Printf("%#v\n", person2)
// 	// &main.Person{name:"李四", age:30, sex:"女"}

// 	// 第三种方式实例化
// 	var person3 = &Person{}
// 	person3.name = "赵四"
// 	person3.age = 28
// 	person3.sex = "男"
// 	fmt.Printf("%#v\n", person3)
// 	// &main.Person{name:"赵四", age:28, sex:"男"}

// 	// 第四种方式初始化
// 	var person4 = Person{
// 		name: "张三",
// 		age:  10,
// 		sex:  "女",
// 	}
// 	fmt.Printf("%#v\n", person4)
// 	// main.Person{name:"张三", age:10, sex:"女"}

// 	// 第五种方式初始化
// 	var person5 = &Person{
// 		name: "孙五",
// 		age:  10,
// 		sex:  "女",
// 	}
// 	fmt.Printf("%#v\n", person5)
// 	// &main.Person{name:"孙五", age:10, sex:"女"}

// 	// 第六种方式初始化
// 	var person6 = Person{
// 		"张三",
// 		5,
// 		"女",
// 	}
// 	fmt.Printf("%#v\n", person6)
// 	// main.Person{name:"张三", age:5, sex:"女"}

// 	person.PrintInfo()
// 	person.SetInfo("李四", 18, "男")
// 	person.PrintInfo()
// }
// type Person struct {
// 	name     string
// 	age      int
// 	hobby    []string
// 	mapValue map[string]string
// }

// func main() {
// 	// 结构体的匿名字段
// 	var person = Person{}
// 	person.name = "张三"
// 	person.age = 10

// 	// 给切片申请内存空间
// 	person.hobby = make([]string, 4, 4)
// 	person.hobby[0] = "睡觉"
// 	person.hobby[1] = "吃饭"
// 	person.hobby[2] = "打豆豆"
// 	// person.hobby[3] = "TTTT"
// 	// 给map申请存储空间
// 	person.mapValue = make(map[string]string)
// 	person.mapValue["address"] = "北京"
// 	person.mapValue["phone"] = "123456789"
// 	// person.mapValue["phone1"] = "123456789"
// 	// 加入#打印完整信息
// 	fmt.Printf("%#v", person)
// }

// // 用户结构体
// type User struct {
// 	userName string
// 	password string
// 	sex      string
// 	age      int
// 	address  Address // User结构体嵌套Address结构体
// }

// // 收货地址结构体
// type Address struct {
// 	name  string
// 	phone string
// 	city  string
// }

// func main() {
// 	var u User
// 	u.userName = "张三"
// 	u.password = "123456"
// 	u.sex = "男"
// 	u.age = 18

// 	var address Address
// 	address.name = "张三家"
// 	address.phone = "110"
// 	address.city = "北京"
// 	u.address = address
// 	fmt.Printf("%#v", u)
// }

// // 用户结构体
// type Animal struct {
// 	name string
// }

// func (a Animal) run() {
// 	fmt.Printf("%v 在运动 \n", a.name)
// }

// // 子结构体
// type Dog struct {
// 	age int
// 	// 通过结构体嵌套，完成继承
// 	Animal
// }

// func (dog Dog) wang() {
// 	fmt.Printf("%v 在汪汪汪 \n", dog.name)
// }

// func main() {
// 	var dog = Dog{
// 		age: 10,
// 		Animal: Animal{
// 			name: "阿帕奇",
// 		},
// 	}
// 	dog.run()
// 	dog.wang()
// }

// // 定义一个学生结构体，注意结构体的首字母必须大写，代表公有，否则将无法转换
// type Student struct {
// 	ID     string
// 	Gender string
// 	Name   string
// 	Sno    string
// }

// func main() {
// 	var s1 = Student{
// 		ID:     "12",
// 		Gender: "男",
// 		Name:   "李四",
// 		Sno:    "s001",
// 	}
// 	// 结构体转换成Json（返回的是byte类型的切片）
// 	jsonByte, _ := json.Marshal(s1)
// 	jsonStr := string(jsonByte)
// 	fmt.Printf(jsonStr)
// 	fmt.Println()
// 	// Json字符串转换成结构体
// 	var str = `{"ID":"12","Gender":"男","Name":"李四","Sno":"s001"}`
// 	var s2 = Student{}
// 	// 第一个是需要传入byte类型的数据，第二参数需要传入转换的地址
// 	err := json.Unmarshal([]byte(str), &s2)
// 	if err != nil {
// 		fmt.Printf("转换失败 \n")
// 	} else {
// 		fmt.Printf("%#v \n", s2)
// 	}
// }

// // 定义一个Student体，使用结构体标签
// type Student2 struct {
// 	Id     string `json:"id"` // 通过指定tag实现json序列化该字段的key
// 	Gender string `json:"gender"`
// 	Name   string `json:"name"`
// 	Sno    string `json:"sno"`
// }

// func main() {
// 	var s1 = Student2{
// 		Id:     "12",
// 		Gender: "男",
// 		Name:   "李四",
// 		Sno:    "s001",
// 	}
// 	// 结构体转换成Json
// 	jsonByte, _ := json.Marshal(s1)
// 	jsonStr := string(jsonByte)
// 	fmt.Println(jsonStr)

// 	// Json字符串转换成结构体
// 	var str = `{"Id":"12","Gender":"男","Name":"李四","Sno":"s001"}`
// 	var s2 = Student2{}
// 	// 第一个是需要传入byte类型的数据，第二参数需要传入转换的地址
// 	err := json.Unmarshal([]byte(str), &s2)
// 	if err != nil {
// 		fmt.Printf("转换失败 \n")
// 	} else {
// 		fmt.Printf("%#v \n", s2)
// 	}
// }

// 嵌套结构体 到 Json的互相转换

// 定义一个Student结构体
type Student3 struct {
	Id     int
	Gender string
	Name   string
}

// 定义一个班级结构体
type Class struct {
	Title    string
	Students []Student3
}

func main() {
	var class = Class{
		Title:    "1班",
		Students: make([]Student3, 0),
	}
	for i := 0; i < 10; i++ {
		s := Student3{
			Id:     i + 1,
			Gender: "男",
			Name:   fmt.Sprintf("stu_%v", i+1),
		}
		class.Students = append(class.Students, s)
	}
	fmt.Printf("%#v \n", class)

	// 转换成Json字符串
	strByte, err := json.Marshal(class)
	if err != nil {
		fmt.Println("打印失败")
	} else {
		fmt.Println(string(strByte))
	}
}
