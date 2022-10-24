package main

import (
	"fmt"
	"strings"
)

func main() {
	var manyG = "gggggggggg"
	var str string = "This is an example of a string"
	var origS = "I am Groot! "
	//var newS string
	fmt.Printf("T/F? Does the string \"%s\" have suffix %s? ", str, "'ing'")
	fmt.Printf("%t\n", strings.HasSuffix(str, "ing")) //字符串是否以ing结尾
	fmt.Printf("T/F? Does the string \"%s\" have presuffix %s? ", str, "'Th'")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th")) //字符串是否以Th开头
	fmt.Printf("T/F? Does the string \"%s\" have Contains  %s? ", str, "'A'")
	fmt.Printf("%t\n", strings.Contains(str, "A")) //字符串是否包含A，区分大小写
	fmt.Printf("T/F? Does the string \"%s\" have Contains  %s? ", str, "'ne'")
	fmt.Printf("%t\n", strings.Contains(str, "ne"))
	fmt.Printf("The position of \"an\" is: ")
	fmt.Printf("%v\n", strings.Index(str, "an"))
	fmt.Printf("The position of \"Th\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Th"))
	fmt.Printf("The position of \"ing\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "ing"))
	fmt.Printf("The position of \"a\" is: ")
	//查找字符串中某个字符串多次出现时，打印第一次出现位置；不包含输出-1
	fmt.Printf("%d\n", strings.Index(str, "a"))
	fmt.Printf("%s will become ", str)
	//返回将s中前n个不重叠a串都替换为new的新字符串，如果n<0会替换所有old子串,n=0不会替换。
	fmt.Printf("%s\n", strings.Replace(str, "a", "a simple ", 1))
	fmt.Printf("Number of H's in %s is: ", str)
	//统计"H's"在句中出现次数
	fmt.Printf("%d\n", strings.Count(str, "H's"))
	fmt.Printf("Number of double g's in %s is: ", manyG)
	fmt.Printf("%d\n", strings.Count(manyG, "gg"))
	//重复字符串并返回
	newS := strings.Repeat(origS, 3)
	fmt.Printf("The new repeated string is: %s\n", newS)
	a := 10
	b := 20
	c := 30
	//Println打印的每一项之间都会有空格，并且自带换行
	fmt.Println("a=", a, ",b=", b, ",c=", c)
	//Printf 中的占位符与后面的数字变量一一对应,没有自动换行
	//更多的占位符参考：http://docscn.studygolang.com/pkg/fmt/
	fmt.Printf("a=%d,b=%d,c=%d\n", a, b, c)
	//Print 没有自动换行，也没空格
	fmt.Print("a=", a, ",b=", b, ",c=", c, "\n")
	lower := strings.ToLower(origS)
	upper := strings.ToUpper(origS)
	//大小写替换
	fmt.Printf("The lowercase is: %s\n", lower)
	fmt.Printf("The uppercase is: %s\n", upper)
	//修剪字符
	space := " It's my precious! "
	//去除头尾空字符
	nospace := strings.TrimSpace(space)
	fmt.Printf("%s!Cut space is: \n%s!\n", space, nospace)
	//去除头尾指定字符串
	noit := strings.Trim(space, " ")
	fmt.Printf("%s cut \" \" is: \n%s\n", space, noit)
	//行中去掉字符串可以用替换方法实现
	noc := strings.Replace(space, "'", "", 1)
	fmt.Printf("%s cut \"'\" is: %s\n", space, noc)
	//去除头的字符
	nohead := strings.TrimLeft(nospace, "I")
	fmt.Printf("No head is: %s\n", nohead)
	//去尾字符
	notail := strings.TrimRight(nospace, "!")
	fmt.Printf("No tail is: %s\n", notail)
	tt := strings.Fields(nospace)
	l := len(tt)
	slicestr := strings.Split(nospace, " ")
	l2 := len(slicestr)
	fmt.Printf("%s slice is: %s length is %d\n", nospace, tt, l)
	fmt.Printf("%s slice is: %s length is %d\n", nospace, slicestr, l2)
	for _, va1 := range slicestr {
		fmt.Printf("%s - ", va1)
	}
	fmt.Println()
	groot := "I am Groot,i am Groot,i am Groot"
	groot2 := strings.Split(groot, ",")
	for _, va2 := range groot2 {
		fmt.Printf("%s - ", va2)
	}
	fmt.Println()
	va3 := strings.Join(groot2, "!")
	fmt.Printf("%s joined by !: %s\n", groot2, va3)
	coding := "每个中文汉字到底占多少byte"
	//11*3+4*1 共占37个字节
	//中文汉字UTF-8编码占用字节数
	fmt.Println(len(coding))
	/*rune
	字符串由字符组成，字符的底层由字节组成，
	而一个字符串在底层的表示是一个字节序列。在 Go 语言中，
	字符可以被分成两种类型处理：对占 1 个字节的英文类字符，
	可以使用 byte（或者 unit8 ）；对占 1 ~ 4 个字节的其他字符，
	可以使用 rune（或者 int32 ），如中文、特殊符号等。
	*/
	//字数
	fmt.Println(len([]rune(coding)))
	fmt.Println([]byte(coding))
	fmt.Println([]rune(coding))

}

/* 输出
[Running] go run "d:\Docments\work\code\unknow_name\demo1\presuffix.go"
T/F? Does the string "This is an example of a string" have suffix 'ing'? true
T/F? Does the string "This is an example of a string" have presuffix 'Th'? true
T/F? Does the string "This is an example of a string" have Contains  'A'? false
T/F? Does the string "This is an example of a string" have Contains  'ne'? false
The position of "an" is: 8
The position of "Th" is: 0
The position of "ing" is: 27
The position of "a" is: 8
This is an example of a string will become This is a simple n example of a string
Number of H's in This is an example of a string is: 0
Number of double g's in gggggggggg is: 5
The new repeated string is: I am Groot! I am Groot! I am Groot!
a= 10 ,b= 20 ,c= 30
a=10,b=20,c=30
a=10,b=20,c=30
*/
