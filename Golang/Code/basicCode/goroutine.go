package main

import (
	"fmt"
	"math"
	"sync"
)

// // 定义一个协程计数器
// var wg sync.WaitGroup

// func test() {
// 	// 这是主进程执行的
// 	for i := 0; i < 1000; i++ {
// 		fmt.Println("test1 你好golang", i)
// 		//time.Sleep(time.Millisecond * 100)
// 	}
// 	// 协程计数器减1
// 	wg.Done()
// }

// func test2() {
// 	// 这是主进程执行的
// 	for i := 0; i < 1000; i++ {
// 		fmt.Println("test2 你好golang", i)
// 		//time.Sleep(time.Millisecond * 100)
// 	}
// 	// 协程计数器减1
// 	wg.Done()
// }

// func test(num int) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("协程（%v）打印的第%v条数据 \n", num, i)
// 	}
// 	// 协程计数器减1
// 	vg.Done()
// }

// var vg sync.WaitGroup

func main() {

	// // 通过go关键字，就可以直接开启一个协程
	// wg.Add(1)
	// go test()

	// // 协程计数器加1
	// wg.Add(1)
	// go test2()

	// // 这是主进程执行的
	// for i := 0; i < 1000; i++ {
	// 	fmt.Println("main 你好golang", i)
	// 	//time.Sleep(time.Millisecond * 100)
	// }
	// // 等待所有的协程执行完毕
	// wg.Wait()
	// fmt.Println("主线程退出")

	// // 获取cpu个数
	// npmCpu := runtime.NumCPU()
	// fmt.Println("cup的个数:", npmCpu)
	// // 设置允许使用的CPU数量
	// runtime.GOMAXPROCS(runtime.NumCPU() - 1)

	// for i := 0; i < 10; i++ {
	// 	go test(i)
	// 	vg.Add(1)
	// }
	// vg.Wait()
	// fmt.Println("主线程退出")
	// // 创建管道
	// ch := make(chan int, 3)

	// // 给管道里面存储数据
	// ch <- 10
	// ch <- 21
	// ch <- 32

	// // 获取管道里面的内容
	// a := <-ch
	// fmt.Println("打印出管道的值：", a)
	// fmt.Println("打印出管道的值：", <-ch)
	// fmt.Println("打印出管道的值：", <-ch)

	// // 管道的值、容量、长度
	// fmt.Printf("地址：%v 容量：%v 长度：%v \n", ch, cap(ch), len(ch))

	// // 管道的类型
	// fmt.Printf("%T \n", ch)

	// // 管道阻塞（当没有数据的时候取，会出现阻塞，同时当管道满了，继续存也会）
	// <-ch // 没有数据取，出现阻塞
	// ch <- 10
	// ch <- 10
	// ch <- 10
	// ch <- 10 // 管道满了，继续存，也出现阻塞

	// 	// 创建管道
	// 	ch := make(chan int, 10)
	// 	// 循环写入值
	// 	for i := 0; i < 10; i++ {
	// 		ch <- i
	// 	}
	// 	// for i := 0; i < 10; i++ {
	// 	// 	fmt.Println(<-ch)
	// 	// }
	// 	// 关闭管道
	// 	close(ch)

	// 	// for range循环遍历管道的值(管道没有key)
	// 	for value := range ch {
	// 		fmt.Println(value)
	// 	}
	// 	// 通过上述的操作，能够打印值，但是出出现一个deadlock的死锁错误，也就说我们需要关闭管道
	// }
	// func write(ch chan int) {
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println("写入:", i)
	// 		ch <- i
	// 		time.Sleep(time.Microsecond * 10)
	// 	}
	// 	wg.Done()

	// ch := make(chan int, 10)
	// wg.Add(1)
	// go write(ch)
	// wg.Add(1)
	// go read(ch)

	// // 等待
	// wg.Wait()
	// fmt.Println("主线程执行完毕")

	// 写入数字
	intChan := make(chan int, 1000)

	// 存放素数
	primeChan := make(chan int, 1000)

	// 存放 primeChan退出状态
	exitChan := make(chan bool, 16)

	// 开启写值的协程
	go putNum(intChan)

	// 开启计算素数的协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go primeNum(intChan, primeChan, exitChan)
	}

	// 开启打印的协程
	wg.Add(1)
	go printPrime(primeChan)

	// 匿名自运行函数
	wg.Add(1)
	go func() {
		for i := 0; i < 16; i++ {
			// 如果exitChan 没有完成16次遍历，将会等待
			<-exitChan
		}
		// 关闭primeChan
		close(primeChan)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("主线程执行完毕")
}

// func write(ch chan int) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println("写入:", i)
// 		ch <- i
// 		time.Sleep(time.Microsecond * 10)
// 	}
// 	wg.Done()
// }
// func read(ch chan int) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println("读取:", <-ch)
// 		time.Sleep(time.Microsecond * 10)
// 	}
// 	wg.Done()
// }

// 想intChan中放入 1~ 120000个数
func putNum(intChan chan int) {
	for i := 2; i < 120000; i++ {
		intChan <- i
	}
	wg.Done()
	close(intChan)
}

// cong intChan取出数据，并判断是否为素数，如果是的话，就把得到的素数放到primeChan中
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	for value := range intChan {
		var flag = true
		for i := 2; i <= int(math.Sqrt(float64(value))); i++ {
			if i%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			// 是素数
			primeChan <- value
			break
		}
	}

	// 这里需要关闭 primeChan，因为后面需要遍历输出 primeChan
	exitChan <- true

	wg.Done()
}

// 打印素数
func printPrime(primeChan chan int) {
	for value := range primeChan {
		fmt.Println(value)
	}
	wg.Done()
}

var wg sync.WaitGroup
