
package main

import (
	"flag"
	"fmt"
)

func main() {
	 //string // [1]
	//flag.StringVar(&name, "name", "everyone", "The greeting object.") // [2]
	var name = flag.String( "name", "everyone", "The greeting object.") // [2]
	//返回指针型,看提示窗的最右边
	flag.Parse()
	//fmt.Printf("Hello, %v!\n", name)
	fmt.Printf("Hello, %v!\n", *name)
}