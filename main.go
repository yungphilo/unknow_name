package main

import (
	"fmt"
	teststatus "unknow_name/whatever"
)

func main() {
	fmt.Println(teststatus.TestStatus())
	/*teststatus whatever目录中go文件中的teststatus包名
	TestStatus 是teststatus包中的TestStatus函数*/
	fmt.Println(teststatus.Tt())
	fmt.Println(teststatus.Hello())

}
