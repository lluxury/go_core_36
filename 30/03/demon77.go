package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader1 := strings.NewReader("NewReader returns a new Reader reading from s. " +
		"It is similar to bytes.NewBufferString but more efficient and read-only.")
	fmt.Printf("The size of reader: %d\n", reader1.Size())
	fmt.Printf("The reading index in reader: %d\n",
		reader1.Size()-int64(reader1.Len()))
	// 全部 119 已读 0
	
	buf1 := make([]byte, 47)
	n, _ := reader1.Read(buf1)
	fmt.Printf("%d bytes were read. (callRead)\n",n)
	fmt.Printf("The reading index in reader: %d]\n",
		reader1.Size()-int64(reader1.Len()))
	fmt.Println("",)
	// 阅读 47， 计数 47, reader1.Size()  119， reader1.Len() 72未读数

	buf2 := make([]byte, 21)
	offset1 := int64(64)
	n, _ = reader1.ReadAt(buf2, offset1)
	fmt.Printf("%d bytes were read. (call ReadAt, offset: %d)\n",n, offset1)
	fmt.Printf("The reading index in reader: %d\n",
		reader1.Size()-int64(reader1.Len()))
	//指定阅读 提供 offset1，有n没有计数

	fmt.Println("",)
	
	offset2 :=  int64(17)
	//保留之前的值， 119，72，右向偏移
	expectedIndex := reader1.Size() - int64(reader1.Len()) + offset2
	fmt.Printf("Seek with offset %d and whence %d ...\n", offset2, io.SeekCurrent)
	readingIndex, _ := reader1.Seek(offset2, io.SeekCurrent)
	fmt.Printf("The reading index in reader: %d (returned by Seek)\n", readingIndex)
	fmt.Printf("The reading index in reader: %d (computed by me)\n", expectedIndex)

	//已经跳转到了 64
	n, _ = reader1.Read(buf2)
	fmt.Printf("%d bytes were read. (call Read)\n",n)
	fmt.Printf("The reading index in reader: %d\n",
		reader1.Size()-int64(reader1.Len()))
}