package main

import (
	"fmt"
)

var A,B int32 = 1,2
var C float32 = 1.2

func main(){
//	noReturn(A,B)
//	fmt.Println(A)
	
	a,_ := add3(A,B);//使用 _ 舍弃不需要的返回值
	fmt.Println(a)
}

func add1(a,b int32)( int32, int32){//省略变量声明
	suma := a + b
	sumb := a + b + 1
	return suma,sumb
}

func add3(a,b int32)( int32, int32){//省略变量声明
	suma := a + b
	sumb := a + b + 1
	return suma,sumb
}

func add2(a,b int32) int32 { //省略变量声明，只保留类型声明,且只返回一个值，省略括号。
	return a + b
}

func noReturn(a,b int){
	a = a + b 
	fmt.Println(a)
}
