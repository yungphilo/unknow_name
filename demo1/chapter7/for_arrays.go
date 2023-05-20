package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	var arr1 [5]int

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}

	for i := 0; i < len(arr1); i++ {
		fmt.Printf("Array at index %d is %d\n", i, arr1[i])
	}
	// for i, _ := range arr1 {
	// 	fmt.Printf("Array at index %d is %d\n", i, arr1[i])
	// }
	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("The progress took this amount of time: %s\n", delta)
}
