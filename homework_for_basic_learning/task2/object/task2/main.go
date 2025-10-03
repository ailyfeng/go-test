package main

import "fmt"

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

func main(){
	/*
题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
考察点 ：组合的使用、方法接收者。
	*/

	emp:=Employee{}
	emp.Age=22
	emp.EmployeeID=10001
	emp.Name="历史剧"

	emp.PrintInfo()

}