package main

import "fmt"

func main()  {
	a1 := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("a1: %v (len: %d, cap: %d)\n",
		a1, len(a1), cap(a1))
	s9:=a1[1:4]  // 索引值，对应 2及后，5之前
	fmt.Printf("s9: %v (len: %d, cap: %d)\n",
		s9, len(s9), cap(s9))
	for i := 1; i < 5; i++ {
	  s9 = append(s9,i)
		fmt.Printf("s9(%d): %v (len: %d, cap: %d)\n",
			i, s9, len(s9), cap(s9))
	}
	fmt.Printf("a1:%v (len:%d, cap: %d)\n",
		a1, len(a1), cap(a1))
	// 原来的数组变了
	fmt.Println()

}



//赫林老师最牛逼的就是，把这些知识还写出例子了，yann 光理解就精疲力竭了。