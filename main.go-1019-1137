package main

var a string

//输出结果为"GO" a的初始值为空，也不是空格
func main() {
	a := "G"
	print(a)
	f2()
	f1()

}

func f1() {
	a := "O"
	print(a)
	f2()
}

func f2() {
	print(a)
	//这个是全局变量，f1的局部变量不会传递给所调用的函数，变量只作用在{}内，即使是在main中声明的变量也不会传递下来
}
