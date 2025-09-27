
# 作业

## 1、控制流程

只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。

 [使用异或方式解答，但是数组元素中，仅且有个元素出现一次，且其他元素都出现两次。](homework_for_basic_learning/control_flow/task1/main.go)

时间复杂度 = O(n)
空间复杂度 = O(n)

 ```go

func findOnlyOne(nums []int) int{
	
	m:=make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	var res int
	for k,v:=range m{
		if v==1{
			res= k
		}
	}
	return res
}
```
优化：
时间复杂度 O(n)
空间复杂度 O(1)

 ```go

    func findOnlyOne2(nums []int) int{ 
        res := 0
        for _,num:= range nums {
            res =res^num
        }
        return res
    }

 ```

 ## 2、回文数

考察：数字操作、条件判断
题目：判断一个整数是否是回文数 

回文数（Palindrome Number）是指一个数字，无论从左往右读还是从右往左读，其数值都完全相同的数。

[使用字符串切片方式解答](homework_for_basic_learning/control_flow/task2/main.go)

自己写的代码：

时间复杂度	O(log x)	

空间复杂度：O(log x) strconv.Itoa(x) 创建了一个长度为 log x 的字符串

```go

func isPalindrome(x int) bool{

	// 换成字符串
	randNumberB:=strconv.Itoa(x)
	tmp:=[]byte(randNumberB)
	lenght:=len(tmp)

	for i,v:=range tmp{
		if v != tmp[lenght-1-i] {
			return false
		}
	}
	return true
}
```

优化：

时间复杂度	O(log x)	只处理一半的位数，增长极慢

空间复杂度	O(1)	只用一个变量，不额外开空间

```go

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
```
 
