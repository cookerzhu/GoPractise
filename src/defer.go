package main 

import (
	"fmt"
)

//defer语句，函数执行到最后，在返回前执行

func main() {
	for index := 1;index <= 5;index ++{
		defer fmt.Println(index);
	}
	
	defer fmt.Println("后定义的defer语句 先执行，后进先出，类似于栈")
}

