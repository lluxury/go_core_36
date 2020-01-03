package main

import "strings"

func main() {
	var builder1 strings.Builder
	builder1.Grow(1)
	//
	//f1 := func(b strings.Builder) {
	//	//b.Grow(1)   // 函数
	//}
	//f1(builder1)
	//
	//ch1 := make(chan strings.Builder, 1)
	//ch1 <- builder1
	//builder2 := <- ch1
	////builder2.Grow(1)	// 通道
	//_ = builder2
	//
	//builder3 := builder1
	////builder3.Grow(1)  // 赋值
	//_ = builder3

	f2 := func(bp *strings.Builder) {
		(*bp).Grow(1)   //不是并发安全
		builder4 := *bp
		//builder4.Grow(1)  // panic
		_ = builder4
	}
	f2(&builder1)

	builder1.Reset()
	builder5 := builder1
	builder5.Grow(1)


}