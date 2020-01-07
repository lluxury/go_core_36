package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	var ioTypes = []reflect.Type{
		reflect.TypeOf((*io.Reader)(nil)).Elem(),
		reflect.TypeOf((*io.Writer)(nil)).Elem(),
		reflect.TypeOf((*io.Closer)(nil)).Elem(),

		reflect.TypeOf((*io.ByteReader)(nil)).Elem(),
		reflect.TypeOf((*io.RuneReader)(nil)).Elem(),
		reflect.TypeOf((*io.ReaderAt)(nil)).Elem(),
		reflect.TypeOf((*io.Seeker)(nil)).Elem(),
		reflect.TypeOf((*io.WriterTo)(nil)).Elem(),
		reflect.TypeOf((*io.ByteWriter)(nil)).Elem(),
		reflect.TypeOf((*io.WriterAt)(nil)).Elem(),
		reflect.TypeOf((*io.ReaderFrom)(nil)).Elem(),

		reflect.TypeOf((*io.ByteScanner)(nil)).Elem(),
		reflect.TypeOf((*io.RuneScanner)(nil)).Elem(),
		reflect.TypeOf((*io.ReadSeeker)(nil)).Elem(),
		reflect.TypeOf((*io.ReadCloser)(nil)).Elem(),
		reflect.TypeOf((*io.WriteCloser)(nil)).Elem(),
		reflect.TypeOf((*io.WriteSeeker)(nil)).Elem(),
		reflect.TypeOf((*io.ReadWriter)(nil)).Elem(),
		reflect.TypeOf((*io.ReadWriteCloser)(nil)).Elem(),
		reflect.TypeOf((*io.ReadWriteSeeker)(nil)).Elem(),
	}

	file1 := (*os.File)(nil)
	fileType := reflect.TypeOf(file1)
	var buf bytes.Buffer
	fmt.Fprintf(&buf,"Type %T implements\n",file1)
	for _, t := range ioTypes {
		if fileType.Implements(t){
			buf.WriteString(t.String())
			buf.WriteByte(',')
			buf.WriteByte('\n')
			//t是上面类型的遍历，如果实现就打出来
		}
	}
	output := buf.Bytes()
	output[len(output)-2] = '.'
	fmt.Printf("%s\n",output)



}