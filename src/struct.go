package main 

import (
	f "fmt"
)

type Skills []string

type Human struct{
	name string
	age int
	weight int 
}

type Student struct{
	Human //匿名字段。
	Skills
	spec string
	weight int 
}

//作为Student类型的内置method,接收者是Student类型。
func (t Student) sayHi(){
	f.Println("my name is" + t.name)
}

//使用 field:value的方式 初始化，可以任意顺序，也可只初始化部分字段
func main() {
	mark := Student{Human:Human{"Mark",20,120},spec:"Computer Science"}//初始化部分字段
//	peter := Student{Human{weight:130,name:"Peter",age:30},"Computer Science"}//初始化
	peter := Student{Human:Human{weight:130,name:"Peter",age:30},spec:"Computer Science"}//部分初始化
	
	john := Student{Human{"John",10,90},Skills{"ABC","BCD"},"BIO",99}//完整定义每一个字段，按照顺序
	
	john.Skills = append(john.Skills,"CDE")//添加一个到数组中
	f.Println(john.Skills)
	f.Println(john.Human.weight)//90,访问的是Human里面的weight.
	f.Println(john.weight)//99,访问的是Student里面的weight.
	
	f.Println("name is", mark.name)
	f.Println("age is", mark.age)
	f.Println("weight is", peter.weight)//为0,因为没有初始化
	
}

