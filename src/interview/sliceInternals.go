package main

import (
"fmt"
"unsafe"
	"reflect"
)

const A = 10
const B = 5

func printOriginalSlice(subslice *[]int) {
	data := (*[B]int)(unsafe.Pointer(((*reflect.SliceHeader)(unsafe.Pointer(subslice))).Data)) //A

	fmt.Printf("original\t%p:%+v\n", data, *data)
}

func main() {
	slice := make([]int, A)
	for i, _ := range slice {
		slice[i] = i
	}
	subslice := slice[0:B]

	fmt.Printf("slice\t%p:%+v\n", &slice, slice)
	fmt.Printf("subslice\t%p:%+v\n", &subslice, subslice)
	printOriginalSlice(&subslice)
}


