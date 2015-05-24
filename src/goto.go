package main 

import "fmt"
//import (
//	"fmt"
//)

func main() {
//	array and slice类型
	array  := [6]byte{'a','b','c','d','e','f'}
 	aslice := array[:4]//包含'a','b','c',aslice[1] is 'a'
//	rating := map[string]int32 {"a":5,"b":4,"c":3}
	flag := false
	
	Here:
	if flag {
		fmt.Println("Here :")
		goto Here1
	}
	for index := range aslice{
		
		fmt.Println(aslice[index])
		
		if(index == 0){
			flag = true
			goto Here
		}
	}
	
	Here1: 
	fmt.Println("it's over")
}


