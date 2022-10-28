package main

import (
	"fmt"
	"strings"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for ix := 0; ix < len(str); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str[ix])
	}
	str2 := "汉语"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for ix := 0; ix < len(str2); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str2[ix])
	}
	for j := 0; j < 15; j++ {
		fmt.Println(j)
	}
	for k := 1; k <= 25; k++ {
		fmt.Println(strings.Repeat("G", k))
	}
}
