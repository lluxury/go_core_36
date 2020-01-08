package main

import (
	"fmt"
	"net"
	"time"
)

type dailArgs struct {
	network	string
	address string
	timeout time.Duration
}

func main() {
	dialArgsList := []dailArgs{
		{
			"tcp",
			"google.cn:80",
			time.Millisecond * 500,
		},
		{
			"tcp",
			"google.cn:80",
			time.Millisecond * 2,
			// "connect: operation timed out"，
			// 错误表示底层的socket在连接网络服务的时候先超时
		},
		{
			"tcp",
			"google.cn:80",
			//time.Millisecond * 4,
			time.Millisecond * 25,
		},
	}

	for _, args := range dialArgsList {
		fmt.Printf("Dial %q with network %q and timeout %s ...\n",
			args.address, args.network, args.timeout)
	ts1 := time.Now()
	conn, err := net.DialTimeout(args.network,args.address, args.timeout)
	ts2 := time.Now()
	fmt.Printf("Elapsed time: %s\n",time.Duration(ts2.Sub(ts1)) )
		if err != nil {
			fmt.Printf("dial error: %v\n", err)
			fmt.Println("",)
			continue
		}
		defer conn.Close()
	fmt.Printf("The local address: %s\n", conn.LocalAddr())
	fmt.Printf("The remote address: %s\n", conn.RemoteAddr())
	fmt.Println("",)


	}
}