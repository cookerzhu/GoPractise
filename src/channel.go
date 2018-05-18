package main

import (
	"fmt"
)

func main(){
	ch := make(chan int)
	v := <- ch  //先从channel取，就会deadlock，如果先往channel放（和下面一行代码交换，就ok）
	go put(ch)
	fmt.Print("get value :" ,v)
}

func put(ch chan int){
	ch <- 1
	fmt.Println("put value : v = 1 " )

}