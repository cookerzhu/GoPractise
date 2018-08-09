package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const M = 26


type Ss struct {
	v int
}

func main() {
	const GOMAXPROCS = 1
	runtime.GOMAXPROCS(GOMAXPROCS)

	var wg sync.WaitGroup
	wg.Add(2 * M)
	for i := 0; i < M; i++ {
		go func(i int) {
			time.Sleep(time.Second * 5 )
			defer wg.Done()
			runtime.Gosched()// A
			fmt.Printf("%c", 'a'+i)
		}(i)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("%c", 'A'+i)
		}(i)
	}
	fmt.Println("main proc begin")
	wg.Wait()

	fmt.Println("\nmain proc finished")



	//s := []Ss{{1}, {3}, {5}, {2}}
	//sort.Slice(s, func(i, j int) bool { return s[i].v < s[j].v })
	//fmt.Printf("%#v", s)
}

