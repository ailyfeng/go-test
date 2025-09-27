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
func isPalindromeOther(x int)bool{
	if(x<0 || (x%10==0 && x!=0)){
		return false
	}
	res:=0
	for x>res{
		res=res*10+x%10
		x/=10

	}
	return x==res || x==res/10
}


func main() {
	/*
	回文数

	考察：数字操作、条件判断
	题目：判断一个整数是否是回文数 

	回文数（Palindrome Number）是指一个数字，无论从左往右读还是从右往左读，其数值都完全相同的数。
	这个概念来源于“回文”（Palindrome），即正着读和反着读都一样的字符串（如 "level"、"radar"、"上海海上"）。
	*/
	
	var randNumber int =123321// rand.Intn(5)
	fmt.Println(randNumber)

	// flag:=isPalindrome(randNumber)
	flag:=isPalindromeOther(randNumber)
	
	if flag {
		fmt.Println("回文数")
	} else {
		fmt.Println("非回文数")
	}



}