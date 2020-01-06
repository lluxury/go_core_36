package main

import (
	"io"
	"strings"
)

func main() {
	comment := "Because these interfaces and primitives wrap lower-level operations with various implementations, " +
		"unless otherwise informed clients should not assume they are safe for parallel execution."
	basicReader := strings.NewReader(comment)
	//basicWriter := new(strings.Builder)

	reader1 := io.LimitReader(basicReader, 90)
	_ = interface{}(reader1).(io.Reader)


}