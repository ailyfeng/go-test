package main

import (
	"fmt"
	"math/rand"
	"strings"
)

/**
 * 测试生成随机数组
 */
func generateArray(n int) []int{
	arr :=make([]int,n)
	for i:=0;i<n;i++{
		arr[i] = rand.Intn(n/2)
	}
	return arr
}

/**
 * 查找数组中只出现一次的元素
 */
func FindOnce(paramMap map[int]int) []int {
	arr:=[]int{}
	for element,total:=range paramMap{
		if total==1{
			arr = append(arr, element)
		}
	}
	return arr
}
/**
 * 统计数组元素出现的次数
 */
func elementCount(arr []int) map[int]int{
	result :=make(map[int]int,len(arr))
	for _,v:=range arr{
		if _,exists :=result[v]; exists{
			result[v] ++
		}else{
			result[v] = 1
		}
	}
	return result
}

func printLine(n int){
	fmt.Println(strings.Repeat("-",n))
}

// 打印5列N行的矩阵
func printMatrix(arr []int){
	printLine(35)

	for i,v:=range arr{
		if i%5 == 0{
			fmt.Println()
		}
		fmt.Printf("%d\t",v)
	}
	fmt.Println()
	printLine(35)
}
func main(){

	/*

	136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
	找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
	例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
	*/
	arr := generateArray(10) // 随机生成10个元素
	fmt.Println("随机数组：")
	printMatrix(arr)

	// 数组去重
	result := elementCount(arr)

	fmt.Println(result)

	onceArr :=FindOnce(result)
	if(len(onceArr)==0){
		fmt.Println("没有只出现一次的元素")
	}else{
		fmt.Println("只出现一次的元素：")
		printMatrix(onceArr)
	}

}