package main

import (
	"fmt"
	"sort"
	"strings"
)

func merge1(intervals [][]int) [][]int {

	length:=len(intervals)
	if length<=1 {
		return intervals
	}

	// 排序
	sort.Slice(intervals,func (j,i int)  bool{
		return intervals[j][0]<intervals[i][0]
	})
	// 初始化返回结果
	var merged =[][]int{intervals[0]}
	for i:=1;i<length;i++{
		last :=&merged[len(merged)-1] // 指针读取，方便修改或替换
		cur :=intervals[i]
		if((*last)[1]>=cur[0]){ // 如果放回的结果值低一个元素 大于 当前值的开始值
			(*last)[1]=max((*last)[1],cur[1]) // 指针修改直接替换
		}else{
			merged =append(merged,cur) // 追加到结果值中
		}
		
	}

	return merged
}

func merge(intervals [][]int) (merged [][]int){
	
	if len(intervals)<=1 {
		return intervals
	}

	// 排序
	sort.Slice(intervals,func (j,i int)  bool{
		return intervals[j][0]<intervals[i][0]
	})
	
	// 初始化返回结果
	// merged:=[][]int{intervals[0]}
	
	for _,cur:=range(intervals){
		length:=len(merged)
		if length>0 && cur[0]<=merged[length-1][1]{ 
			merged[length-1][1]=max(merged[length-1][1],cur[1])
		}else{
			merged =append(merged,cur) // 追加到结果值中
		}
		
	}

	return merged
	

}


func main(){

	/*
 	合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
	请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
	可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，
	将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。

	https://leetcode.cn/problems/merge-intervals/description/

	解题思路：
		排序数组
		当前元素和结果集中最后一个元素比较：当前元素的最后一个值大于结果集中最后一个元素的第一个元素值时，就替换，否则追加到返回结果集中
	*/


	var arr =[][][]int{
		{{0,0},{2,6},{8,10},{15,18}},
		{{1,3},{2,6},{8,10},{15,18}},
		{{1,4},{0,0}},
		{{1,2},{3,4},{5,6},{7,8},{1,23}},
		{{1,4},{0,2},{3,5}},
		{{1,5}},
		{{2,3},{5,5},{2,2},{3,4},{3,4}},
	}

	for _,nums:=range(arr){
		fmt.Println(strings.Repeat("-",50))
		// fmt.Println(nums)
		res:=merge(nums)
		fmt.Println(res)
	}


}