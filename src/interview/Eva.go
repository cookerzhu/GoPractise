//order of evaluation of variables in return statement is not determined

package main

import "fmt"

func main() {
	m, i, j := testChange()
	fmt.Println("retun ", m, i, j) //2 0 99
}

func testChange() (int, interface{}, int) {
	m := 0
	i := 0

	f := func() int {
		m = 2
		i = 2
		return 99
	}

	return m, i, f()
}