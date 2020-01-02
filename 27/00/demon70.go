package main

import (
	"bytes"
	"fmt"
	"io"
	"sync"
)

var bufPool sync.Pool

type Buffer interface {
	Delimiter() byte
	Write(contents string) (err error)
	Read()(contents string, err error)
	Free()
}

type myBuffer struct {
	buf bytes.Buffer
	delimiter byte
}

func (b *myBuffer) Delimiter() byte {
	return b.delimiter
}

func (b *myBuffer) Write(contents string) (err error){
	if _, err = b.buf.WriteString(contents); err != nil {
		return
	}
	return b.buf.WriteByte(b.delimiter)
}

func (b *myBuffer) Read() (contents string, err error) {
	return b.buf.ReadString(b.delimiter)
}

func (b *myBuffer) Free() {
	bufPool.Put(b)
}

var delimiter = byte('\n')

func init() {
	bufPool = sync.Pool{	// 初始化 零时变量，赋给 bufPool
		New: func() interface{} {
			return &myBuffer{delimiter:delimiter}
		},
	}
}

func GetBuffer() Buffer {
	return bufPool.Get().(Buffer)
}

func main() {
	buf := GetBuffer()
	defer buf.Free()
	buf.Write("A Pool is a set of temporary objects that" +
		"may be individually saved and retrieved.")
	buf.Write("A Pool is safe for use by multiple goroutines simultaneously.")
	buf.Write("A Pool must not be copied after first use.")

	fmt.Println("The data blocks in buffer:",)
	for {
		block, err := buf.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(fmt.Errorf("unexpected error: %s", err))
		}
		fmt.Print(block)
	}
}

