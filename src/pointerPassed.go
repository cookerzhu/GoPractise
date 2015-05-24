package main 

import (
	"fmt"
)

func main() {
	x := 4 ;
	fmt.Println(&x)
	x1 := add(&x)//传入内存地址，即指针。
	
	fmt.Println(x1)//6
	fmt.Println(&x1)
	
}

func add(a *int) int{//a 表示指针 *int 类型,前面加*表示该指针对应的内存里面存储的值
	b := 10
	a = &b
	c := *a
	fmt.Println(c)
	
	*a = *a + 2
	return *a 
}


