package main

import "fmt"

func main() {
	str := "运维专家谈转型go"
	fmt.Printf("The string: %q\n",str)
	fmt.Printf(" => runes(char):%q\n",[]rune(str))   // 单字
	fmt.Printf(" => runes(hex):%x\n",[]rune(str))    // 代码
	fmt.Printf(" => bytes(hex):%x\n",[]byte(str))    // 一串代码
}