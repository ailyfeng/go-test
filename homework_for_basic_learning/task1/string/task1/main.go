package main

import "fmt"


func isValid(s string) bool{
	var stack []rune
	for _,v:=range s{
		if v=='(' || v=='{' || v=='['{
			stack=append(stack, v)
		}else{
			stackLen:=len(stack)-1
			if len(stack)==0{
				return false
			}else if v==')' && stack[stackLen]=='(' {
				stack=stack[:stackLen]
			}else if v=='}' && stack[stackLen]=='{'{
				stack=stack[:stackLen]
			}else if v==']' && stack[stackLen]=='['{
				stack=stack[:stackLen]
			}else{
				return false
			}

		}
	}

	return len(stack)==0
}
 
func main(){

	/*
	考察：字符串处理、栈的使用
	题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
	https://leetcode.cn/problems/valid-parentheses/

	思路：利用栈的先进先出的特性来处理
	1. 创建一个栈，用于存储左括号
	2. 遍历字符串，如果遇到左括号，将其压入栈中
	3. 如果遇到右括号，则判断栈是否为空，如果不为空，则将栈顶的左括号弹出，否则返回false
	*/

	var s string = "([{}]){}()"

	res:=isValid(s)
	fmt.Println(res)
 
}