package main

import (
	"fmt"
	"os"
	"os/exec"
)

func underlyingError(err error) error {
	switch err := err.(type) {
	case *os.PathError:
		return err.Err   // 识别类型，返回类型对应方法
	case *os.LinkError:
		return err.Err
	case *os.SyscallError:
		return err.Err
	case *exec.Error:
		return err.Err
	}
	return err
}
func main(){
	_, _, err := os.Pipe()

	paths := []string{
		os.Args[0], // 当前文件
		"/it/must/not/exist", // 肯定不存在的目录。
		os.DevNull,  // 存在的目录
	}
	printError := func(i int, err error) {
		if err != nil {
			fmt.Println("nil errors")
			return
		}
		err = underlyingError(err)
		switch err {
		case os.ErrClosed:
			fmt.Printf("error(closed)[%d]: %s\n", i, err)
		case os.ErrInvalid:
			fmt.Printf("error(invalid)[%d]: %s\n", i, err)
		case os.ErrPermission:
			fmt.Printf("error(permission)[%d]: %s\n", i, err)
		}
	}
	var f *os.File
	var index int
	{
		index = 0
		f, err := os.Open(paths[index])
		if err != nil {
			fmt.Printf("unexpected error: %s\n", err)
			return
		}
		//制造错误为 os.ErrClosed
		f.Close()
		_, err = f.Read([]byte{})
		printError(index,err)
	}

	{
		index = 1
		//制造 os.ErrInvalid 错误
		f, _ = os.Open(paths[index])
		_, err = f.Stat()
		printError(index,err)
	}
	{
		index = 2
		//制造 os.ErrPermission 的错误
		_, err = exec.LookPath(paths[index])
		printError(index,err)
	}
	if f != nil {
		f.Close()
	}
	fmt.Println()

	}