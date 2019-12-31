package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var box atomic.Value
	fmt.Println("Copy box to box2.",)
	box2 := box
	v1 := [...]int{1, 2, 3}
	fmt.Printf("Store %v to box.\n", v1)
	box.Store(v1)
	fmt.Printf("The value load from box is %v.\n", box.Load())
	fmt.Printf("The value load from box2 is %v.\n", box2.Load())
	fmt.Println("",)




}