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
	var err error
	_, err = exec.LookPath(os.DevNull)
	fmt.Printf("error: %s\n", err)
	if execErr, ok := err.(*exec.Error); ok {
		execErr.Name = os.TempDir()
		execErr.Err = os.ErrNotExist
	}
	fmt.Printf("error: %s\n", err)
	fmt.Println()
}