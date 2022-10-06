package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

//计算int64随机数各位数和
func sumInt64(x int64) int64 {
	var sum int64 = 0
	for x > 0 {
		sum += x % 10
		x /= 10
	}
	return sum
}

func getRand64(jobChan chan<- int64) {
	defer wg.Done()
	for i := 0; i < 24; i++ {
		var randNum int64 = rand.Int63()
		// fmt.Println(randNum)
		jobChan <- randNum
	}
	close(jobChan)
}

func getResult(jobChan <-chan int64, resultChan chan<- int64) {
	defer wg.Done()
	i := <-jobChan
	resultChan <- sumInt64(i)
}

func main() {
	// 使用无缓冲通道，则不能使用sync.Wait()等待。因为如果进行等待，则有一个通道既无法读，也无法写
	jobChan := make(chan int64, 111)
	resultChan := make(chan int64, 111)
	wg.Add(1)
	go getRand64(jobChan)

	for i := 0; i < 24; i++ {
		wg.Add(1)
		go getResult(jobChan, resultChan)

	}
	wg.Wait()
	close(resultChan)
	for i := range resultChan {
		fmt.Println(i)
	}
}
