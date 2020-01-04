package main

import (
	"fmt"
	"bytes"
)

func main() {
	contents := "ab"
	buffer1 := bytes.NewBufferString(contents)
	fmt.Printf("The capacity of new buffer with contents %q: %d\n", 
		contents, buffer1.Cap())
	fmt.Println("")   // 原始
	
	unreadBytes := buffer1.Bytes()
	fmt.Printf("The unread bytes of the buffer: %v\n", unreadBytes)
	fmt.Println("")   // 未读取

	contents = "cdefg"
	fmt.Printf("Write contents %q ...\n", contents)
	buffer1.WriteString(contents)
	fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap())
	fmt.Println("")   // 写入，容量8

	unreadBytes = unreadBytes[:cap(unreadBytes)]
	fmt.Printf("The unread bytes of the buffer: %v\n", unreadBytes)
	fmt.Println("")  // 获取未读

	value := byte('X')
	fmt.Printf("Set a byte in the unread bytes to %v ...\n", value)
	unreadBytes[len(unreadBytes)-2] = value
	fmt.Printf("The unread bytes of the buffer: %v\n", buffer1.Bytes())
	fmt.Println("") // 修改

	contents = "hijklmn"
	fmt.Printf("Write contents %q ...\n", contents)
	buffer1.WriteString(contents)
	// fmt.Printf("The unread bytes of the buffer: %v\n", unreadBytes)
	fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap())
	fmt.Println() // 自动扩充

	unreadBytes = unreadBytes[:cap(unreadBytes)]
	fmt.Printf("The unread bytes of the buffer: %v\n", unreadBytes)
	fmt.Println("")  // 无法使用
}