package main

import (
	"errors"
	"fmt"
)

func main() {
	// fmt.Println(sumFn(1, 2))
	// fmt.Println(sumFn2(1, 2, 3, 4, 5, 7, 8))
	// fmt.Println(sumFn3(1, 2))
	// var _, b = sumFn3(1, 3)
	// fmt.Println(b)

	// var c calc
	// c = sumFn
	// fmt.Println(c(1, 2))

	// fmt.Println(sunFn(1, 23, subFn))

	// add := do("+")
	// fmt.Println(add(1, 5))

	// func() {
	// 	fmt.Println("匿名自执行函数")
	// }()
	// var fn = adder()
	// fmt.Println(fn())
	// fmt.Println(fn())
	// fmt.Println(fn())
	// var fn2 = adder2()
	// fmt.Println(fn2(10))
	// fmt.Println(fn2(10))
	// fmt.Println(fn2(10))

	// // defer函数
	// fmt.Println("1")
	// defer fmt.Println("2")
	// defer fmt.Println("3")
	// fmt.Println("4")

	// fmt.Println("开始")
	// defer func() {
	// 	fmt.Println("1")
	// 	fmt.Println("2")
	// }()
	// fmt.Println("结束")

	// fn11()
	// fn22()
	// fmt.Println("结束")

	myFn()
}

// 求两个数的和
func sumFn(x int, y int) int {
	return x + y
}
func subFn(x int, y int) int {
	return x - y
}

// 可变参数
func sumFn2(x ...int) int {
	sum := 0
	for _, num := range x {
		sum = sum + num
	}
	return sum
}

// 方法多返回值
func sumFn3(x int, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y
	return
}

type calc func(int, int) int

// 传递两个参数和一个方法
func sunFn(a int, b int, sum func(int, int) int) int {
	return sum(a, b)
}

// 返回一个方法
type calcType func(int, int) int

func do(o string) calcType {
	switch o {
	case "+":
		return func(i int, i2 int) int {
			return i + i2
		}
	case "-":
		return func(i int, i2 int) int {
			return i - i2
		}
	case "*":
		return func(i int, i2 int) int {
			return i * i2
		}
	case "/":
		return func(i int, i2 int) int {
			return i / i2
		}
	default:
		return nil
	}

}

// 闭包的写法：函数里面嵌套一个函数，最后返回里面的函数就形成了闭包
func adder() func() int {
	var i = 10
	return func() int {
		return i + 1
	}
}
func adder2() func(y int) int {
	var i = 10
	return func(y int) int {
		i = i + y
		return i
	}
}

func fn11() {
	fmt.Println("fn1")
}

func fn22() {
	// panic("抛出一个异常")

	// 使用recover监听异常
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	panic("抛出一个异常")
}

func readFile(fileName string) error {
	if fileName == "main.go" {
		return nil
	} else {
		return errors.New("读取文件失败")
	}
}

func myFn() {
	defer func() {
		e := recover()
		if e != nil {
			fmt.Println("给管理员发送邮件")
		}
	}()
	err := readFile("XXX.go")
	if err != nil {
		panic(err)
	}
}
