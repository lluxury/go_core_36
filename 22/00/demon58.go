package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"sync"
)

var protecting uint

func init() {
	flag.UintVar(&protecting, "protecting", 1,
		"It indicates whether to use a mutex to protect data writing.")
}

func main() {
	flag.Parse()
	var buffer bytes.Buffer

	const (
		max1 = 5  // 启用的goroutine的数量
		max2 = 10 // 每个goroutine需要写入的数据块的数量
		max3 = 10 // 每个数据块中需要有多少个重复的数字
	)

	var mu sync.Mutex // 互斥锁

	sign := make(chan struct{}, max1)
	//for i := 1; i < max1; i++ {    // 并发程序的故障率急剧上升
	for i := 1; i <= max1; i++ {	// 三层循环
		go func(id int, writer io.Writer) {
			defer func() {
				sign <- struct{}{}  // 依次写入通道
			}()
			//for j := 1; j < max2; j++ {
			for j := 1; j <= max2; j++ {  // 准备数据
				header := fmt.Sprintf("\n[id: %d, iteration: %d]",
					id, j)
				data := fmt.Sprintf(" %d", id*j)
				if protecting > 0 {  // 写入数据
					mu.Lock()
				}
				_, err := writer.Write([]byte(header))
				if err != nil {
					log.Printf("error: %s [%d]", err, id)
				}

				for k := 0; k < max3; k++ {
					_, err := writer.Write([]byte(data))
					if err != nil {
						log.Printf("error: %s [%d]", err, id)
					}
				}
				if protecting > 0 {
					mu.Unlock()
				}
			}
		}(i, &buffer)
	}  // 并发调 goroutine
	for i := 0; i < max1; i++ {   // 依次返回，对应所引值
		<-sign
	}
	data, err := ioutil.ReadAll(&buffer)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	log.Printf("The contents:\n%s", data)
}

//初始化，变量，常量，锁，信号，遍历，便利，读data