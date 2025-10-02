package main

import "fmt"

/*
时间复杂度：O(N)
空间复杂度：O(1)
*/
func removeDuplicates(nums []int) int{
	numsLen :=len(nums)
	if(numsLen==0){
		return 0
	}

	index:=1 // 不相等的索引值递增，第0元素视为独立值
	for i:=1;i<numsLen;i++{
		if(nums[i]!=nums[i-1]){ // 比较当前值和前一个值不相等
			nums[index]=nums[i]// 不相等的值放在前面
			index++ // 不相等的索引值递增
		}
	}
	fmt.Println(nums)
	return  index
}
func main(){

	nums :=[]int{0,0,1,1,1,2,2,3,3,4}
	// 解题思路
	// 把当前元素和前一个元素比较，不相等就放在前面（需要一个不相等的索引值递增）
	

	res:=removeDuplicates(nums)
	fmt.Println(res)
}