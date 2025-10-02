# 1、任务1
## 1.1、控制流程

### 1.1.1、只出现一次的数字

只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。

 [使用异或方式解答，但是数组元素中，仅且有个元素出现一次，且其他元素都出现两次。](task1/control_flow/task1/main.go)

* 时间复杂度 = O(n)
* 空间复杂度 = O(n)


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

* 时间复杂度 O(n)
* 空间复杂度 O(1)


 ```go

    func findOnlyOne2(nums []int) int{ 
        res := 0
        for _,num:= range nums {
            res =res^num
        }
        return res
    }

 ```

 ### 1.1.2、回文数

考察：数字操作、条件判断
题目：判断一个整数是否是回文数 

回文数（Palindrome Number）是指一个数字，无论从左往右读还是从右往左读，其数值都完全相同的数。

[使用字符串切片方式解答](task1/control_flow/task2/main.go)

自己写的代码：

* 时间复杂度	O(log x)	

* 空间复杂度：O(log x) strconv.Itoa(x) 创建了一个长度为 log x 的字符串

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

* 时间复杂度	O(log x)	只处理一半的位数，增长极慢

* 空间复杂度	O(1)	只用一个变量，不额外开空间

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

## 1.2、字符串的操作

### 1.2.1、有效的括号 
考察：字符串处理、栈的使用

题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效

链接：https://leetcode-cn.com/problems/valid-parentheses/

[使用栈方式解答](task1/string/task1/main.go)

* 时间复杂度 = O(n). 遍历字符串 s
* 空间复杂度 = O(n) 创建一个栈

```go

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
```

### 1.2.2、最长公共前缀

考察：字符串处理、循环嵌套

题目：查找字符串数组中的最长公共前缀

链接：https://leetcode-cn.com/problems/longest-common-prefix/ 


[使用嵌套循环方式解答](task1/string/task2/main.go)

* 时间复杂度 = O(n*m) n为字符串数组的长度，m为字符串数组中字符串的长度
* 空间复杂度 = O(1) 不使用额外空间

```go

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
```

## 1.3、基本值类型

### 1.3.1、加一 

难度：简单

考察：数组操作、进位处理

题目：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一

链接：https://leetcode-cn.com/problems/plus-one/ 

[解答](task1/base_type/task1/main.go)


- 时间复杂度：O(N)
- 空间复杂度：O(N) 这里数组元素都是9的情况下，数组长度就是n+1,否则是O(1)

```go

func plusOne(digits []int)[]int{
	digitsLen:=len(digits)-1
	for i:=digitsLen;i>=0;i--{
		if(digits[i]<9){
			digits[i]++
			return digits
		}
		digits[i]=0
	}
	return append([]int{1},digits...)
	
}
```

## 1.4、引用类型：切片

### 1.4.1、 删除有序数组中的重复项


删除有序数组中的重复项：给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。

链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/

[解答](task1/slice/task1/main.go)

解题思路
	把当前元素和前一个元素比较，不相等就放在前面（需要一个不相等的索引值递增）


- 时间复杂度：O(N)
- 空间复杂度：O(1)

```go


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
	return  index
}

```

### 1.4.2 合并区间
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

 

示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2：

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
示例 3：

输入：intervals = [[4,7],[1,4]]
输出：[[1,7]]
解释：区间 [1,4] 和 [4,7] 可被视为重叠区间。
 

提示：

1 <= intervals.length <= 104
intervals[i].length == 2
0 <= starti <= endi <= 104

解题思路

- 排序数组
- 当前元素和结果集中最后一个元素比较：当前元素的最后一个值大于结果集中最后一个元素的第一个元素值时，就替换，否则追加到返回结果集中

[代码](task1/slice/task2/main.go)

- 时间复杂度：O(n Log n)
- 空间复杂度：O(2n)
```go

func merge(intervals [][]int) [][]int {

	length:=len(intervals)
	if length<=1 {
		return intervals
	}

	// 排序 O(n Log n)
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

```

优化

- 时间复杂度：O(n Log n)
- 空间复杂度：O(n)

```go

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

```

## 1.3、 基础

### 1.3.1、两数之和 


考察：数组遍历、map使用

题目：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数

链接：https://leetcode-cn.com/problems/two-sum/ 


[代码](task1/base/task1/main.go)

自写
- 时间复杂度：O(n^2) 暴力破解
- 空间复杂度：O(1)

```go


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

```

优化
- 时间复杂度：O(n)
- 空间复杂度：O(n)


```go
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

```

