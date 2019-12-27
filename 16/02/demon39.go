package main

import (
	"fmt"
)

func main()  {
	num := 10
	sign := make(chan struct{}, num)
	for i := 0; i < num; i++ {
	  go func() {
	  	fmt.Println(i)
	  	sign <- struct{}{}
	  }()
	}
	for j := 0; j < num; j++ {
	  <-sign
	}  // 值还是不确认，这是另一个问题
}