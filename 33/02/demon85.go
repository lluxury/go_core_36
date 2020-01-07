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

	size := 300
	fmt.Printf("New a buffered reader with size %d ...\n",size)
	reader1 := bufio.NewReaderSize(basicReader,size)
	fmt.Println("",)

	fmt.Print("[ About 'Peek' method]\n\n")
	peekNum := 38
	fmt.Printf("Peek %d bytes ...\n",peekNum)
	bytes , err := reader1.Peek(peekNum)
	if err != nil {
		fmt.Printf("error: %v\n",err)
	}
	fmt.Printf("Peeked contents(%d): %q\n", len(bytes),bytes)
	fmt.Printf("The number of unread bytes in the buffer: %d\n",reader1.Buffered())
	fmt.Println("",)

	//fmt.Print("[ About 'Read' method ]\n\n")


	
	
	



}