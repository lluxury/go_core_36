package main

import (
	"context"
	"fmt"
)

type myKey int

func main()  {
	keys := []myKey{
		myKey(20),
		myKey(30),
		myKey(60),
		myKey(61),
	}

	values := []string{
		"value in node2",
		"value in node3",
		"value in node6",
		"value in node6Branch",
	}

	rootNode := context.Background()		// 根
	node1, cancleFunc1 := context.WithCancel(rootNode)  // 撤销节点
	defer cancleFunc1()

	node2 := context.WithValue(node1,keys[0],values[0])
	node3 := context.WithValue(node2,keys[1],values[1])  // 数据节点
	fmt.Printf("The value of the key %v found in the node3: %v\n",
		//keys[0], node3.Value([keys[0]]))
		keys[0], node3.Value(keys[0]))
	fmt.Printf("The value of the key %v found in the node3: %v\n",
		keys[1], node3.Value(keys[1]))
	fmt.Printf("The value of the key %v found in the node3: %v\n",
		keys[2], node3.Value(keys[2]))  // 一跟反查

}