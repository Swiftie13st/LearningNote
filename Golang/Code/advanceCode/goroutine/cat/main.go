package main

import (
	"fmt"
	"sync"
)

var (
	countNum int = 100
	wg       sync.WaitGroup
)

//3个函数分别打印cat、dog、fish，要求每个函数都要起一个goroutine，按照cat、dog、fish顺序打印在屏幕上10次
func main() {
	dogCh := make(chan struct{}, 1)
	defer close(dogCh)
	catCh := make(chan struct{}, 1)
	defer close(catCh)
	fishCh := make(chan struct{}, 1)
	defer close(fishCh)

	wg.Add(3)
	go catPrint(&catCh, &dogCh)
	go dogPrint(&dogCh, &fishCh)
	go fishPrint(&fishCh, &catCh)

	catCh <- struct{}{}
	wg.Wait()
}
func catPrint(catCh *chan struct{}, dogCh *chan struct{}) {
	count := 0
	for {
		if count >= countNum {
			wg.Done()
			//fmt.Println("cat quit")
			return
		}
		<-*catCh
		fmt.Println("cat", count+1)
		count++
		*dogCh <- struct{}{}
	}
}

func dogPrint(dogCh *chan struct{}, fishCh *chan struct{}) {
	count := 0
	for {
		if count >= countNum {
			wg.Done()
			//fmt.Println("dog quit")
			return
		}
		<-*dogCh
		fmt.Println("dog")
		count++
		*fishCh <- struct{}{}
	}
}

func fishPrint(dogCh *chan struct{}, catCh *chan struct{}) {
	count := 0
	for {
		if count >= countNum {
			wg.Done()
			//fmt.Println("fish quit")
			return
		}
		<-*dogCh
		fmt.Println("fish")
		count++
		*catCh <- struct{}{}
	}
}
