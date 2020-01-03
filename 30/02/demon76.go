package main

import "strings"

func main() {
	var builder1 strings.Builder
	builder1.Grow(1)

	f1 := func(b strings.Builder) {
		//b.Grow(1)   // 函数
	}
	f1(builder1)

	ch1 := make(chan strings.Builder, 1)
	ch1 <- builder1
	builder2 := <- ch1
	//builder2.Grow(1)	// 通道
	_ = builder2

	builder3 := builder1
	builder3.Grow(1)  // 赋值
	_ = builder3
}