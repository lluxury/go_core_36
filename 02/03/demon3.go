
package main

import (
	// 需在此处添加代码。[1]
	"flag"
	"fmt"
	"os"
)

var name string
//var name = flag.String("name", "everyone", "The greeting object.")
var cmdLine = flag.NewFlagSet("question",flag.ExitOnError)

func init() {
	cmdLine.StringVar(&name,"name","everyyone","The greeting object.")
	//flag.StringVar(&name,"name","everyyone","The greeting object.")
	//flag.String()
}

func main() {
	cmdLine.Parse(os.Args[1:])
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}
