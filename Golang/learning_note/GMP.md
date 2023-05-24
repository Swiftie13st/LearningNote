
## Goroutine

### 介绍一下goroutine

Goroutine，经 Golang 优化后的特殊协程，协程是一种更细粒度的调度，可以满足多个不同处理逻辑的协程共享一个线程资源，它的创建销毁和调度的成本都非常的小。它与线程存在映射关系，为 M：N，可以利用多个线程实现并行，并且go实现了GMP调度模型，可以通过调度器的调度来实现线程间的动态绑定和灵活调度。
GMP：M 是操作系统线程的抽象，P 则是用于管理 G的调度器。P 结构负责管理 G的调度，包括创建、销毁、挂起和唤醒等，而 M 结构则负责将 G 绑定到操作系统线程上并执行它们。调度流程：M如果想要运行G就需要与P绑定，调度时先找p.runnext然后p的本地队列，再找全局队列，最后找准备就绪的网络协程，这样的好处是取本地队列时可以接近于无锁化减少锁的竞争。其中每61次调度会直接从全局队列中取G执行，并且把一个G放入本地队列，避免全局队列的G被饿死。如果都没有的话就会触发work stealing机制，会尝试从其他P的本地队列中偷取一半的G来执行。当M执行G时遇到了系统调用或其他阻塞行为，M会阻塞，此时runtime会通过handoff机制将M与P解绑，P与空闲的M或新建一个M绑定执行后续的G。当M结束阻塞后，G会尝试获取一个空闲的P，如果获取不到则这个M会变成休眠状态

优先取P本地队列，其次取全局队列，最后取wait队列

（1）M 是线程的抽象；G 是 goroutine；P 是承上启下的调度器；
（2）M调度G前，需要和P绑定；
（3）全局有多个M和多个P，但同时并行的G的最大数量等于P的数量；
（4）G的存放队列有三类：P的本地队列；全局队列；和wait队列（图中未展示，为io阻塞就绪态goroutine队列）；
（5）M调度G时，优先取P本地队列，其次取全局队列，最后取wait队列；这样的好处是，取本地队列时，可以接近于无锁化，减少全局锁竞争；
（6）为防止不同P的闲忙差异过大，设立work-stealing机制，本地队列为空的P可以尝试从其他P本地队列偷取一半的G补充到自身队列.

（1）与线程存在映射关系，为 M：N；
（2）创建、销毁、调度在用户态完成，对内核透明，足够轻便；
（3）可利用多个线程，实现并行；
（4）通过调度器的斡旋，实现和线程间的动态绑定和灵活调度；
（5）栈空间大小可动态扩缩，因地制宜.

调度假设当前正在执行G1，G1阻塞（如系统调用），此时P与G1，M1解绑，P被挂载到M2上继续执行G队列中其他任务。G1解除阻塞后，如果有空闲的P就加入到P队列中，如果没有就放到全局可运行队列runqueue中。P会周期性扫描全局（61次）可运行队列，执行里面的G；如果全局runqueue为空，就会从其他的P的执行队列中取一半G来执行。

在 GPM 模型，有一个全局队列（Global Queue）：存放等待运行的 G，还有一个 P 的本地队列：也是存放等待运行的 G，但数量有限，不超过 256 个。

GPM 的**调度流程**从 go func()开始创建一个 goroutine，新建的 goroutine 优先保存在 P 的本地队列中，如果 P 的本地队列已经满了，则会保存到全局队列中。

M 会从 P 的队列中取一个可执行状态的 G 来执行，如果 P 的本地队列为空，就会从其他的 MP 组合偷取一个可执行的 G 来执行，

当 M 执行某一个 G 时候发生系统调用或者阻塞，M 阻塞，

如果这个时候 G 在执行，runtime 会把这个线程 M 从 P 中摘除，然后创建一个新的操作系统线程来服务于这个 P，当 M 系统调用结束时，这个 G 会尝试获取一个空闲的 P 来执行，并放入到这个 P 的本地队列，如果这个线程 M 变成休眠状态，加入到空闲线程中，然后整个 G 就会被放入到全局队列中。

**work stealing（工作量窃取） 机制**：会优先从全局队列里进行窃取，之后会从其它的P队列里窃取一半的G，放入到本地P队列里。
**hand off （移交）机制**：当前线程的G进行阻塞调用时，例如睡眠，则当前线程就会释放P，然后把P转交给其它空闲的线程执行，如果没有闲置的线程，则创建新的线程

### go中线程VS协程：

-  创建销毁：
   - 协程goroutine由Go runtime负责管理，创建和销毁的销毁都非常小，是用户级
   - 线程创建和销毁开销巨大，因为是内核级的，通常的解决方法是线程池
- 创建数量：
  - 协程：轻松创建上百万个
  - 线程：通常最多不超过1w个
- 内存占用：
	- 协程：2kb，初始分配4k堆栈，随着程序的执行自动增长删除
	- 线程：1M，创建线程是必须指定堆栈且固定，通常M为单位
- 切换成本：
	- 协程：协程切换只需保存3个寄存器，耗时约200纳秒
	- 线程：线程切换需要保存几十个寄存器，耗时约1000纳秒
- 调度方式：
	- 协程：非抢占式，由GO runtime主动交出控制权
	- 线程：在时间片用完后，由CPU中断任务强行将其调度走，此时需要保存很多信息

### gorountine的优势

1. Go程序会智能地将 goroutine 中的任务合理地分配给每个CPU，最大限度的使用cpu的性能
2. 开启一个goroutine消耗是非常小的（大概是2kb），所以可以轻松的创建数以百计的goroutine。
3. 速度快，还可以用channel进行通信

## GMP

[Golang并发调度的GMP模型 - 掘金 (juejin.cn)](https://juejin.cn/post/6886321367604527112)
[GO GMP协程调度实现原理 5w字长文史上最全 - 掘金 (juejin.cn)](https://juejin.cn/post/7105971301151408141)

- G(Goroutine)，表示一个 goroutine，即我需要分担出去的任务；
- M(Machine)，对应一个内核线程，用于将一个 G 搬到线程上执行；
- P(Processor)，一个装满 G 的队列，用于维护一些任务；

-   G：Groutine协程，拥有运行函数的指针、栈、上下文（指的是sp、bp、pc等寄存器上下文以及垃圾回收的标记上下文），在整个程序运行过程中可以有无数个，代表一个用户级代码执行流（用户轻量级线程）；
-   P：Processor，调度逻辑处理器，同样也是Go中代表资源的分配主体（内存资源、协程队列等），默认为机器核数，可以通过GOMAXPROCS环境变量调整
-   M：Machine，代表实际工作的执行者，对应到操作系统级别的线程；M的数量会比P多，但不会太多，最大为1w个。

Golang 中的 GMP 模型是指 Goroutine、M（Machine）、P（Processor）三个组件组成的并发模型，其中 Goroutine 是 Go 语言中的轻量级线程，M 是操作系统线程的抽象，P 则是用于管理 Goroutine 的调度器。P 结构负责管理 Goroutine 的调度，包括 Goroutine 的创建、销毁、挂起和唤醒等，而 M 结构则负责将 Goroutine 绑定到操作系统线程上并执行它们。

![image-202303301522203](https://raw.githubusercontent.com/Swiftie13st/Figurebed/main/img/202305241549788.png)

1.  **全局队列**（Global Queue）：存放等待运行的G。
2.  **P的本地队列**：同全局队列类似，存放的也是等待运行的G，存的数量有限，不超过256个。新建G'时，G'优先加入到P的本地队列，如果队列满了，则会把本地队列中一半的G移动到全局队列。接近于无锁化，但没有达到真正意义的无锁，由于 work-stealing 机制的存在，其他 p 可能会前来执行窃取动作，因此操作仍需加锁.
3.  **P列表**：所有的P都在程序启动时创建，并保存在数组中，最多有`GOMAXPROCS`(可配置)个。
4.  **M**：线程想运行任务就得获取P，从P的本地队列获取G，P队列为空时，M也会尝试从全局队列**拿**一批G放到P的本地队列，或从其他P的本地队列**偷**一半放到自己P的本地队列。M运行G，G执行之后，M会从P获取下一个G，不断重复下去。

（1）M 是线程的抽象；G 是 goroutine；P 是承上启下的调度器；
（2）M调度G前，需要和P绑定；
（3）全局有多个M和多个P，但同时并行的G的最大数量等于P的数量；
（4）G的存放队列有三类：P的本地队列；全局队列；和wait队列（图中未展示，为io阻塞就绪态goroutine队列）；
（5）M调度G时，优先取P本地队列，其次取全局队列，最后取wait队列；这样的好处是，取本地队列时，可以接近于无锁化，减少全局锁竞争；
（6）为防止不同P的闲忙差异过大，设立work-stealing机制，本地队列为空的P可以尝试从其他P本地队列偷取一半的G补充到自身队列.

### 调度

![image-202303091653149](https://raw.githubusercontent.com/Swiftie13st/Figurebed/main/img/202305241549789.png)

调度假设当前正在执行G1，G1阻塞（如系统调用），此时P与G1，M1解绑，P被挂载到M2上继续执行G队列中其他任务。G1解除阻塞后，如果有空闲的P就加入到P队列中，如果没有就放到全局可运行队列runqueue中。P会周期性扫描全局可运行队列，执行里面的G；如果全局runqueue为空，就会从其他的P的执行队列中取一半G来执行。

在 GPM 模型，有一个全局队列（Global Queue）：存放等待运行的 G，还有一个 P 的本地队列：也是存放等待运行的 G，但数量有限，不超过 256 个。

GPM 的**调度流程**从 go func()开始创建一个 goroutine，新建的 goroutine 优先保存在 P 的本地队列中，如果 P 的本地队列已经满了，则会保存到全局队列中。

M 会从 P 的队列中取一个可执行状态的 G 来执行，如果 P 的本地队列为空，就会从其他的 MP 组合偷取一个可执行的 G 来执行，

当 M 执行某一个 G 时候发生系统调用或者阻塞，M 阻塞，

如果这个时候 G 在执行，runtime 会把这个线程 M 从 P 中摘除，然后创建一个新的操作系统线程来服务于这个 P，当 M 系统调用结束时，这个 G 会尝试获取一个空闲的 P 来执行，并放入到这个 P 的本地队列，如果这个线程 M 变成休眠状态，加入到空闲线程中，然后整个 G 就会被放入到全局队列中。

**work stealing（工作量窃取） 机制**：会优先从全局队列里进行窃取，之后会从其它的P队列里窃取一半的G，放入到本地P队列里。
**hand off （移交）机制**：当前线程的G进行阻塞调用时，例如睡眠，则当前线程就会释放P，然后把P转交给其它空闲的线程执行，如果没有闲置的线程，则创建新的线程

### 调度器的设计策略

**复用线程**：避免频繁的创建、销毁线程，而是对线程的复用。

1) `work stealing`机制
当本线程无可运行的G时，尝试从其他线程绑定的P偷取G，而不是销毁线程。  
2) `hand off`机制
当本线程因为G进行系统调用阻塞时，线程释放绑定的P，把P转移给其他空闲的线程执行。

**利用并行**：`GOMAXPROCS`设置P的数量，最多有`GOMAXPROCS`个线程分布在多个CPU上同时运行。`GOMAXPROCS`也限制了并发的程度，比如`GOMAXPROCS = 核数/2`，则最多利用了一半的CPU核进行并行。

**抢占**：在coroutine中要等待一个协程主动让出CPU才执行下一个协程，在Go中，一个goroutine最多占用CPU 10ms，防止其他goroutine被饿死，这就是goroutine不同于coroutine的一个地方。

**全局G队列**：在新的调度器中依然有全局G队列，但功能已经被弱化了，当M执行work stealing从其他P偷不到G时，它可以从全局G队列获取G。

`runnext优先`

### 为什么要有P？

加了 P 之后会带来什么改变呢？

- 每个 P 有自己的本地队列，大幅度的减轻了对全局队列的直接依赖，所带来的效果就是锁竞争的减少。而 GM 模型的性能开销大头就是锁竞争。
- 每个 P 相对的平衡上，在 GMP 模型中也实现了 Work Stealing （工作量窃取机制）算法，如果 P 的本地队列为空，则会从全局队列或其他 P 的本地队列中窃取可运行的 G 来运行，减少空转，提高了资源利用率。

如果在M上实现P类似的组件：

- 一般来讲，M 的数量都会多于 P。像在 Go 中，M 的数量默认是 10000，P 的默认数量的 CPU 核数。另外由于 M 的属性，也就是如果存在系统阻塞调用，阻塞了 M，又不够用的情况下，M 会不断增加。
- M 不断增加的话，如果本地队列挂载在 M 上，那就意味着本地队列也会随之增加。这显然是不合理的，因为本地队列的管理会变得复杂，且 Work Stealing 性能会大幅度下降。

在 Golang 中，为了提高并发性能，一个 M 可以绑定多个 P，并且一个 P 也可以被多个 M 共享。这种设计可以在多核 CPU 上更好地利用硬件资源，从而提高并发性能。

如果直接将 P 结构实现在 M 上，会使得一个 M 管理的 Goroutine 的数量非常大，导致调度器的复杂度增加，可能会影响到并发性能。而将 P 结构单独实现，并允许多个 P 绑定到同一个 M 上，可以有效降低调度器的复杂度，提高并发性能。

另外，将 P 结构单独实现还可以方便地进行负载均衡和资源调度。因为不同的 Goroutine 可能会对系统资源产生不同的需求，单独实现 P 结构可以让调度器更加灵活地进行资源调度，从而提高系统的并发性能。

### 限制P的个数 

可以通过 `runtime.GOMAXPROCS()` 来设定 `P` 的值，当前 Go 版本的 `GOMAXPROCS` 默认值已经设置为 CPU 的（逻辑核）核数， 这允许我们的 Go 程序充分使用机器的每一个 CPU, 最大程度的提高我们程序的并发性能。不过从实践经验中来看，IO 密集型的应用，可以稍微调高 `P` 的个数；

**限制M**：Go语⾔本身是限定M的最⼤量是10000，也可以通过"runtime/debug"包下`debug.SetMaxThreads(1)`设置

### goroutine什么时候会发生阻塞？

1. 由于原子、互斥量或通道操作调用导致 Goroutine 阻塞，调度器将把当前阻塞的 Goroutine 切换出去，重新调度 LRQ 上的其他 Goroutine；
2. 由于网络请求和 IO 操作导致 Goroutine 阻塞。
3. 当调用一些系统方法的时候（如文件 I/O），如果系统方法调用的时候发生阻塞，
4. 如果在 Goroutine 去执行一个 sleep 操作，导致 M 被阻塞了

### 每个线程/协程占用多少内存？

创建一个协程需要2kb, 栈空间不够会自动扩容， 创建一个线程需要1M空间。

### 在GPM调度模型，goroutine 有哪几种状态？线程呢？

**有9种状态**

-   **\_Gidle**：刚刚被分配并且还没有被初始化
-   **\_Grunnable**：没有执行代码，没有栈的所有权，存储在运行队列中
-   **\_Grunning**：可以执行代码，拥有栈的所有权，被赋予了内核线程 M 和处理器 P
-   **\_Gsyscall**：正在执行系统调用，拥有栈的所有权，没有执行用户代码，被赋予了内核线程 M 但是不在运行队列上
-   **\_Gwaiting**：由于运行时而被阻塞，没有执行用户代码并且不在运行队列上，但是可能存在于 Channel 的等待队列上
-   **\_Gdead**：没有被使用，没有执行代码，可能有分配的栈
-   **\_Gcopystack**：栈正在被拷贝，没有执行代码，不在运行队列上
-   **\_Gpreempted**：由于抢占而被阻塞，没有执行用户代码并且不在运行队列上，等待唤醒
-   **\_Gscan**：GC 正在扫描栈空间，没有执行代码，可以与其他状态同时存在

![image-202304241515076](https://raw.githubusercontent.com/Swiftie13st/Figurebed/main/img/202305241549790.png)


### 如果goroutine一直占用资源怎么办，PMG模型怎么解决这个问题

如果有一个goroutine一直占用资源的话，GMP模型会从正常模式转为饥饿模式，通过信号协作强制处理在最前的 goroutine 去分配使用

#### 基于信号的抢占式调度

golang在之前的版本中已经实现了抢占调度，不管是陷入到大量计算还是系统调用，大多可被sysmon扫描到并进行抢占。但有些场景是无法抢占成功的。比如轮询计算 for { i++ } 等，这类操作无法进行newstack、morestack、syscall，所以无法检测stackguard0 = stackpreempt。

go team已经意识到抢占是个问题，所以在1.14中加入了**基于信号的协程调度抢占**。原理是这样的，首先注册绑定 SIGURG 信号及处理方法runtime.doSigPreempt，sysmon会间隔性检测超时的p，然后发送信号，m收到信号后休眠执行的goroutine并且进行重新调度。

sysmon启动后会间隔性的进行监控，最长间隔10ms，最短间隔20us。如果某协程独占P超过`10ms`，那么就会被抢占！

### 如果若干个线程中有一个发生OOM会发生什么？如果goroutine发生呢？

当一个线程抛出OOM异常后，它所占据的内存资源会全部被释放掉，从而不会影响其他线程的运行

go中的内存泄漏一般都是goroutine泄露，就是goroutine没有被关闭，或者没有添加超时控制，让goroutine一只处于阻塞状态，不能被GC。

当一个线程发生OOM（内存溢出）时，通常会导致整个进程崩溃，因为线程共享进程的地址空间。然而，当一个Goroutine发生OOM时，情况可能会有所不同。由于Goroutine的堆栈是动态扩展的，当一个Goroutine的堆栈无法扩展时，Go运行时会尝试回收其他Goroutine的内存，以便为当前Goroutine分配更多的内存。如果回收失败，Go运行时会抛出一个运行时错误（如`runtime: out of memory`），但不会导致整个进程崩溃。

### 什么是协程泄露

协程泄露指的是在 Go 语言程序中，由于某种原因而导致协程无法正常结束，从而造成内存泄露的情况。这种情况通常发生在使用协程时处理异步任务时，如果没有正确地处理协程的终止条件，它们将一直保持活动状态，不断占用内存，最终导致内存泄露。为了避免协程泄露，应该确保在程序结束时关闭所有协程，并释放其占用的内存。

- 可以通过`leaktest`库来检测
- 使用 golang 自带的`pprof`监控工具，可以发现内存上涨情况

1. channel缺少消费者，导致发送阻塞；没有生产者，读取阻塞
2. 死锁
	1.  同一个goroutine中，使用同一个chnnel读写；
	2.  2个 以上的goroutine中， 使用同一个 channel 通信。 读写channel 先于 go程创建；
	3. channel 和 读写锁、互斥锁混用；
3. select 所有的case都阻塞
4. 无限死循环
	I/O 操作上的堵塞也可能造成泄露，例如发送请求到 API 服务器，而没有使用超时；或者程序单纯地陷入死循环中。

### 如果有若干个goroutine，其中有一个panic，会发生什么

有一个panic，那么剩余goroutine也会退出，程序退出。如果不想程序退出，那么必须通过调用 recover() 方法来捕获 panic 并恢复将要崩掉的程序。

### defer可以捕获到其goroutine中的子goroutine的panic吗？

不能,它们处于不同的调度器P中。对于子goroutine，必须通过 **recover() 机制来进行恢复**，然后结合日志进行打印（或者通过channel传递error），下面是一个例子

### 主协程如何等待所有协程都完成

1. `sync.WaitGroup`
   - Add：WaitGroup 类型有一个计数器，默认值是 0，通常通过个方法来标记需要等待的子协程数量
   - Done：当某个子协程执行完毕后，可以通过 Done 方法标记已完成，常用 defer 语句来调用
   - Wait 阻塞当前协程，直到对应 WaitGroup 类型实例的计数器值归零
2. 使用channel
   - 声明一个和子协程数量一致的通道数组，然后为每个子协程分配一个通道元素，在子协程执行完毕时向对应的通道发送数据；然后在主协程中，依次读取这些通道接收子协程发送的数据，只有所有通道都接收到数据才会退出主协程。
3. 使用context
   - 使用 `context.WithCancel`在子协程退出:

### 过起多个协程，怎么控制他们的退出 

channel通知，select监控，ctx中Done退出，waitGroup等待

### 编程题：3个函数分别打印cat、dog、fish，要求每个函数都要起一个goroutine，按照cat、dog、fish顺序打印在屏幕上100次。 

```go
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
		fmt.Println("cat", count)
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

func fishPrint(fishCh *chan struct{}, catCh *chan struct{}) {
	count := 0
	for {
		if count >= countNum {
			wg.Done()
			//fmt.Println("fish quit")
			return
		}
		<-*fishCh
		fmt.Println("fish")
		count++
		*catCh <- struct{}{}
	}
}
```

### 请描述避免多线程竞争时有哪些手段？ 

Go语言提供了传统的同步 goroutine 的机制，就是对共享资源加锁。atomic 和 sync 包里的一些函数就可以对共享的资源进行加锁操作。

#### 原子函数

原子函数能够以很底层的加锁机制来同步访问整型变量和指针。

n++ 可能会出现脏写，n = a, a = a + 1, n = a
解决：把n++封装成原子操作，解除资源竞争，避免脏写
```go
func atomic.AddInt32(addr *int32, delta int32) (new int32)

func atomic.LoadInt32(addr *int32) (val int32)
```

```go
var isInit uint32
if atomic.LoadUint32(&isInit) == 1 {
	return
}
atomic.StoreUint32(&isInit, 1)
```

#### 互斥锁

另一种同步访问共享资源的方式是使用互斥锁，互斥锁这个名字来自互斥的概念。  
互斥锁用于在代码上创建一个临界区，保证同一时间只有一个 goroutine 可以执行这个临界代码。

### 特殊的goroutine  G0

在 Go 中创建的所有 Goroutine 都会被一个内部的调度器所管理。Go 调度器尝试为所有的 Goroutine 分配运行时间，并且在当前的 Goroutine 阻塞或者终止的时候，Go 调度器会通过运行 Goroutine 的方式使所有 CPU 保持忙碌状态。这个调度器实际上是作为一个特殊的 Goroutine 运行的。

**M0**

`M0`是启动程序后的编号为0的主线程，这个M对应的实例会在全局变量runtime.m0中，不需要在heap上分配，M0负责执行初始化操作和启动第一个G， 在之后M0就和其他的M一样了。

**G0**

`G0`是每次启动一个M都会第一个创建的gourtine，G0仅用于负责调度的G，G0不指向任何可执行的函数, 每个M都会有一个自己的G0。在调度或系统调用时会使用G0的栈空间, 全局变量的G0是M0的G0。

职责：
1. Goroutine 创建与调度
2. defer 函数分配。
3. 垃圾收集操作，比如 STW（ stopping the world ），扫描 Goroutine 的栈，以及一些标记清理操作。
4. 栈增长。当需要的时候，Go 会增加 Goroutine 的大小。这个操作是由 `g0` 的 prolog 函数完成的。

### 底层源码分析

https://github.com/golang/go/blob/master/src/runtime/runtime2.go
https://github.com/golang/go/blob/master/src/runtime/proc.go

#### g

（1）g 即goroutine，是 golang 中对协程的抽象；
（2）g 有自己的运行栈、状态、以及执行的任务函数（用户通过 go func 指定）；
（3）g 需要绑定到 p 才能执行，在 g 的视角中，p 就是它的 cpu.

```go
type g struct { 
	goid int64 // 唯一的goroutine的ID 
	m *m //负责当前g的m
	sched gobuf // goroutine切换时，用于保存g的上下文 
	stack stack // 栈 
	gopc // pc of go statement that created this goroutine 
	startpc uintptr // pc of goroutine function 
	// ... 
} 
type gobuf struct {
	sp uintptr // 栈指针位置 
	pc uintptr // 运行到的程序位置 
	g guintptr // 指向 goroutine 
	ret uintptr // 保存系统调用的返回值 
	// ... 
} 
type stack struct {
	lo uintptr // 栈的下界内存地址 
	hi uintptr // 栈的上界内存地址 
}
```

（1）m：在 p 的代理，负责执行当前 g 的 m；
（2）sched.sp：保存 CPU 的 rsp 寄存器的值，指向函数调用栈栈顶；
（3）sched.pc：保存 CPU 的 rip 寄存器的值，指向程序下一条执行指令的地址；
（4）sched.ret：保存系统调用的返回值；
（5）sched.bp：保存 CPU 的 rbp 寄存器的值，存储函数栈帧的起始位置.

#### m

（1）m 即 machine，是 golang 中对线程的抽象；
（2）m 不直接执行 g，而是先和 p 绑定，由其实现代理；
（3）借由 p 的存在，m 无需和 g 绑死，也无需记录 g 的状态信息，因此 g 在全生命周期中可以实现跨 m 执行.

```go
type m struct {    
	g0 *g  // goroutine with scheduling stack    
// ...    
	tls [tlsSlots]uintptr // thread-local storage  (for x86 extern register)    
// ...
}
```

（1）g0：==一类特殊的调度协程，不用于执行用户函数，负责执行 g 之间的切换调度==. 与 m 的关系为 1:1；
（2）tls：thread-local storage，线程本地存储，存储内容只对当前线程可见. 线程本地存储的是 m.tls 的地址，m.tls\[0\] 存储的是当前运行的 g，因此线程可以通过 g 找到当前的 m、p、g0 等信息.

#### p

（1）p 即 processor，是 golang 中的调度器；
（2）p 是 gmp 的中枢，借由 p 承上启下，实现 g 和 m 之间的动态有机结合；
（3）对 g 而言，p 是其 cpu，g 只有被 p 调度，才得以执行；
	（4）对 m 而言，p 是其执行代理，为其提供必要信息的同时（可执行的 g、内存分配情况等），并隐藏了繁杂的调度细节；
（5）p 的数量决定了 g 最大并行数量，可由用户通过 GOMAXPROCS 进行设定（超过 CPU 核数时无意义）.

```go
type p struct {    
	// ...    
	runqhead uint32    
	runqtail uint32    
	runq     [256]guintptr        
	runnext guintptr    
	// ...
}
```

（1）runq：本地 goroutine 队列，最大长度为 256.
（2）runqhead：队列头部；
（3）runqtail：队列尾部；
（4）runnext：下一个可执行的 goroutine.

#### schedt 全局队列

```go
type schedt struct {    
	// ...    
	lock mutex    
	// ...    
	runq     gQueue    
	runqsize int32    
	// ...
}
```

sched 是全局 goroutine 队列的封装：
（1）lock：一把操作全局队列时使用的锁；
（2）runq：全局 goroutine 队列；
（3）runqsize：全局 goroutine 队列的容量.

#### 调度过程

goroutine 的类型可分为两类：

1.  负责调度普通 g 的 g0，执行固定的调度流程，与 m 的关系为一对一；
2. 负责执行用户函数的普通 g.

m 通过 p 调度执行的 goroutine 永远在普通 g 和 g0 之间进行切换，当 g0 找到可执行的 g 时，会调用 gogo 方法，调度 g 执行用户定义的任务；当 g 需要主动让渡或被动调度时，会触发 mcall 方法，将执行权重新交还给 g0.

（1）主动调度：
一种用户主动执行让渡的方式，主要方式是，用户在执行代码中调用了 runtime.Gosched 方法，此时当前 g 会当让出执行权，主动进行队列等待下次被调度执行.
（2）被动调度：
因当前不满足某种执行条件，g 可能会陷入阻塞态无法被调度，直到关注的条件达成后，g才从阻塞中被唤醒，重新进入可执行队列等待被调度.
常见的被动调度触发方式为因 channel 操作或互斥锁操作陷入阻塞等操作，底层会走进 gopark 方法.goready 方法通常与 gopark 方法成对出现，能够将 g 从阻塞态中恢复，重新进入等待执行的状态.
（3）正常调度：
g 中的执行任务已完成，g0 会将当前 g 置为死亡状态，发起新一轮调度.
（4）抢占调度：
倘若 g 执行系统调用超过指定的时长，且全局的 p 资源比较紧缺，此时将 p 和 g 解绑，抢占出来用于其他 g 的调度. 等 g 完成系统调用后，会重新进入可执行队列中等待被调度.
值得一提的是，前 3 种调度方式都由 m 下的 g0 完成，唯独抢占调度不同.
因为发起系统调用时需要打破用户态的边界进入内核态，此时 m 也会因系统调用而陷入僵直，无法主动完成抢占调度的行为.
因此，在 Golang 进程会有一个==全局监控协程 monitor g ==的存在，这个 g 会越过 p 直接与一个 m 进行绑定，不断轮询对所有 p 的执行状况进行监控. 倘若发现满足抢占调度的条件，则会从第三方的角度出手干预，主动发起该动作.

![image-202304241542137](https://raw.githubusercontent.com/Swiftie13st/Figurebed/main/img/202305241549791.png)

#### 抢占调度

抢占调度的执行者不是 g0，而是一个全局的 monitor g，代码位于 runtime/proc.go 的 retake 方法中：

1. 加锁后，遍历全局的 p 队列，寻找需要被抢占的目标：
2. 倘若某个 p 同时满足下述条件，则会进行抢占调度：
	1. 执行系统调用超过 10 ms；
	2. p 本地队列有等待执行的 g；
	3. 或者当前没有空闲的 p 和 m.
3. 抢占调度的步骤是，先将当前 p 的状态更新为 idle，然后步入 handoffp 方法中，判断是否需要为 p 寻找接管的 m（因为其原本绑定的 m 正在执行系统调用）
4. 当以下三个条件满足其一时，则需要为 p 获取新的 m：
	1. 当前 p 本地队列还有待执行的 g；
	2. 全局繁忙（没有空闲的 p 和 m，全局 g 队列为空）
	3. 需要处理网络 socket 读写请求
5. 获取 m 时，会先尝试获取已有的空闲的 m，若不存在，则会创建一个新的 m.

Go 进程在启动的时候，会开启一个后台线程 sysmon，监控执行时间过长的 goroutine，进而发出抢占。另一方面，GC 执行 stw 时，会让所有的 goroutine 都停止，其实就是抢占。这两者都会调用 `preemptone()` 函数。

`preemptone()` 函数会沿着下面这条路径：

```go
preemptone->preemptM->signalM->tgkill
```

向正在运行的 goroutine 所绑定的的那个 M（也可以说是线程）发出 `SIGURG` 信号。

1.  M 注册一个 SIGURG 信号的处理函数：sighandler。
2.  sysmon 线程检测到执行时间过长的 goroutine、GC stw 时，会向相应的 M（或者说线程，每个线程对应一个 M）发送 SIGURG 信号。
3.  收到信号后，内核执行 sighandler 函数，通过 pushCall 插入 asyncPreempt 函数调用。
4.  回到当前 goroutine 执行 asyncPreempt 函数，通过 mcall 切到 g0 栈执行 gopreempt_m。
5.  将当前 goroutine 插入到全局可运行队列，M 则继续寻找其他 goroutine 来运行。
6.  被抢占的 goroutine 再次调度过来执行时，会继续原来的执行流。
