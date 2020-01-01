package main

import (
	"fmt"
	"sync"
	//"sync/atomic"
	"time"
)

func main() {
	//var counter uint32
	//var once sync.Once
	//once.Do(func() {
	//	atomic.AddUint32(&counter, 1)
	//})
	//fmt.Printf("The counter: %d\n",counter)
	//once.Do(func() {
	//	atomic.AddUint32(&counter,2)
	//})
	//fmt.Printf("The counter: %d\n", counter)
	//fmt.Println("",)
	
	once := sync.Once{}
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		once.Do(func() {
			for i := 0; i < 3; i++ {
				fmt.Printf("Do task. [1-%d]\n",	i)
				time.Sleep(time.Second)
			}
		})
		fmt.Println("Done. [1]",)
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 500)
		once.Do(func() {
			fmt.Println("Do task. [2]",)
		})
		fmt.Println("Done. [2]",)
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 500)
		once.Do(func() {
			fmt.Println("Do task, [3]",)
		})
		fmt.Println("Done. [3]",)
	}()
	wg.Wait()
	fmt.Println("",)
}