package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	comment := "Package bufio implements buffered I/O. " +
		"It wraps an io.Reader or io.Writer object, " +
		"creating another object (Reader or Writer) that " +
		"also implements the interface but provides buffering and " +
		"some help for textual I/O."
	basicReader := strings.NewReader(comment)
	fmt.Printf("The size of basic reader: %d\n",basicReader.Size())

	//size := 300
	//fmt.Printf("New a buffered reader with size %d ...\n",size)
	//reader1 := bufio.NewReaderSize(basicReader,size)
	//fmt.Println("",)
	//
	//fmt.Print("[ About 'Peek' method]\n\n")
	//peekNum := 38
	//fmt.Printf("Peek %d bytes ...\n",peekNum)
	//bytes , err := reader1.Peek(peekNum)
	//if err != nil {
	//	fmt.Printf("error: %v\n",err)
	//}
	//fmt.Printf("Peeked contents(%d): %q\n", len(bytes),bytes)
	//fmt.Printf("The number of unread bytes in the buffer: %d\n",reader1.Buffered())
	//fmt.Println("",)
	//fmt.Print("[ About 'Read' method ]\n\n")

	//readNum := 38
	//buf1 := make([]byte,readNum)  // 准备切片
	//fmt.Printf("Read %d bytes ...\n",readNum) // 读数
	//n, err := reader1.Read(buf1)
	//if err != nil {
	//	fmt.Printf("error: %\n", err)
	//}
	//fmt.Printf("Read contents(%d): %q\n",n,buf1)  // 内容
	//fmt.Printf("The number of unread bytes in the buffer : %d\n",reader1.Buffered()) // 读计数
	//fmt.Println("",)

	//fmt.Print("[ About 'ReadSlice' method ]\n\n")
	//
	//fmt.Println("Reset the basic reader ...")
	//basicReader.Reset(comment)
	//fmt.Println("Reset the buffered reader ...")
	//reader1.Reset(basicReader)
	////重置读取器和缓冲
	//fmt.Println("", )
	//
	//delimiter := byte('(')
	//fmt.Printf("Read slice with delimiter %q...\n", delimiter)
	//line, err := reader1.ReadSlice(delimiter)
	//if err != nil {
	//	fmt.Printf("error: %v\n", err)
	//}
	//fmt.Printf("Read contents(%d) %q\n", len(line),line)
	//fmt.Printf("The number of unread bytes in the buffer: %d\n",reader1.Buffered())
	//fmt.Println("",)
	//
	//delimiter = byte('[')
	//fmt.Printf("Read slice with delimiter %q...\n", delimiter)
	//line, err = reader1.ReadSlice(delimiter)
	//if err != nil {
	//	fmt.Printf("error: %v\n", err)
	//}
	//fmt.Printf("Read contents(%d) %q\n", len(line),line)
	//fmt.Printf("The number of unread bytes in the buffer: %d\n",reader1.Buffered())
	//fmt.Println("",)
	//// 找不到，把剩余部分读了，报了个错
	//
	//fmt.Println("Reset the basic reader ...")
	//basicReader.Reset(comment)
	//size = 200
	//fmt.Printf("New a buffered reader with size %d ...\n",size)
	//reader2 := bufio.NewReaderSize(basicReader,size)
	//fmt.Println("",)
	//
	//delimiter = byte('[')
	//fmt.Printf("Read slice with delimiter %q...\n",delimiter)
	//line, err = reader2.ReadSlice(delimiter)
	//if err != nil {
	//	fmt.Printf("error: %v\n", err)
	//}
	////error: bufio: buffer full
	//
	//fmt.Printf("Read contents(%d): %q\n",len(line),line)
	//fmt.Printf("The number of unread bytes in the buffer: %d\n",reader2.Buffered())
	//fmt.Println("",)
	//
	//fmt.Print("[ About 'ReadBytes' method ]\n\n")
	//fmt.Println("Reset the basic reader ...")
	//basicReader.Reset(comment)
	//size = 200
	//fmt.Printf("New a buffered reader with size %d ...\n", size)
	//reader3 := bufio.NewReaderSize(basicReader,size)
	//fmt.Println("",)
	//
	//delimiter = byte('[')
	//fmt.Printf("Read bytes with delimiter %q...\n", delimiter)
	//line, err = reader3.ReadBytes(delimiter)
	//if err != nil {
	//	fmt.Printf("error: %v\n", err)
	//}
	////error: EOF
	//fmt.Printf("Read contents(%d): %q\n",len(line),line)
	//fmt.Printf("The number of unread bytes in the buffer: %d\n",reader3.Buffered())
	//fmt.Println("",)
	//
	//fmt.Print("[ About contents leak ]\n\n")
	showContentsLeak(comment)

}

func showContentsLeak(comment string)  {
	basicReader := strings.NewReader(comment)
	fmt.Printf("The size of basic reader1: %d\n", basicReader.Size())

	size := len(comment)
	fmt.Printf("New a buffered reader with size %d ...\n",size)
	reader4 := bufio.NewReaderSize(basicReader,size)
	fmt.Println("",)

	peekNum := 7
	fmt.Printf("Peek %d bytes ...\n",peekNum)
	bytes, err := reader4.Peek(peekNum)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("Peeked contents(%d): %q\n", len(bytes),bytes)
	fmt.Println("",)

	bytes = bytes[:cap(bytes)]
	fmt.Printf("The all of the contents in the buffer:\n%q\n",bytes)
	fmt.Println("",)   // 看到了未读取内容

	//blank := bytes
	blank := byte('X')
	fmt.Println("Set blanks into the contents in the buffer ...")
	for _, i := range []int{55, 56, 57, 58, 66, 67, 68} {
		bytes[i] = blank
	}
	fmt.Println("",)

	peekNum = size
	fmt.Printf("Peek %d bytes ...\n",peekNum)
	bytes, err = reader4.Peek(peekNum)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("Peeked contents(%d):\n%q\n", len(bytes), bytes)
	fmt.Println("",)  // 修改了缓存

}