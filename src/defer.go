package main 

//defer语句，函数执行到最后，在返回前执行

//func main() {
//	//for index := 1;index <= 5;index ++{
//	//	defer fmt.Println(index)
//	//}
//	//
//	//defer fmt.Println("后定义的defer语句 先执行，后进先出，类似于栈")
//
//
//}
//Panic Stack
func f1() {
	defer println("f1-begin")
	f2()
	defer println("f1-end")
}

func f2() {
	defer println("f2-begin")
	f3()
	defer println("f2-end")
}

func f3() {
	defer println("f3-begin")
	panic(0)
	defer println("f3-end")
}

func main() {
	f1()
}
