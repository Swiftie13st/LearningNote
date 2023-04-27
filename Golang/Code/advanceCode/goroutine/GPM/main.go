package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

//定义任务Task类型,每一个任务Task都可以抽象成一个函数
type Task struct {
	f func() error //一个task中必须包含一个具体的业务
}

//通过NewTask来创建一个Task
func NewTask(arg_f func() error) *Task {
	t := Task{
		f: arg_f,
	}
	return &t
}

//Task也需要一个执行业务的方法
func (t *Task) Execute() {
	t.f() //调用任务中已经绑定好的业务方法
}

//定义池类型
type Pool struct {
	EntryChannel chan *Task
	WorkerNum    int
	JobsChanel   chan *Task
}

//创建一个协程池
func NewPool(cap int) *Pool {
	p := Pool{
		EntryChannel: make(chan *Task),
		JobsChanel:   make(chan *Task),
		WorkerNum:    cap,
	}
	return &p
}

//协程池创建worker并开始工作
func (p *Pool) worker(workerId int) {
	//worker不断的从JobsChannel内部任务队列中拿任务
	for task := range p.JobsChanel {
		task.Execute()
		fmt.Println("workerId", workerId, "执行任务成功")
	}
}

// EntryChannel获取Task任务
func (p *Pool) ReceiveTask(t *Task) {
	p.EntryChannel <- t
}

//让协程池开始工作
func (p *Pool) Run() {
	//1:首先根据协程池的worker数量限定，开启固定数量的worker
	for i := 0; i < p.WorkerNum; i++ {
		go p.worker(i)
	}
	//2:从EntryChannel协程出入口取外部传递过来的任务
	//并将任务送进JobsChannel中
	for task := range p.EntryChannel {
		p.JobsChanel <- task
	}
	//3:执行完毕需要关闭JobsChannel和EntryChannel
	close(p.JobsChanel)
	close(p.EntryChannel)
}

func TestProcs() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(300)
	// for i := 0; i < 10; i++ {
	// 	// time.Sleep(100 * time.Millisecond)
	// 	go func() {
	// 		fmt.Println("i: ", i)
	// 		wg.Done()
	// 	}()
	// }
	for i := 0; i < 300; i++ {
		go func(i int) {
			fmt.Println("j: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func main() {
	//创建trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	//启动trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	//创建一个task
	t := NewTask(func() error {
		fmt.Println(time.Now())
		return nil
	})

	//创建一个协程池，最大开启5个协程worker
	p := NewPool(5)
	//开启一个协程，不断的向Pool输送打印一条时间的task任务
	go func() {
		for {
			p.ReceiveTask(t) //把任务推向EntryChannel
		}
	}()
	//启动协程池p
	p.Run()
	TestProcs()
}
