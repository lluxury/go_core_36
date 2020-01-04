package main

import (
	"bytes"
	"fmt"
)

func main() {
	var contents string



	var buffer3 bytes.Buffer
	fmt.Printf("The length of buffer: %d\n", buffer3.Len())
	fmt.Printf("The capacity of buffer: %d\n", buffer3.Cap())
	fmt.Print("\n\n")

	contents = "xyz"
	fmt.Printf("Write contents %q ...\n", contents)
	buffer3.WriteString(contents)
	fmt.Printf("The length of buffer: %d\n", buffer3.Len())
	fmt.Printf("The capacity of buffer: %d\n", buffer3.Cap())

}