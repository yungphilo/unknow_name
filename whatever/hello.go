package teststatus

//第一行且一个文件夹下的所有文件必须使用同一个包名，一个.go文件不能声明两个package
//函数名首字母必须大写，go约定大写开头的标识符才能被import用
func Hello() string {
	return "Hello func"
}
