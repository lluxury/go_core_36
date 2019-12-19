
package main

import (
	// 需在此处添加代码。[1]
	"flag"
	"fmt"
	"os"
)

var name string
//var name = flag.String("name", "everyone", "The greeting object.")

func init() {
	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr,"Usage of %s:\n","yann")
		flag.PrintDefaults()
	}
	flag.StringVar(&name,"name","everyyone","The greeting object.")
	//flag.String()
}

func main() {

	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}
