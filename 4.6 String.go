package main

import (
	"fmt"
	"strings"
	//"unicode/utf8"
)

func main() {
	str1 := "asSASA ddd dsjkdsjs dk"
	//fmt.Print(len(str1))
	//fmt.Printf("The number of characters in string str1 is %d\n", utf8.RuneCountInString(str1))
	//str2 := "asSASA ddd dsjkdsjsこん dk"
	//fmt.Printf("The number of bytes in string str2 is %d\n", len(str2))
	//fmt.Printf("The number of characters in string str2 is %d", utf8.RuneCountInString(str2))
	a := strings.Index(str1, "s")
	fmt.Println(a)
}
