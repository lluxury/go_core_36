package main

import (
	"fmt"
	"os"
)

type argDesc struct {
	action	string
	flag 	int
	perm 	os.FileMode
} 

func main() {
	fmt.Printf("The mode for dir:\n%32b\n", os.ModeDir)
	fmt.Printf("The mode for named pipe:\n%32b\n", os.ModeNamedPipe)
	fmt.Printf("The mode for all of the irregular files:\n%32b\n", os.ModeType)
	fmt.Printf("The mode for permissions:\n%32b\n", os.ModePerm)
	fmt.Println("",)
	

	
	
	

}