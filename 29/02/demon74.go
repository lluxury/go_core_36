package main

import "fmt"

func main() {
	str := "运维专家谈转型go"
	for i, c := range str {
		//fmt.Printf("%d: %q [% x]\n",  i, c, []byte(string(c)))
		fmt.Printf("%d: %q [% x]\n",  i, c, []byte(string(c)))
	}
}