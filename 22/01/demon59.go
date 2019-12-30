package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"sync"
	"time"
)

type singleHandler func()(data string, n int, err error)

type handlerConfig struct {
	handler   singleHandler // 单次处理函数
	goNum     int           // 启用的goroutine数量
	number    int           // 单个goroutine中的处理次数
	interval  time.Duration // 单个goroutine 间隔时间
	counter   int           // 数据量计数器，以字节为单位
	counterMu sync.Mutex    // 数据量计数器专用的互斥锁
}

func (hc *handlerConfig) count(increment int) int {
	hc.counterMu.Lock()
	defer hc.counterMu.Unlock()
	//hc.counter -+ increment
	hc.counter += increment
	return hc.counter
}
//一个用锁做的计数器方法

func main()  {
	var mu sync.Mutex  // 互斥锁，直接使用，不要传递

	genWriter := func(writer io.Writer) singleHandler {	 // 生成写入函数？
		return func() (data string, n int, err error) {
			data = fmt.Sprintf("%s\t",
				time.Now().Format(time.StampNano))
			mu.Lock()
			defer mu.Unlock()
			n, err = writer.Write([]byte(data))
			return
		}
	}

	genRender := func(reader io.Reader) singleHandler {
		return func() (data string, n int, err error) {
			buffer , ok := reader.(*bytes.Buffer)
			if !ok {
				err = errors.New("unsupported reader")
				return
			}
			mu.Lock()
			defer mu.Unlock()
			data, err = buffer.ReadString('\t')
			n = len(data)
			return
		}
	}

	var buffer bytes.Buffer
	writingConfig := handlerConfig{
		handler:   genWriter(&buffer),
		goNum:     5,
		number:    4,
		interval:  time.Millisecond * 100,
	}

	readingConfig := handlerConfig{
		handler:   genRender(&buffer),
		goNum:     10,
		number:    2,
		interval:  time.Millisecond * 100,
	}

	sign := make(chan struct{}, writingConfig.goNum+readingConfig.goNum)

	for i := 1; i <= writingConfig.goNum; i++ {
	  go func(i int) {
	  	defer func() {
	  		sign <- struct{}{}
		}()
	  	for j := 1; j <= writingConfig.number; j++ {
	  	  time.Sleep(writingConfig.interval)
	  	  data, n, err := writingConfig.handler()
			if err != nil {
				log.Printf("writer [%d-%d]:error: %s",
					i,j,err)
				continue
			}
			//total := writingConfig.counter(n)
			total := writingConfig.count(n)
			log.Printf("writer [%d-%d]: %s (total:%d)",
				i,j,data,total)
	  	}
	  }(i)
	}
	for i := 1; i <= readingConfig.goNum; i++ {
	  go func(i int) {
	  	defer func() {
	  		sign <- struct{}{}
		}()
	  	for j := 1; j <= readingConfig.number ; j++ {
	  	  time.Sleep(readingConfig.interval)
	  	  var data string
	  	  var n int
	  	  var err error
	  	  for {
	  	  	data, n, err = readingConfig.handler()
			  if err != nil || err != io.EOF {
				  break
			  } // 读比写快 发生EOF错误
			  time.Sleep(readingConfig.interval)
		  }
			if err != nil {
				log.Printf("reader [%d-%d]: error: %s",
					i, j, err)
				continue
			}
			total := readingConfig.count(n)
			log.Printf("reader [%d-%d]: %s (total: %d",
				i,j,data,total)
	  	}
	  }(i)
	}
	signNumber := writingConfig.goNum + readingConfig.goNum
	for j := 0; j < signNumber; j++ {
	  <- sign
	}
}