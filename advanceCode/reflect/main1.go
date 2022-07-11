package main

import (
	"fmt"
	"reflect"
)

type myInt int64

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	// fmt.Printf("type:%v\n", v)
	fmt.Printf("type(.Name()): %v kind(.Kind()): %v\n", v.Name(), v.Kind())
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

func main1() {
	// var a float32 = 3.14
	// reflectType(a) // type:float32
	// var b int64 = 100
	// reflectType(b) // type:int64

	// var a *float32 // 指针
	// var b myInt    // 自定义类型
	// var c rune     // 类型别名
	// reflectType(a) // type: kind:ptr
	// reflectType(b) // type:myInt kind:int64
	// reflectType(c) // type:int32 kind:int32

	// type person struct {
	// 	name string
	// 	age  int
	// }
	// type book struct{ title string }
	// var d = person{
	// 	name: "沙河小王子",
	// 	age:  18,
	// }
	// var e = book{title: "《跟小王子学Go语言》"}
	// reflectType(d) // type:person kind:struct
	// reflectType(e) // type:book kind:struct

	var a float32 = 3.14
	var b myInt = 100
	reflectValue(a) // type is float32, value is 3.140000
	reflectValue(b) // type is int64, value is 100
	// 将int类型的原始值转换为reflect.Value类型
	c := reflect.ValueOf(10)
	fmt.Printf("type c :%T\n", c) // type c :reflect.Value
}
