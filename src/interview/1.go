package main

import "fmt"

type Param map[string]interface{}

type Show struct {
	Param
}

func main() {
	s := new(Show)
	//s.Param["RMB"] = 10000  //error map没有正确的初始化
	s.Param = map[string]interface{} {"a":1}

	s.Param = Param{"b":2}

	s.Param = make(map[string]interface{})
	s.Param["c"] = 1
	fmt.Print(s.Param["c"])
}
