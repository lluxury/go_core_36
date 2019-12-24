package main

import "fmt"

var channels = [3]chan int{  // 切处通道
	nil,
	make(chan int),
	nil,
}

var numbers = []int{1,2,3} // 切片

func main()  {
	select {
	case getChan(0) <- getNumber(0):
		fmt.Println("The first candidate case is selected.")
	case getChan(1) <- getNumber(1):
		fmt.Println("The second candidate case is selected.")
	case getChan(2) <- getNumber(2):
		fmt.Println("The third candidate case is selected")
	default:
		fmt.Println("No candidate case is selected!")
	}
// 匹配通道，从切片里获取值，依次遍历了一遍
}

func getNumber(i int) int  {
	fmt.Printf("numbers[%d]\n",i)
	return numbers[i]

}

func getChan(i int) chan int  {
	fmt.Printf("channels[%d]\n",i)
	return channels[i]

}

