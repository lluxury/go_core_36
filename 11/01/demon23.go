package main

import (
	"fmt"
	"math/rand"

)

func main()  {
	var uselessChan = make(chan<- int, 1)  // 只发不收
	var anotherUselessChan = make(<-chan int, 1) // 只收不发
	fmt.Printf("The useless channels: %v, %v\n",
		uselessChan,anotherUselessChan)

	intChan1 := make(chan int,3)
	SendInt(intChan1)

	intChan2 := getIntChan()
	for elem := range intChan2{  // 只能从通道接收
		fmt.Printf("The element in intChan2: %v\n", elem)
	}

	_ = GetIntChan(getIntChan)
}

//func SendInt(ch chan<- int)  {
//	ch <- rand.Int(1000)
//}

func SendInt(ch chan<- int) {
	ch <- rand.Intn(1000)
}
// 定义函数，约束参数

type Notifier interface {
	SendInt(ch chan<- int)
}
// 定义接口，约束方法

func getIntChan() <-chan int {	// 声明结果
	num := 5
	ch := make(chan int, num)
	for i := 0; i < num; i++ {   // 带有range子句的for语句
		ch <- i
	}
	close(ch)
	return ch
}

type GetIntChan func() <-chan int    // 声明函数类型