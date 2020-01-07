package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	comment := "Writer implements buffering for an io.Writer object. " +
		"If an error occurs writing to a Writer, " +
		"no more data will be accepted and all subsequent writes, " +
		"and Flush, will return the error. After all data has been written, " +
		"the client should call the Flush method to guarantee all data " +
		"has been forwarded to the underlying io.Writer."
	basicWriter1 := &strings.Builder{}

	size := 300
	fmt.Printf("New a buffered writer with size %d ...\n", size)
	//writer1 := bufio.NewReaderSize(basicWriter1,size)
	writer1 := bufio.NewWriterSize(basicWriter1, size)
	fmt.Println("",)


	begin, end := 0, 53
	fmt.Printf("Write %d bytes into the writer ...\n", end-begin)
	writer1.WriteString(comment[begin:end])
	fmt.Printf("The number of buffered bytes: %d\n",writer1.Buffered())
	fmt.Printf("The number of unused bytes in the buffer: %d\n",writer1.Available())
	fmt.Println("Flush the buffer in the writer ...")
	writer1.Flush()
	fmt.Printf("The number of buffered bytes: %d\n",writer1.Buffered())
	fmt.Printf("The number of unrused bytes in the buffer: %d\n",writer1.Available())
	fmt.Println("",)





}