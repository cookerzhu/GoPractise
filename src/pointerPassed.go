package main 

import (
	"fmt"
)

func main() {
	//x := 4 ;
	//fmt.Println(&x)
	//x1 := add(&x)//传入内存地址，即指针。
	//
	//fmt.Println(x1)//6
	//fmt.Println(&x1)

	var a Integer = 1
	a.Add(2)
	fmt.Println("a=",a)
	fmt.Println("&a=",&a)

	
}

func add(a *int) int{//a 表示指针 *int 类型,前面加*表示该指针对应的内存里面存储的值
	b := 10
	a = &b
	c := *a
	fmt.Println(c)
	
	*a = *a + 2
	return *a 
}

type Integer int

func (a Integer) Add(b Integer){
	a += b  // a == 1
	fmt.Println(a)
	fmt.Println(&a) //此处的a 只是一个副本，和main函数中的a指向不同的内存地址
}

//func main(){
//	var a Integer = 1;
//	a.add(2)
//	fmt.Println("a=",a) // a == 3
//}


