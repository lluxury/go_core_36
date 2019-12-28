package main

import (
	"errors"
	"fmt"
)

func main()  {
	fmt.Println("Enter function main.")
	fmt.Printf("no panic: %v\n", recover())
	panic(errors.New("somethine wrong"))
	p := recover()
	fmt.Printf("panic: %s\n", p)

	fmt.Println("Exit function main.")

}