package main

import (
	"fmt"
	"sync"
)

func main(){

	/*
	
题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。
	*/

	var wg sync.WaitGroup
	var mu sync.Mutex
	var m=map[int]int{}

	for i := 0; i < 10; i++ {
		//  m[i]=0. // 这行一定不能写，因为它是在goroutine 外部执行的。会出现main goroutine 和 worker goroutine 同时写 map 的局面。
		wg.Add(1)
		go func(idx int){
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()
			for j:=0;j<1000;j++{
				m[idx]++
			}

		}(i)
	}

	wg.Wait()

	for k,v:=range m{
		fmt.Println(k,v)
	}
}