package main

import (
	"fmt"
	"sync"
)

type IntStrMap struct {
	m sync.Map
}

func (iMap *IntStrMap) Delete(key int)  {
	iMap.m.Delete(key)
}

func (iMap *IntStrMap) Load(key int) (value string, ok bool) {
	//v, ok := iMap.Load(key)
	v, ok := iMap.m.Load(key)
	if v != nil {
		value = v.(string)
	}
	return
}

func (iMap *IntStrMap) LoadOrStore(key int, value string) (actual string, loaded bool) {
	a,loaded := iMap.m.LoadOrStore(key, value)
	actual = a.(string)
	return
}

func (iMap *IntStrMap) Range(f func(key int, value string) bool) {
	f1 := func(key,value interface{}) bool {
		return f(key.(int), value.(string))
	}
	iMap.m.Range(f1)
}

func (iMap *IntStrMap) Store(key int, value string)  {
	iMap.m.Store(key, value)
}





var pairs = []struct{
	k int
	v string
}{
	{k: 1, v: "a"},
	{k: 2, v: "b"},
	{k: 3, v: "c"},
	{k: 4, v: "d"},
}



func main() {
	//var sMap sync.Map
	//sMap.Store([]int{1, 2, 3},4)
	//_ = sMap

	var iMap IntStrMap
	iMap.Store(pairs[0].k, pairs[0].v)
	iMap.Store(pairs[1].k, pairs[1].v)
	iMap.Store(pairs[2].k, pairs[2].v)
	fmt.Println("[Three pairs have been stored in the IntStrMap instance]")
	
	iMap.Range(func(key int, value string) bool {
		fmt.Printf("The result of an iteration in Range: %d, %s\n",
			key,value)
		return true
	})

	k0 := pairs[0].k
	v0, ok := iMap.Load(k0)
	fmt.Printf("The result of Load: %v, %v (key: %v)\n",
		v0, ok, k0)

	k3 := pairs[3].k
	v3, ok := iMap.Load(k3)
	fmt.Printf("The result of Load: %v, %v (key: %v)\n",
		v3, ok, k3)

	k2, v2 := pairs[2].k, pairs[2].v
	actual2, loaded2 := iMap.LoadOrStore(k2,v2)
	fmt.Printf("The result of LoadOrStore: %v, %v (key: %v, value: %v)\n",
		actual2,loaded2,k2,v2)
	v3 = pairs[3].v
	actual3, loaded3 := iMap.LoadOrStore(k3,v3)
	fmt.Printf("The result of LoadOrStore: %v, %v (key: %v, value: %v)\n",
		actual3, loaded3, k3, v3)

	k1 := pairs[1].k
	iMap.Delete(k1)
	fmt.Printf("[The pair with the key of %v has been removed from the IntStrMap instance]\n",
		k1)
	v1,ok := iMap.Load(k1)
	fmt.Printf("The result of Load: %v, %v (key: %v)\n",
		v1, ok, k1)
	v1 = pairs[1].v
	actual1, loaded1 := iMap.LoadOrStore(k1,v1)
	fmt.Printf("The result of LoadOrStore: %v, %v (key: %v, value: %v)\n",
		actual1, loaded1, k1, v1)
	
	iMap.Range(func(key int, value string) bool {
		fmt.Printf("The result of an iteration in Range: %d, %s\n",
			key,value)
		return true
	})
	fmt.Println("",)
}