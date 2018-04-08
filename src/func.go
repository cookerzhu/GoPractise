package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := plusX(2)//a is a func returned by plusX
	fmt.Println(a(6))
	fmt.Println(reflect.TypeOf(a) )

}

//闭包
//返回一个函数fun(y int) int {x + y }
func plusX(x int) func(int)  int{
	return func(y int) int {
		return x + y
	}
}
