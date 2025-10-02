package main

import "fmt"

func modify(p *int) {
	*p += 10
}
func main() {
	/*
		题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
		考察点 ：指针的使用、值传递与引用传递的区别。
	*/
	var a int = 10

	// var p *int
	p := &a
	modify(p)
	fmt.Printf("a=%d\tp=%d\n",a,*p)

}
