package main

import (
	"fmt"
	"time"
)

func main()  {
	num := 10
	for i := 0; i < num; i++ {
	  go func() {
	  	fmt.Println(i)
	  }()
	}
	time.Sleep(time.Millisecond * 500) // 随机出现数字，我的电脑比较慢
}