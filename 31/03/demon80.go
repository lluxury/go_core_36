package main

import (
	"fmt"
	"bytes"
)

func main() {
	// contents := "ab"
	// buffer1 := bytes.NewBufferString(contents)
	// fmt.Printf("The capacity of new buffer with contents %q: %d\n", 
	// 	contents, buffer1.Cap())
	// fmt.Println("")   // 原始
	
	// unreadBytes := buffer1.Bytes()
	// fmt.Printf("The unread bytes of the buffer: %v\n", unreadBytes)
	// fmt.Println("")   // 未读取

	// contents = "cdefg"
	// fmt.Printf("Write contents %q ...\n", contents)
	// buffer1.WriteString(contents)
	// fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap())
	// fmt.Println("")   // 写入，容量8

	// unreadBytes = unreadBytes[:cap(unreadBytes)]
	// fmt.Printf("The unread bytes of the buffer: %v\n", unreadBytes)
	// fmt.Println("")  // 获取未读

	// value := byte('X')
	// fmt.Printf("Set a byte in the unread bytes to %v ...\n", value)
	// unreadBytes[len(unreadBytes)-2] = value
	// fmt.Printf("The unread bytes of the buffer: %v\n", buffer1.Bytes())
	// fmt.Println("") // 修改

	// contents = "hijklmn"
	// fmt.Printf("Write contents %q ...\n", contents)
	// buffer1.WriteString(contents)
	// // fmt.Printf("The unread bytes of the buffer: %v\n", unreadBytes)
	// fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap())
	// fmt.Println() // 自动扩充

	// unreadBytes = unreadBytes[:cap(unreadBytes)]
	// fmt.Printf("The unread bytes of the buffer: %v\n", unreadBytes)
	// fmt.Println("")  // 无法使用

	contents := "12"
	buffer2 := bytes.NewBufferString(contents)
	fmt.Printf("The capacity of new buffer with contents %q: %d\n", 
		contents, buffer2.Cap())
	fmt.Println("")

	nextBytes := buffer2.Next(2)   // 12 拿到的切片
	fmt.Printf("The next bytes of the buffer: %v\n", nextBytes)
	fmt.Println("")

	contents = "34567"
	fmt.Printf("Write contents %q ...\n", contents)
	buffer2.WriteString(contents)
	fmt.Printf("The capacity of buffer: %d\n", buffer2.Cap())
	fmt.Println("")

	nextBytes = nextBytes[:cap(nextBytes)]  // 扩充切片
	fmt.Printf("The next bytes of the buffer: %v\n", nextBytes)
	fmt.Println("")

	value := byte('X')
	fmt.Printf("Set a byte in the next bytes to %v ...\n", value)
	nextBytes[len(nextBytes)-2] = value
	fmt.Printf("The unread bytes of the buffer: %v\n", buffer2.Bytes())
	fmt.Println("")

	contents = "89101112"
	fmt.Printf("Write contents %q ...\n", contents)
	buffer2.WriteString(contents)
	fmt.Printf("The capacity of buffer: %d\n", buffer2.Cap())
	fmt.Println("")

	nextBytes = nextBytes[:cap(nextBytes)]
	fmt.Printf("The next bytes of the buffer: %v\n", nextBytes)
}