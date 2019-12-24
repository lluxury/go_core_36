package main

import "fmt"

func main()  {
	ch1 := make(chan int, 3)
	ch1 <-2
	ch1 <- 1  // 空格是容许的
	ch1 <- 3
	elem1 := <-ch1
	elem1 = <-ch1
	fmt.Printf("The first element received form channel ch1:%v\n",
		elem1)

}