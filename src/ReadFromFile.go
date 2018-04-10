/**
	从文件读取内容
 */
package main

import (
	"os"
	"bufio"
	"fmt"
)

func main(){
	buf:=make([]byte,254)

	f,_ := os.Open("F:\\1.txt")
	defer f.Close()

	//for{
	//	n,_ := f.Read(buf)
	//	if n == 0 {break}
	//	fmt.Println(n)
	//	os.Stdout.Write(buf[0:n])
	//}


	r := bufio.NewReader(f)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for{
		n,_:=r.Read(buf)
		fmt.Println(n)
		if n==0 {break}
		w.Write(buf[0:n])
	}
}
