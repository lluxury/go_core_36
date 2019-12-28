package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
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

	paths2 := []string{
		runtime.GOROOT(), // Go根目录
		"/it/must/not/exist", // 肯定不存在的目录。
		os.DevNull,  // 存在的目录
	}
	printError2 := func(i int, err error) {
		if err != nil {
			fmt.Println("nil errors")
			return
		}
		err = underlyingError(err)
		if os.IsExist(err) {
			fmt.Printf("error(closed)[%d]: %s\n", i, err)
		} else if os.IsNotExist(err){
			fmt.Printf("error(invalid)[%d]: %s\n", i, err)
		} else if os.IsPermission(err) {
			fmt.Printf("error(permission)[%d]: %s\n", i, err)
		} else {
			fmt.Printf("error(other)[%d]: %s\n", i, err)
		}
	}
	var f *os.File
	var index int
	{
		index = 0
		err = os.Mkdir(paths2[index], 0700)
		printError2(index,err)
	}

	{
		index = 1
		f, err = os.Open(paths2[index])
		printError2(index,err)
	}
	{
		index = 2
		_, err = exec.LookPath(paths2[index])
		printError2(index,err)
	}
	if f != nil {
		f.Close()
	}
	fmt.Println()

	}