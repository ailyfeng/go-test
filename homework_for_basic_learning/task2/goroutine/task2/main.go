package main

import (
	"fmt"
	"sync"
	"time"
)

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
func main(){
	/*
	题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
	考察点 ：协程原理、并发任务调度。
	*/

	tasks :=map[string]Task{
		"爬虫抓取":func() {
			fmt.Println("任务[爬取网页]：开始...")
			time.Sleep(2 * time.Second)
			fmt.Println("任务[爬取网页]：完成")
		},
		"邮件发送":func() {
			fmt.Println("任务[邮件发送]：开始...")
			time.Sleep(1 * time.Second)
			fmt.Println("任务[邮件发送]：完成")
		},
		"短信发送":func() {
			fmt.Println("任务[短信发送]：开始...")
			time.Sleep(3 * time.Second)
			fmt.Println("任务[短信发送]：完成")
		},
		"消息队列":func() {
			fmt.Println("任务[消息队列]：开始...")
			time.Sleep(2 * time.Second)
			fmt.Println("任务[消息队列]：完成")
		},
		"群公告发送":func() {
			fmt.Println("任务[群公告发送]：开始...")
			time.Sleep(4 * time.Second)
			fmt.Println("任务[群公告发送]：完成")
		},
		"活动推送":func() {
			fmt.Println("任务[活动推送]：开始...")
			time.Sleep(1 * time.Second)
			fmt.Println("任务[活动推送]：完成")
		},
		
	}
	fmt.Println("🚀 开始并发执行任务...\n")
	res:=ExecuteTasks(tasks)
	fmt.Println("🚀 结束并发执行任务...\n")

	for _,v:=range res{
		fmt.Printf("  📌 %s: %v\n",v.Name,v.Duration)
	}

}