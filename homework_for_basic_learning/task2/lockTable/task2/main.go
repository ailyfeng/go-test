package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	/*
		题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
		考察点 ：原子操作、并发数据安全。
	*/

	var wg sync.WaitGroup
	var m [10]atomic.Int32 // 使用 atomic.Int32 数组，每个协程一个计数器
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				m[idx].Add(1)
			}
		}(i)
	}

	wg.Wait()
	for k,v:=range m {
		fmt.Printf("%d计数器最终值:%d\n", k,v.Load())
	}
}