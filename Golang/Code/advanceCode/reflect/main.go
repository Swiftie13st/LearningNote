package main

import (
	"fmt"
	"reflect"
)

// func reflectSetValue1(x interface{}) {
// 	v := reflect.ValueOf(x)
// 	if v.Kind() == reflect.Int64 {
// 		v.SetInt(200) //修改的是副本，reflect包会引发panic
// 	}
// }

// func reflectSetValue2(x interface{}) {
// 	v := reflect.ValueOf(x)
// 	// 反射中使用 Elem()方法获取指针对应的值
// 	if v.Elem().Kind() == reflect.Int64 {
// 		v.Elem().SetInt(200)
// 	}
// }

// func main() {
// 	var a1 int64 = 100
// 	// reflectSetValue1(a1) //panic: reflect: reflect.Value.SetInt using unaddressable value
// 	reflectSetValue2(&a1)
// 	fmt.Println(a1)

// 	// *int类型空指针
// 	var a *int
// 	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
// 	// nil值
// 	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
// 	// 实例化一个匿名结构体
// 	b := struct{}{}
// 	// 尝试从结构体中查找"abc"字段
// 	fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())
// 	// 尝试从结构体中查找"abc"方法
// 	fmt.Println("不存在的结构体方法:", reflect.ValueOf(b).MethodByName("abc").IsValid())
// 	// map
// 	c := map[string]int{"abc": 1}
// 	// 尝试从map中查找一个不存在的键
// 	fmt.Println("map中不存在的键：", reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜扎")).IsValid())
// 	fmt.Println("map中存在的键：", reflect.ValueOf(c).MapIndex(reflect.ValueOf("abc")).IsValid())
// }
type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

// 给student添加两个方法 Study和Sleep(注意首字母大写)
func (s student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println("method num:", t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name: %s\n", t.Method(i).Name)
		fmt.Printf("method: %s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}

func main() {
	stu1 := student{
		Name:  "小王子",
		Score: 90,
	}

	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind()) // student struct
	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}

	// 通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}
	printMethod(stu1)
}
