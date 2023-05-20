package main

import "fmt"

func main() {
	var i1 = 5
	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)
	var intP *int = &i1

	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
	s := "How dare you"
	var p *string = &s
	*p = "What's up"
	fmt.Printf("Here is the pointer p: %p\n", p)
	fmt.Printf("Here is the string *p: %s\n", *p)
	fmt.Printf("Here is the string s: %s\n", s)
	//对一个空指针的反向引用是不合法的，并且会使程序崩溃
	var q *int = nil
	*q = 0

}
