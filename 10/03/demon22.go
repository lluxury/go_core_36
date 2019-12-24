package main

import "fmt"

func main()  {
	ch1 := make(chan int, 2)
	go func(){
		for i := 0; i < 10; i++ {
			fmt.Printf("Sender: sending element %v...\n",i)
			ch1 <-i
		}
		fmt.Println("Sender: close the channel...")
		close(ch1)
	}()

	for{
		elem,ok := <-ch1
		if !ok {
			fmt.Println("Receiver: closed channel")
			break
		}
		fmt.Printf("Receiver: received an element: %v\n",elem)
	}
	fmt.Println("End.")

}



//赫林老师最牛逼的就是，把这些知识还写出例子了，yann 光理解就精疲力竭了。