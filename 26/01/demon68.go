package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("Enter function main.")
	coordinateWithContext()
}

func coordinateWithContext()  {
	total := 12
	var num int32
	fmt.Printf("The number: %d [with context.Context]\n", num)
	cxt, canceFunc := context.WithCancel(context.Background())
	for i := 1; i <=total ; i++ {
	  go addNum(&num, i, func() {
	  	if atomic.LoadInt32(&num) == int32(total){
	  		canceFunc()
		}
	  })  // 匿名函数，虽然用的多，但还是强调一下，addNum的参数之一
	}
	<-cxt.Done()
	fmt.Println("End.",)
}

func addNum(numP *int32, id int, deferFunc func())  {
	defer func() {
		deferFunc()
	}()
	for i := 0; ; i++ {
	  currNum := atomic.LoadInt32(numP)
	  newNum := currNum + 1
	  time.Sleep(time.Millisecond * 200)
	  if atomic.CompareAndSwapInt32(numP,currNum,newNum){
	  	fmt.Printf("The number: %d [%d-%d]\n", newNum,id,i)
	  	break
	  } else {
		  //fmt.Printf("The CAS operation failed. [%d-%d]\n", id, i)
	  }
	}
}


