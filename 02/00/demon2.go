
package main

import (
	// 需在此处添加代码。[1]
	"flag"
	"fmt"
)

var name string
//var name = flag.String("name", "everyone", "The greeting object.")

func init() {
	// 需在此处添加代码。[2]
	flag.StringVar(&name,"name","everyyone","The greeting object.")
	//flag.String()
}

func main() {
	// 需在此处添加代码。[3]
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}
