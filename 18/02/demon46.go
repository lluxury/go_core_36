package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func underlyingError(err error) error {
	switch err := err.(type) {
	case *os.PathError:
		return err.Err
	case *os.LinkError:
		return err.Err
	case *os.SyscallError:
		return err.Err
	case *exec.Error:
		return err.Err
	}
	return err
}

type Errno int

func (e Errno) Error() string {
	return "errno" + strconv.Itoa(int(e))
}

func main()  {

	const (
		ERR0 = Errno(0)
		ERR1 = Errno(1)
		ERR2 = Errno(2)
	)
	var myErr error = Errno(0)
	switch myErr {
	case ERR0:
		fmt.Println("ERR0")
	case ERR1:
		fmt.Println("ERR1")
	case ERR2:
		fmt.Println("ERR2")
	}
}
