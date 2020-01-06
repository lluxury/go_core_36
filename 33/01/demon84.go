package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	comment := "Package bufio implements buffered I/O. " +
		"It wraps an io.Reader or io.Writer object, " +
		"creating another object (Reader or Writer) that " +
		"also implements the interface but provides buffering and " +
		"some help for textual I/O."
	basicReader := strings.NewReader(comment)
	fmt.Printf("The size of basic reader: %d\n",basicReader.Size())
	fmt.Println("",)

	//fmt.Println("New a buffered reader ...")
	reader1 := bufio.NewReader(basicReader)
	//fmt.Printf("The default size of buffered reader: %d\n",reader1.Size())
	//// 缓冲未填充
	fmt.Printf("The number of unread bytes in the buffer: %d\n",reader1.Buffered())
	//fmt.Println("",)
	
	bytes, err := reader1.Peek(7)
	if err != nil {
		fmt.Printf("error: %v\n",err)
	}
	fmt.Printf("Peeked contents(%d): %q\n",len(bytes),bytes)
	fmt.Printf("The number of unread bytes in the buffer: %d\n",reader1.Buffered())
	fmt.Println("",)   // 缓冲填满
	
	buf1 := make([]byte,7)
	n, err := reader1.Read(buf1)
	if err != nil {
		fmt.Printf("error: %v\n",err)
	}
	fmt.Printf("Read contents(%d): %q\n",n,buf1)
	fmt.Printf("The number of unread bytes in the buffer: %d\n",reader1.Buffered())
	fmt.Println("",)  // reader1 是 bufio 

	
	
	

	


}