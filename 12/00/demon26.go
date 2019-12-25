package main

import "fmt"

type Printer func(contents string) (n int, err error)
//函数 类型，参数，结果声明

func printToStd(contents string) (bytesNum int, err error) {
	return fmt.Println(contents)
}
//函数

func main()  {
	var p Printer
	p = printToStd	// 特征一样，我们是兄弟？
	p("something")
}

