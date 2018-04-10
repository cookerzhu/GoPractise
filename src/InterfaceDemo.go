package main

import (
	"fmt"
	"reflect"
)

type I interface{
	Put(int)
	Get() int
}

type S struct {
	i int
}

func(p *S) Get() int{
	return p.i
}

func(p *S) Put(v int){
	p.i = v
}


type R struct {
	i int
}

func(p *R) Get() int{
	return p.i
}

func(p *R) Put(v int){
	p.i = v
}

//接口类型作为参数
func foo(p I){
	fmt.Println(p.Get())
	p.Put(1)

	switch t:= p.(type){
		case *S:
		case *R:
		case S:
		case R:
		default:
			fmt.Print(t)
	}

	if t1 , ok:= p.(I);ok{
		fmt.Println(reflect.TypeOf(t1))
	}
}

/**
t := something.(I)   .(I) 是 类型
断言，用于转换 something 到 I 类型的接口。

在 switch 之外使用 (type) 是非法的。类型判断不是唯一的运行时得到类型的方法。
为了在运行时得到类型，同样可以使用 “comma, ok” 来判断一个接口类型是否实现了
某个特定接口：
i f t, ok := something.(I) ; ok {
// 对于某些实现了接口 I 的
// t 是其所拥有的类型
}
 */
