package main

import "fmt"

const N = 3
func main() {
	m := make(map[int]*int)

	for i := 0; i < N; i++ {
		m[i] = &i //A m[i]  pointer to i (&i),i increase in for loop , i==3
	}

	for k, v := range m {
		fmt.Println(k,v)
	}
}
