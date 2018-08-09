package main

/**
defer 是后进先出。
panic 需要等defer 结束后才会向上传递。 出现panic恐慌时候，会先按照defer的后入先出的顺序执行，最后才会执行panic。
 */
import (
	"fmt"
)

func main() {
	defer_call()
}

func defer_call() {
	recover()
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}