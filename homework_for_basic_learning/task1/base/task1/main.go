package main

import (
	"fmt"
	"strings"
	"time"
)

/*
- 时间复杂度：O(n^2) 暴力破解
- 空间复杂度：O(1)
*/
func twoSum(nums []int,target int)[]int{
	if len(nums)<2{
		return []int{}
	}

	
	for i:=0;i<len(nums);i++{
		
		for j := i+1; j < len(nums); j++ {
			if(nums[i]+nums[j]==target){
				return []int{i,j}
			}
		}
	}
	return []int{}
}

/*
- 时间复杂度：O(n)
- 空间复杂度：O(n)
*/
func twoSum1(nums []int,target int)[]int{
	m:=make(map[int]int)
	for index,v:=range(nums){
		if res,exsits:=m[target-v];exsits{
			return []int{res,index}
		}
		m[v]=index
	}
	return nil
}

type Arr struct{
	Target int
	Nums []int
}
func main(){
	var arr = map[int]Arr{
		0:Arr{Target:9,Nums:[]int{2,7,11,15}},
		1:Arr{Target:6,Nums:[]int{3,2,4}},
		2:Arr{Target:6,Nums:[]int{}},
		3:Arr{Target:9,Nums:[]int{11,15,2,7}},
	}


	for _,nums:=range(arr){
    start := time.Now() // 记录开始时间
		fmt.Println(strings.Repeat("-",50))
		res:=twoSum1(nums.Nums,nums.Target)
		fmt.Printf("执行时间：%v\n",time.Since(start))
		fmt.Println(res,nums.Nums,nums.Target)
	}
}
