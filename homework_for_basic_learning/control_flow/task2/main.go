package main

import (
	"fmt"
	"strconv"
)


func isPalindrome(x int) bool{

	// 换成字符串
	randNumberB:=strconv.Itoa(x)
	tmp:=[]byte(randNumberB)
	lenght:=len(tmp)

	var flag bool=true
	for i,v:=range tmp{
		if v != tmp[lenght-1-i] {
			flag=false
		}
	}
	return flag
}


func main() {
	/*
	回文数

	考察：数字操作、条件判断
	题目：判断一个整数是否是回文数 

	回文数（Palindrome Number）是指一个数字，无论从左往右读还是从右往左读，其数值都完全相同的数。
	这个概念来源于“回文”（Palindrome），即正着读和反着读都一样的字符串（如 "level"、"radar"、"上海海上"）。
	*/
	
	var randNumber int =1234321// rand.Intn(50000)
	fmt.Println(randNumber)

	flag:=isPalindrome(randNumber)
	if flag {
		fmt.Println("回文数")
	} else {
		fmt.Println("非回文数")
	}
}