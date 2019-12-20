
package main

import (
	"flag"
	"github.com/lluxury/go_core_36/03/02/lib"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	//lib5.Hello(name)
	lib.Hello(name)
}