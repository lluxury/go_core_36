package main

import (
	"flag"
	"github.com/lluxury/go_core_36/03/04/lib"
	"os"
	//in "puzzlers/article3/q4/lib/internal" // 此行无法通过编译。
	//"os"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	lib.Hello(name)
	//in.Hello(os.Stdout, name)
}