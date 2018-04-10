package main

import (
	"fmt"
	"net"
	"bufio"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:8053")
	if err != nil {
		fmt.Printf("fail to listen :%s\n", err.Error())
	}
	for{
		if c,err := l.Accept() ; err == nil {
			Echo(c)
		}
	}
}


func Echo(c net.Conn){
	defer c.Close()

	line,err := bufio.NewReader(c).ReadString('\n')

	if err != nil {
		fmt.Printf("fail to read : %s\n",err.Error())
		return
	}

	_,err0 := c.Write([]byte(line))

	if err != nil {
		fmt.Printf("Failure to write: %s\n", err0.
			Error())
		return
	}
}