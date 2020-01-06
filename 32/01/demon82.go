package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	comment := "Package io provides basic interfaces to I/O primitives. " +
		"Its primary job is to wrap existing implementations of such primitives, " +
		"such as those in package os, " +
		"into shared public interfaces that abstract the functionality, " +
		"plus some other related primitives."

	//fmt.Println("New a string reader and name it \"reader1\" ...")
	reader1 := strings.NewReader(comment)
	buf1 := make([]byte, 7)
	n, err := reader1.Read(buf1)
	//var offset1, index1 int64
	//executeIfNoErr(err, func() {
	//	//fmt.Printf("Read(%d): %q\n", buf1[:n])
	//	fmt.Printf("Read(%d): %q\n", n, buf1[:n])
	//	offset1 = int64(53)
	//	index1, err = reader1.Seek(offset1, io.SeekCurrent)
	//})
	//executeIfNoErr(err, func() {
	//	fmt.Printf("The new index after seeking from current with offset %d: %d\n",
	//		offset1, index1)
	//	n, err = reader1.Read(buf1)
	//})
	//executeIfNoErr(err, func() {
	//	fmt.Printf("Read(%d): %q\n", n, buf1[:n])
	//	fmt.Println("", )
	//
	//})

	//reader1.Reset(comment)
	//num1 := int64(7)
	//fmt.Printf("New a limited reader with reader1 and number %d ...\n", num1)
	//reader2 := io.LimitReader(reader1,7)
	//buf2 := make([]byte,10)
	//for i := 0; i < 3; i++ {
	//	n, err = reader2.Read(buf2)
	//	executeIfNoErr(err, func() {
	//		fmt.Printf("Read(%d): %q\n",n, buf2[:n])
	//	})
	//	fmt.Println("",)
	//}

	//reader1.Reset(comment)
	//offset2 := int64(56)
	//num2 := int64(72)
	//fmt.Printf("New a section reader with reader1, offset %d and number %d ...\n", offset2, num2)
	//reader3 := io.NewSectionReader(reader1, offset2,num2)
	//buf3 := make([]byte, 20)
	//for i := 0; i < 5; i++ {
	//  n, err = reader3.Read(buf3)
	//  executeIfNoErr(err, func() {
	//	  fmt.Printf("Read(%d): %q\n",n, buf3[:n])
	//  })
	//}
	//fmt.Println("",)

	reader1.Reset(comment)
	writer1 := new(strings.Builder)
	fmt.Println("New a tee reader with reader1 and write1 ...",)
	reader4 := io.TeeReader(reader1, writer1)
	buf4 := make([]byte, 40)
	for i := 0; i < 8; i++ {
		n, err = reader4.Read(buf4)
	  executeIfNoErr(err, func() {
		  fmt.Printf("Read(%d): %q\n",n, buf4[:n])
	  })
	}
	fmt.Println("",)
}
func executeIfNoErr(err error, f func()) {
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
		f()
	}

