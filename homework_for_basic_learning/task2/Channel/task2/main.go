package main

import (
	"fmt"
	"sync"
)


func producer(ch chan<- int,wg *sync.WaitGroup){
	defer wg.Done()
	for i:=0;i<100;i++{
		ch<-i
	}
	close(ch)
}

func consumer(ch <-chan int,wg *sync.WaitGroup){
	defer wg.Done()
	for num:=range ch{
		fmt.Println(num)
	}
}
func main(){
	/*

题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。
	*/

	var ch =make(chan int,100)
	var wg sync.WaitGroup

	wg.Add(1)
	go producer(ch,&wg)

	wg.Add(1)
	go consumer(ch,&wg)

	wg.Wait()


}