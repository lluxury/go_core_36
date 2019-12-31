package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {

	//num := uint32(18)
	//fmt.Printf("The number: %d\n", num)
	//delta := int32(-3)   // 转有符号并赋值
	//atomic.AddUint32(&num, uint32(delta))
	//fmt.Printf("The number: %d\n", num)
	//
	//atomic.AddUint32(&num, ^uint32(-(-3)-1))  // 直接用补码，推荐
	//fmt.Printf("The number: %d\n", num)
	//
	//fmt.Printf("The two's complement of %d: %b\n",
	//	delta, uint32(delta))
	//fmt.Printf("The equivalent: %b\n", ^uint32(-(-3)-1))
	//fmt.Println()

	forAndCAS1()
	fmt.Println("",)
	forAndCAS2()

}

func forAndCAS1()  {  // 简易的自旋锁
	sign := make(chan struct{},2)
	num := int32(0)
	fmt.Printf("The number: %d\n", num)
	go func() {
		defer func() {
			sign <- struct{}{}
		}()
		for {
			time.Sleep(time.Millisecond * 500)
			newNum := atomic.AddInt32(&num, 2)	// 定时增加 num
			fmt.Printf("The number: %d\n", newNum)
			if newNum == 10 {
				break
			}
		}
	}()
	go func() {
		defer func() {
			sign <- struct{}{}
		}()
		for {
			if atomic.CompareAndSwapInt32(&num,10,0){  // 逢10 归0
				fmt.Println("The number has gone to zero.",)
				break
			}
			time.Sleep(time.Millisecond * 500)
		}
	}()
	<-sign
	<-sign
}

func forAndCAS2()  { // 简易的互斥锁
	sign := make(chan struct{},2)
	num := int32(0)
	fmt.Printf("The number: %d\n", num)
	max := int32(20)
	go func(id int, max int32) { // 定时增加num的值
		defer func() {
			sign <- struct{}{}
		}()
		for i := 0; ; i++ {
			currNum := atomic.LoadInt32(&num)
			if currNum >= max {
				break
			}
			newNum := currNum + 2
			time.Sleep(time.Millisecond * 200)
			if atomic.CompareAndSwapInt32(&num,currNum,newNum){
				fmt.Printf("The number : %d [%d-%d]\n", newNum, id,i)
			} else {
				fmt.Printf("The CAS operation failed. [%d-%d]\n", id, i)
			}
		}
	}(1,max)
	go func(id int, max int32) { // 定时增加num的值，2个都是需求，轮流加
		defer func() {
			sign <- struct{}{}
		}()
		for j := 0; ; j++ {
			currNum := atomic.LoadInt32(&num)
			if currNum >= max{
				break
			}
			newNum := currNum + 2
			time.Sleep(time.Millisecond * 200)
			if atomic.CompareAndSwapInt32(&num,currNum,newNum) {
				fmt.Printf("The number: %d [%d-%d]\n", newNum, id, j)
			} else {
				fmt.Printf("The CAS operation failed. [%d-%d]\n", id, j)
			}
		}
	}(2,max)
	<-sign
	<-sign
}

