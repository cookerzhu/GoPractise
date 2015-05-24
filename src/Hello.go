package main 

import "fmt"
//import (
//	"fmt"
//)

var AMap map[rune]int32
//全局变量，只是声明,rune 就是 int32的别称,go数据类型中有int(int 与int32 不可互用),但是没有 float

func main() {
//	s := "hello\t"
	//Raw 字符串，没有字符转义，原样输出
//	m := `hello
//	world`
//	fmt.Printf(m)

//	array and slice类型
	array  := [6]byte{'a','b','c','d','e','f'}
 	aslice := array[:4]//包含'a','b','c',aslice[1] is 'a'
 	
 	aslice = append(aslice,'s')//添加一个元素到slice
	
//	rating := map[string]int32 {"a":5,"b":4,"c":3}
//	c ,ok := rating["d"]//有2个返回值,key "d" 不存在
//	fmt.Println(c)//0
//	fmt.Println(ok)//false
	
//	for index := range aslice {
//		fmt.Println(aslice[index])//打印'a' 等值的 ascii码 数值
//	}
	
//	for index := range rating {
//		fmt.Println(index)//打印 map 里面的元素 ，没有固定顺序。此时打印的是key,a,b,c
//	}
	
	for k,v := range aslice {
//		fmt.Println("map's key is " + k);使用_舍弃不想要的返回值
		fmt.Print(k)//遍历，k是 index索引,v是 值
		fmt.Print("\t")
		fmt.Println(v)
	}
	
//	AMap := make(map[int32]int32)//使用前用make初始化，初始化map内部的数据结构，填充适当的值
//	AMap[1] = 23;
	
}


