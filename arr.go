package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var num []int
	fmt.Println("请输入求数组：")
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		l := len(input.Text())
		for i := 0; i < l; i++ {
			num[i], _ = strconv.Atoi(strings.Split(input.Text(), " ")[i])

		}
		fmt.Print(num)
		/*		b, _ = strconv.Atoi(strings.Split(input.Text(), " ")[1])
				fmt.Println(a+b)
			}*/
	}
}
