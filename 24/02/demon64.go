package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var box atomic.Value
	//fmt.Println("Copy box to box2.",)
	box2 := box
	v1 := [...]int{1, 2, 3}
	//fmt.Printf("Store %v to box.\n", v1)
	box.Store(v1)
	//fmt.Printf("The value load from box is %v.\n", box.Load())
	//fmt.Printf("The value load from box2 is %v.\n", box2.Load())
	//fmt.Println("",)

	v2 := "123"
	//fmt.Printf("Store %q to box2.\n",v2)
	box2.Store(v2)
	//fmt.Printf("The value load from box is %v.\n", box.Load())
	//fmt.Printf("The value load from box2 is %q.\n", box2.Load())
	//fmt.Println()

	fmt.Println("Copy box to box3.",)
	box3 := box
	fmt.Printf("The value load from box3 is %v.\n",box3.Load())
	v3 := 123
	fmt.Printf("Store %d to box2.\n", v3)
	box3.Store(v3)
	_ = box3
	fmt.Println("",)
	
	

}