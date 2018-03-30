package main

import "fmt"


type people struct{
	Name string
}

func (p *people) String() string{
	return fmt.Sprintf("print:%v",p)
}

func main(){
	p := & people{}
	p.String() //fatal error: stack overflow  ???
}