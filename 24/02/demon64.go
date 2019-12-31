package main

import (
	"errors"
	"fmt"
	"io"
	"os"
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

	//fmt.Println("Copy box to box3.",)
	//box3 := box
	//fmt.Printf("The value load from box3 is %v.\n",box3.Load())
	//v3 := 123
	//fmt.Printf("Store %d to box2.\n", v3)
	//box3.Store(v3)
	//_ = box3
	//fmt.Println("",)
	
	var box4 atomic.Value
	v4 := errors.New("something wrong")
	fmt.Printf("Store an error with message %q to box4.\n",v4)
	box4.Store(v4)
	v41 := io.EOF
	fmt.Println("Store a value of the same type to box4",)
	box4.Store(v41)
	v42, ok := interface{}(&os.PathError{}).(error)
	if ok {
		fmt.Printf("Store a value of type %T that implements error interface to box4.\n", v42)
		//box4.Store(v42)		// 类型不同，
		fmt.Println("",)
		
	}

	
	

}