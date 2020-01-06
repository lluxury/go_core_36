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

	reader2:= io.NewSectionReader(basicReader,90,89)
	_ = interface{}(reader2).(io.Reader)
	_ = interface{}(reader2).(io.ReaderAt)
	_ = interface{}(reader2).(io.Seeker)


}