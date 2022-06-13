package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var a byte = 'a'
	// // 输出的是ASCII码值，也就是说当我们直接输出byte（字符）的时候，输出的是这个字符对应的码值
	// fmt.Println(a)
	// // 输出的是字符
	// fmt.Printf("%c\n", a)

	// // for循环打印字符串里面的字符
	// // 通过len来循环的，相当于打印的是ASCII码
	// s := "你好 golang"
	// fmt.Println(s)
	// for i := 0; i < len(s); i++ {
	// 	fmt.Printf("%v(%c)\t", s[i], s[i])
	// }

	// // 通过rune打印的是 utf-8字符
	// for index, v := range s {
	// 	fmt.Println(index, v)
	// }

	// // 字符串转换
	// s1 := "big"
	// byteS1 := []byte(s1)
	// byteS1[0] = 'p'
	// fmt.Println(string(byteS1))

	// // rune类型
	// s2 := "你好golang"
	// byteS2 := []rune(s2)
	// byteS2[0] = '我'
	// fmt.Println(string(byteS2))

	// // 整型和浮点型之间转换
	// var aa int8 = 20
	// var bb int16 = 40
	// fmt.Println(int16(aa) + bb)

	// // 建议整型转换成浮点型
	// var cc int8 = 20
	// var dd float32 = 40
	// fmt.Println(float32(cc) + dd)

	// // 字符串类型转换
	// var i int = 20
	// var f float64 = 12.456
	// var t bool = true
	// var b byte = 'a'
	// str1 := fmt.Sprintf("%d", i)
	// fmt.Printf("类型：%v - %T \n", str1, str1)

	// str2 := fmt.Sprintf("%f", f)
	// fmt.Printf("类型：%v - %T \n", str2, str2)

	// str3 := fmt.Sprintf("%t", t)
	// fmt.Printf("类型：%v - %T \n", str3, str3)

	// str4 := fmt.Sprintf("%c", b)
	// fmt.Printf("类型：%v - %T \n", str4, str4)

	// int类型转换str类型
	var num1 int64 = 20
	s1 := strconv.FormatInt(num1, 10)
	fmt.Printf("转换：%v - %T", s1, s1)

	// float类型转换成string类型
	var num2 float64 = 3.1415926

	/*
	   参数1：要转换的值
	   参数2：格式化类型 'f'表示float，'b'表示二进制，‘e’表示 十进制
	   参数3：表示保留的小数点，-1表示不对小数点格式化
	   参数4：格式化的类型，传入64位 或者 32位
	*/
	s2 := strconv.FormatFloat(num2, 'f', -1, 64)
	fmt.Printf("\n转换：%v - %T", s2, s2)
}
