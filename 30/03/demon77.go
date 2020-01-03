package main

import (
	"fmt"
	"strings"
)

func main() {
	reader1 := strings.NewReader("NewReader returns a new Reader reading from s. " +
		"It is similar to bytes.NewBufferString but more efficient and read-only.")
	fmt.Printf("The size of reader: %d\n", reader1.Size())
	fmt.Printf("The reading index in reader: %d\n",
		reader1.Size()-int64(reader1.Len()))
	
	buf1 := make([]byte, 47)
	n, _ := reader1.Read(buf1)
	fmt.Printf("%d bytes were read. (callRead)\n",n)
	fmt.Printf("The reading index in reader: %d]\n",
		reader1.Size()-int64(reader1.Len()))
	fmt.Println("",)

	buf2 := make([]byte, 21)
	offset1 := int64(64)
	n, _ = reader1.ReadAt(buf2, offset1)
	fmt.Printf("%d bytes were read. (call ReadAt, offset: %d)\n",n, offset1)
	fmt.Printf("The reading index in reader: %d\n",
		reader1.Size()-int64(reader1.Len()))
	fmt.Println("",)
	
	
	
	
	

}