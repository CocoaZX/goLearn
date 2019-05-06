package main

func main() {
	const (
		a = iota
		b = iota
		c = iota
	)
	println(a, b, c)
}

/*Go 语言不存在隐式类型转换，因此所有的转换都必须显式说明
但是初始化赋值的时候可以不写类型，类似swift*/
/*iota 也可以用在表达式中，如：iota + 50。在每遇到一个新的常量块或单个常量声明时， iota 都会重置为 0（ 简单地讲，每遇到一次 const 关键字，iota 就重置为 0 ）。*/
/*声明变量的一般形式是使用 var 关键字：var identifier type。*/
