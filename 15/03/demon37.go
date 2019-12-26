package main

import (
	"fmt"
	"unsafe"
)

type Dog struct {
	name string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func main(){
	dog := Dog{"little pig"}
	dogP := &dog
	//xx := unsafe.Pointer(dogP) 为了看中间值，就是 =（中间部分
	dogPtr := uintptr(unsafe.Pointer(dogP))   // 入口

	namePtr := dogPtr + unsafe.Offsetof(dogP.name)
	nameP := (*string)(unsafe.Pointer(namePtr)) // 字段指针

	//fmt.Printf("nameP == &(dogP.name)? %v\n", nameP = &(dogP.name))
	fmt.Printf("nameP == &(dogP.name)? %v\n",
		nameP == &(dogP.name))
	fmt.Printf("The name of dog is %q.\n", *nameP) // 字段值

	*nameP = "monster"			// 修改，骇
	fmt.Printf("The name of dog is %q.\n", dogP.name)
	fmt.Println()


	// 下面这种不匹配的转换虽然不会引发panic，但是其结果往往不符合预期。
	numP := (*int)(unsafe.Pointer(namePtr))
	num := *numP
	fmt.Printf("This is an unexpected number: %d\n", num)

}