
package main

import "fmt"

var container = []string{"zero","one","two"} // 切片

func main()  {
	container := map[int] string{0: "zero",1: "yann",2: "two"} // 数组
	fmt.Printf("The element is %q.\n",container[1])
}

//不同类型,结果没报错,还输出来了