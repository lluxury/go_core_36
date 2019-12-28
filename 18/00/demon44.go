package main

import (
	"errors"
	"fmt"
)

func echo(request string) (response string, err error)  {
	if request == "" {
		err = errors.New("empty request")
		return
	}
	response = fmt.Sprintf("echo: %s", response)
	return
}

func main() {

	fmt.Println()
	err1 := fmt.Errorf("invalid contents: %s", "#$%")
	err2 := errors.New(fmt.Sprintf("invalid contents: %s", "#$%"))
	if err1.Error() == err2.Error() {
		fmt.Println("The error messages in err1 and err2 are the same.")
	}
	}