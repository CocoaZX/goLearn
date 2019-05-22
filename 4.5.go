package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}

/* := go语言中可以使用“:=”为一个新的变量完成声明以及初始化的工作，如下例所示：

i := 1
等价于：

var i = 1
*/

/*go 也没有字符串类型，相当于C*/
