package main

import (
	"fmt"
	"sync"
)
const N = 10

func Add(x,y int){
	z := x + y
	fmt.Println(z)
}

//Goroutine Closures
func main(){
	//for i:= 0;i<10;i++ {
	//	go Add(i,i)
	//}
	//什么都没打印，当main函数返回时，程序退出，且程序并不等待其他goroutine(非主goroutine)结束，主函数
	//启动了10个goroutine，然后返回，程序就结束了，被启动的goroutine没来得及执行

	m := make(map[int]int)

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			//fmt.Println(i)//10
			m[i] = i
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	println(len(m))
}