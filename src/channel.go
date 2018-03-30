package main

import (
	"fmt"
)

func main(){
	ch := make(chan int)
	go put(ch)
	v := <- ch  //先从channel取，就会阻塞，如果先放（和下面一行代码交换，就ok）
	fmt.Print(v)
}

func put(ch chan int){
	ch <- 1
	fmt.Println("put value : v = 1 " )

}