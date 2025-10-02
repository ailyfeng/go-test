package main

import "fmt"

func modify(p *[]int){
	for i:=range(*p){
		(*p)[i] *=2
	}
}

func main(){
	/*
题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
考察点 ：指针运算、切片操作。

	*/

	var s =[]int{1,2,3,4}
	

	modify(&s)
	fmt.Println(s)
}