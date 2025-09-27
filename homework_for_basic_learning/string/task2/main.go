package main

import "fmt"

func longestCommonPrefix(strs []string) string{
	if len(strs)==0{
		return ""
	}
	var prefix string = strs[0]

	for _,str:=range strs{
		for i:=0;i<len(prefix);i++{
			if i>=len(str) || prefix[i]!=str[i]{
				prefix = prefix[:i]
				break
			}
		}

	}
	return prefix
}

func main() {
  
	/*
	最长公共前缀

	考察：字符串处理、循环嵌套

	题目：查找字符串数组中的最长公共前缀

	https://leetcode.cn/problems/longest-common-prefix/description/

	思路：
	1. 获取数组中第一个元素，作为最长的前缀
	2. 循环遍历数组，将最长的前缀字符串和其他元素做字符串比较，将公共部分作为最长的前缀
	*/

	arr := []string{"flower", "flow", "floght"}

	longestPrefix := longestCommonPrefix(arr)
	fmt.Println(longestPrefix)


}