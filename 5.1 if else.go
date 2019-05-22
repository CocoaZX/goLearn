package main

import (
	"fmt"
	"runtime"
)

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { //Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
	fmt.Print(prompt)
}

func main() {
	var first int = 10
	var cond int

	if first <= 0 {

		fmt.Printf("first is less than or equal to 0\n")
	} else if first > 0 && first < 5 {

		fmt.Printf("first is between 0 and 5\n")
	} else {

		fmt.Printf("first is 5 or greater\n")
	}
	if cond = 5; cond > 10 {

		fmt.Printf("cond is greater than 10\n")
	} else {

		fmt.Printf("cond is not greater than 10\n")
	}
}

var prompt = "Enter a digit, e.g. 3 " + "or %s to quit."

/*注意事项 不要同时在 if-else 结构的两个分支里都使用 return 语句，这将导致编译报错 function ends without a return statement（你可以认为这是一个编译器的 Bug 或者特性）。（ 译者注：该问题已经在 Go 1.1 中被修复或者说改进 ）*/
