package main

import ("fmt")

//常量使用const定义
func main()  {
	const a int=100
	fmt.Printf("a=%d\n",a)
	const(
		b int = 88
		c	//c的值和上一行b一样
	)
	fmt.Printf("b=%d c=%d\n",b,c)

	//iota每一行都会递增,第一行从0开始,iota默认值为0，递增只有在const作用域有效
	/*
	const(
		e = iota
		f
		g
	)
	*/
	const(
		e = 1<<iota
		f
		g
	)
	fmt.Printf("e=%d f=%d g=%d\n",e,f,g)
}