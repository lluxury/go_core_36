package main

import (
	"errors"
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	greeting, err := hello(name)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	fmt.Println(greeting, introduce())
}

func hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	return fmt.Sprintf("Hello, %s!", name), nil
	//S 只带一个值，Printf 带2个
}

func introduce() string {
	return "Welcome to my Golang column."
}