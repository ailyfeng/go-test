package main

import "fmt"

func findOnlyMutl(nums []int) []int{
	
	m:=make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	res:=[]int{}
	if(len(m)==0){
		return res
	}
	for k,v:=range m{
		if v==1{
			res=append(res,k)
		}
	}
	return res
}

func findOnlyOne2(nums []int) int{ 
 	res := 0
    for _,num:= range nums {
        res =res^num
    }
    return res
}

func findOnlyOne(nums []int) (int, bool){
	
	var res int
	if(len(nums)==0){
		return res,false
	}
	m:=make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	for k,v:=range m{
		if v==1{
			return k,true
		}
	}
	return res,false
}

func main(){

	/*

	136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
	找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
	例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
	*/
	arr := []int{2,9,2,3,3}
	m:=findOnlyOne2(arr)
	fmt.Println(m)
 
}