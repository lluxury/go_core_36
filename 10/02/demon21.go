package main

func main()  {
	ch1 := make(chan int, 1)
	ch1 <-1
	//ch1 <-2

	ch2 := make(chan int, 1)
	//elem,ok :=<-ch2
	//_, _ = elem,ok
	ch2 <-1

	var ch3 chan int
	//ch3 <-1
	//<-ch3
	_ = ch3
}