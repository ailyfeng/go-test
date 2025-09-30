
# 1、控制流程

## 1.1、只出现一次的数字

只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。

 [使用异或方式解答，但是数组元素中，仅且有个元素出现一次，且其他元素都出现两次。](homework_for_basic_learning/control_flow/task1/main.go)

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

 ## 1.2、回文数

考察：数字操作、条件判断
题目：判断一个整数是否是回文数 

回文数（Palindrome Number）是指一个数字，无论从左往右读还是从右往左读，其数值都完全相同的数。

[使用字符串切片方式解答](homework_for_basic_learning/control_flow/task2/main.go)

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

# 2、字符串的操作

## 2.1、有效的括号 
考察：字符串处理、栈的使用

题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效

链接：https://leetcode-cn.com/problems/valid-parentheses/

[使用栈方式解答](homework_for_basic_learning/string/task1/main.go)

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

## 2.2、最长公共前缀

考察：字符串处理、循环嵌套

题目：查找字符串数组中的最长公共前缀

链接：https://leetcode-cn.com/problems/longest-common-prefix/ 


[使用嵌套循环方式解答](homework_for_basic_learning/string/task2/main.go)

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

# 3、基本值类型

## 3.1、加一 

难度：简单

考察：数组操作、进位处理

题目：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一

链接：https://leetcode-cn.com/problems/plus-one/ 

[解答](homework_for_basic_learning/base_type/task1/main.go)


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

# 4、引用类型：切片

# 4.1、 删除有序数组中的重复项


删除有序数组中的重复项：给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。

链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/


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

	index:=1 // 不相等的索引值递增，第一元素视为独立值
	for i:=1;i<numsLen;i++{
		if(nums[i]!=nums[i-1]){ // 比较当前值和前一个值不相等
			nums[index]=nums[i]// 不相等的值放在前面
			index++ // 不相等的索引值递增
		}
	}
	fmt.Println(nums)
	return  index
}

```