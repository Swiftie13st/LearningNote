package main

import (
	"fmt"
	"sync"
	"time"
)

// // var x int64
// // var wg sync.WaitGroup
// // var lock sync.Mutex

// // func add1() {
// // 	for i := 0; i < 5000; i++ {
// // 		x = x + 1
// // 	}
// // 	wg.Done()
// // }
// // func add() {
// // 	for i := 0; i < 5000; i++ {
// // 		lock.Lock() // 加锁
// // 		x = x + 1
// // 		lock.Unlock() // 解锁
// // 	}
// // 	wg.Done()
// // }
// // func main() {
// // 	wg.Add(2)
// // 	go add()
// // 	go add()
// // 	wg.Wait()
// // 	fmt.Println(x)
// // }
// var (
// 	x      int64
// 	wg     sync.WaitGroup
// 	lock   sync.Mutex
// 	rwlock sync.RWMutex
// )

// func write() {
// 	// lock.Lock()   // 加互斥锁
// 	rwlock.Lock() // 加写锁
// 	x = x + 1
// 	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
// 	rwlock.Unlock()                   // 解写锁
// 	// lock.Unlock()                     // 解互斥锁
// 	wg.Done()
// }

// func read() {
// 	// lock.Lock()                  // 加互斥锁
// 	rwlock.RLock()               // 加读锁
// 	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
// 	rwlock.RUnlock()             // 解读锁
// 	// lock.Unlock()                // 解互斥锁
// 	wg.Done()
// }

// func main() {
// 	start := time.Now()
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go write()
// 	}

// 	for i := 0; i < 1000; i++ {
// 		wg.Add(1)
// 		go read()
// 	}

// 	wg.Wait()
// 	end := time.Now()
// 	fmt.Println(end.Sub(start))
// }
var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

func write() {
	// lock.Lock()   // 加互斥锁
	rwlock.Lock() // 加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwlock.Unlock()                   // 解写锁
	// lock.Unlock()                     // 解互斥锁
	wg.Done()
}

func read() {
	// lock.Lock()                  // 加互斥锁
	rwlock.RLock()               // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwlock.RUnlock()             // 解读锁
	// lock.Unlock()                // 解互斥锁
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
