/*
要求实现一个 map:
(1) 面向高并发;
(2) 只存在插入和查询操作 0(1);
(3)查询时，若 key 存在，直接返回,val; 若 key 不存在，阻塞直到 key val 对被放入后，获取 val 返回;等待指定时长仍未放入，返回超时错误;
(4)写出真实代码，不能有死锁或者 panic 风险
*/

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type MyConcurrentMap struct {
	sync.Mutex
	mp      map[int]int
	keyToCh map[int]chan struct{}
}

func NewMyConcurrentMap() *MyConcurrentMap {
	return &MyConcurrentMap{
		mp:      make(map[int]int),
		keyToCh: make(map[int]chan struct{}),
	}
}

func (m *MyConcurrentMap) Put(k, v int) {
	m.Lock()
	defer m.Unlock()
	m.mp[k] = v
	ch, ok := m.keyToCh[k]
	if !ok {
		// 没有当前k的ch说明没有需要当前kv对的g在阻塞等待
		return
	}
	// ch <- struct{}{}不能只放入一个值，如果有多个g在等待，则只会唤醒一个，因此选择关闭，则其他等待的g会获取到默认零值，注意考虑不要重复关闭
	//1 通过select去读，当读到了（零值）说明被关闭了, 2 声明一个mychan，内部使用sync.Once进行关闭
	select {
	case <-ch:
		return
	default:
		close(ch)
	}

}

func (m *MyConcurrentMap) Get(k int, timeOut time.Duration) (int, error) {
	m.Lock()
	v, ok := m.mp[k]
	if ok {
		m.Unlock()
		return v, nil
	}

	ch, ok := m.keyToCh[k]
	if !ok {
		ch = make(chan struct{})
		m.keyToCh[k] = ch
	}

	tCtx, cancel := context.WithTimeout(context.Background(), timeOut) // 超时控制
	defer cancel()
	m.Unlock()
	select { // 阻塞
	case <-tCtx.Done():
		return -1, tCtx.Err()
	case <-ch:
	}
	m.Lock()
	v = m.mp[k]
	m.Unlock()
	return v, nil
}

func main() {
	mp := NewMyConcurrentMap()
	for i := 0; i < 1000; i++ {
		go func(i int) {
			v, err := mp.Get(i, time.Duration(time.Microsecond*time.Duration(i)))
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(v)
		}(i)

		go func(i int) {
			time.Sleep(time.Microsecond)
			mp.Put(i, 100)
		}(i)
	}
}
