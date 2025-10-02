package main

import (
	"fmt"
	"sync"
)

// 打印奇数
func printOdd(nums []int,wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Print("\n打印奇数:")
	for _,n:=range nums{
		t:=n%2
		if(t==1){
			fmt.Printf("%d\t",n)
		}
	}
}

// 打印偶数
func printEven(nums []int,wg *sync.WaitGroup){
	wg.Done()
	fmt.Print("\n打印偶数:")
	for _,n:=range nums{
		tmp:=n%2
		if(tmp==0){
			fmt.Printf("%d\t",n)
		}
	}
}

func main(){

	/*
	题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
	考察点 ： go 关键字的使用、协程的并发执行。
	*/
	var nums = []int{1,2,3,4,5,6,7,8,9,10}
	var wg sync.WaitGroup
	wg.Add(2)
	go printOdd(nums ,&wg)
	go printEven(nums,&wg)
	wg.Wait()
}
