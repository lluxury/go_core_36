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
	r, w, err := os.Pipe()
	if err != nil {
		fmt.Printf("unexpected error: %s\n", err)
		return
	}
	r.Close()
	_, err = w.Write([]byte("hi"))
	uError := underlyingError(err)
	fmt.Printf("underlying error: %s (type: %T)\n", uError,uError)
	fmt.Println()
	}