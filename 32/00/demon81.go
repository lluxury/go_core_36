package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	builder := new(strings.Builder)
	_ = interface{}(builder).(io.Writer)
	_ = interface{}(builder).(io.ByteWriter)
	_ = interface{}(builder).(fmt.Stringer)
}

