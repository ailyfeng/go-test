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

# 2 任务2

## 2.1 指针

### 2.1.1、 指针操作

题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
考察点 ：指针的使用、值传递与引用传递的区别。

[解题](homework_for_basic_learning/task2/point/task1/main.go)

```go

func modify(p *int) {
	*p += 10
}
```

### 2.1.2、slice
题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
考察点 ：指针运算、切片操作。

[解题](homework_for_basic_learning/task2/point/task2/main.go)

```go

func modify(p *[]int){
	for i:=range(*p){
		(*p)[i] *=2
	}
}

```

## 2.2 Goroutine

### 2.2.1 协程
题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。
[解题](task2/goroutine/task1/main.go)

```go

// 打印奇数
func printOdd(nums []int,wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Print("\n打印奇数:")
	for _,n:=range nums{
		t:=n%2
		if(t==1){
			fmt.Printf("%d\t",n)
		}
	}
}

// 打印偶数
func printEven(nums []int,wg *sync.WaitGroup){
	wg.Done()
	fmt.Print("\n打印偶数:")
	for _,n:=range nums{
		tmp:=n%2
		if(tmp==0){
			fmt.Printf("%d\t",n)
		}
	}
}
```

### 2.2.2 多任务调度器

	题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
	考察点 ：协程原理、并发任务调度。

[解题](task2/goroutine/task2/main.go)

```go

type Task func()

type TaskInfo struct{
	Name string
	Duration time.Duration
}

func ExecuteTasks(tasks map[string]Task) (result []TaskInfo){
	var wg sync.WaitGroup
	var mu sync.Mutex

	for taskName,task:=range tasks{
		wg.Add(1)
		go func(taskName string,t Task){
			defer wg.Done()
			startTime := time.Now()
			t()
			duration :=time.Since(startTime)

			mu.Lock()
			result=append(result, TaskInfo{
				Name: taskName,
				Duration:duration,
			})
			mu.Unlock()

		}(taskName,task)
	}
	wg.Wait()
	return result
}

```
## 2.3、面向对象

### 2.3.1、接口的实现
题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。
[解题](task2/object/task1/main.go)

```go

const PI float32 =3.141592657

type Shape interface{
	Area() float32
	Perimeter() float32
}

type Rectangle struct{
	Width float32
	Height float32
	
}

type Circle struct{
	Radius int
}

func (rec Rectangle) Area() float32{
	return rec.Width*rec.Height
}

func (rec Rectangle) Perimeter() float32{
	return (rec.Height+rec.Width)*2
}

func (cir Circle) Area() float32{
	return float32(cir.Radius*cir.Radius)*PI
}


func (rec Circle) Perimeter() float32{
	return float32(rec.Radius)*2*PI
}

```


### 2.3.2、结构体的组合
题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
考察点 ：组合的使用、方法接收者。

[解题](task2/object/task2/main.go)
```go

type Person struct{
	Name string
	Age int
}

type Employee struct{
	Person
	EmployeeID int
}

func (emp Employee) PrintInfo(){
	// fmt.Printf("EmployeeID:\t%d\nName:\t\t%s\nAge:\t\t%d\n",emp.EmployeeID,emp.Name,emp.Age)
	fmt.Printf("EmployeeID:\t%d\n",emp.EmployeeID)
	fmt.Printf("Name:\t\t%s\n",emp.Name)
	fmt.Printf("Age:\t\t%d\n",emp.Age)
}

```

## 2.4、Channel

### 2.4.1、通道1
题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
考察点 ：通道的基本使用、协程间通信。

[解题](task2/Channel/task1/main.go)

```go

	var wg sync.WaitGroup

	ch:=make(chan int,10)
	wg.Add(1)
	go func (){
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <-i
		}
		close(ch)
	}()

	wg.Add(1)
	go func(){
		defer wg.Done()
		for num:=range ch{
			fmt.Println(num)
		}
	}()

	wg.Wait()

```

### 2.4.2、通道2
题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。

[解题](task2/Channel/task2/main.go)

```go

func producer(ch chan<- int,wg *sync.WaitGroup){
	defer wg.Done()
	for i:=0;i<100;i++{
		ch<-i
	}
	close(ch)
}

func consumer(ch <-chan int,wg *sync.WaitGroup){
	defer wg.Done()
	for num:=range ch{
		fmt.Println(num)
	}
}

```
