package main

import (
	"fmt"
	"time"
)

type WorkerManager struct {
	//用来监控Worker是否已经死亡的缓冲Channel
	workerChan chan *worker
	// 一共要监控的worker数量
	nWorkers int
}

//创建一个WorkerManager对象
func NewWorkerManager(nworkers int) *WorkerManager {
	return &WorkerManager{
		nWorkers:   nworkers,
		workerChan: make(chan *worker, nworkers),
	}
}

//启动worker池，并为每个Worker分配一个ID，让每个Worker进行工作
func (wm *WorkerManager) StartWorkerPool() {
	//开启一定数量的Worker
	for i := 0; i < wm.nWorkers; i++ {
		i := i
		wk := &worker{id: i}
		go wk.work(wm.workerChan)
	}

	//启动保活监控
	wm.KeepLiveWorkers()
}

//保活监控workers
func (wm *WorkerManager) KeepLiveWorkers() {
	//如果有worker已经死亡 workChan会得到具体死亡的worker然后 打出异常，然后重启
	for wk := range wm.workerChan {
		// log the error
		fmt.Printf("Worker %d stopped with err: [%v] \n", wk.id, wk.err)
		// reset err
		wk.err = nil
		// 当前这个wk已经死亡了，需要重新启动他的业务
		go wk.work(wm.workerChan)
	}
}

type worker struct {
	id  int
	err error
}

func (wk *worker) work(workerChan chan<- *worker) (err error) {
	// 任何Goroutine只要异常退出或者正常退出 都会调用defer 函数，所以在defer中想WorkerManager的WorkChan发送通知
	defer func() {
		//捕获异常信息，防止panic直接退出
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				wk.err = err
			} else {
				wk.err = fmt.Errorf("Panic happened with [%v]", r)
			}
		} else {
			wk.err = err
		}

		//通知 主 Goroutine，当前子Goroutine已经死亡
		workerChan <- wk
	}()

	// do something
	fmt.Println("Start Worker...ID = ", wk.id)

	// 每个worker睡眠一定时间之后，panic退出或者 Goexit()退出
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 1)
	}

	panic("worker panic..")
	//runtime.Goexit()

	return err
}

func main() {
	wm := NewWorkerManager(10)

	wm.StartWorkerPool()
}
