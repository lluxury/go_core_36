package main

import "fmt"

func main(){
	numbers2 := [...]int{1, 2, 3, 4, 5, 6}
	numbers3 := []int{1, 2, 3, 4, 5, 6}
	maxIndex3 := len(numbers2) -1
	for i, e := range numbers3 {
		if i == maxIndex3 {
		numbers3[0] += e
		} else {
			numbers3[i+1] += e
		}
	}
	fmt.Println(numbers3)
	}