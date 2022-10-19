// package main

// import (
// 	"fmt"
// 	"os"
// 	"runtime"
// )

// func main() {
// 	var goos string = runtime.GOOS
// 	fmt.Printf("The operating system is: %s\n", goos)
// 	path := os.Getenv("PATH")
// 	fmt.Printf("Path is %s\n", path)
// }
package main

var a = "G"

func main() {
	n()
	l()
	m()
	n()
}

func n() {
	print(a)
}

func m() {

	a = "O" //对变量"a"重新赋值，而"a"是全局变量，所以再次执行函数n时，打印为"O";
	//a : = "O" //若"a : = "O";则是在这个函数体中初始化了一个和全局变量"a"相同名字的一个局部变量,所以再次执行函数n时,打印的还为为初始值"G"
	//所以为了可读性还是不要把局部变量和全局变量起同一个名字
	print(a)
}
func l() {
	print(a)
}
