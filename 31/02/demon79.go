package main

import (
	"bytes"
	"fmt"
)

func main() {
	var contents string
	buffer1 := bytes.NewBufferString(contents)
	fmt.Printf("The length of new buffer with contents %q: %d\n", 
		contents,buffer1.Len())
	fmt.Printf("The capacity of new buffer with contents %q: %d\n", 
		contents,buffer1.Cap())	
	fmt.Println("")

	contents = "12345"
	// fmt.Printf("Write contents $q ...\n", contents)
	fmt.Printf("Write contents %q ...\n", contents)
	buffer1.WriteString(contents)
	fmt.Printf("The length of buffer: %d\n", buffer1.Len())
	fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap())
	fmt.Println("")

	contents = "67"
	fmt.Printf("Write contents %q ...\n", contents)
	buffer1.WriteString(contents)
	fmt.Printf("The length of buffer: %d\n", buffer1.Len())
	fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap())
	fmt.Println("")

	contents = "89"
	fmt.Printf("Write contents %q ...\n", contents)
	buffer1.WriteString(contents)
	fmt.Printf("The length of buffer: %d\n", buffer1.Len())
	fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap())
	fmt.Println("")
}