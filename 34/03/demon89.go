package main

import (
	"fmt"
	"os"
	"path/filepath"
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

	fileName1 := "something3.txt"
	filepath1 := filepath.Join(os.TempDir(),fileName1)
	fmt.Printf("The file path: %s\n", filepath1)

	argDescList := []argDesc{
		{
			"Create",
			os.O_RDONLY| os.O_CREATE,
			0644,
		},
		{
			"Reuse",
			os.O_RDONLY|os.O_TRUNC,
			0666,
		},
		{
				"Open",
				os.O_RDONLY|os.O_APPEND,
				0777,
		},
	}

	defer os.Remove(filepath1)
	for _, v := range argDescList {
		fmt.Printf("%s the file with perm %o ...\n", v.action, v.perm)
		file1, err := os.OpenFile(filepath1, v.flag, v.perm)
		if err != nil {
			fmt.Printf("error: %v\n	", err)
			continue
		}
		info1, err := file1. Stat()
		if err != nil {
			fmt.Printf("error: %v\n", err)
			continue
		}
		fmt.Printf("The file permissions: %o\n", info1.Mode().Perm())
	}
}