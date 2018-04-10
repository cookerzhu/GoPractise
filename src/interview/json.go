package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

type People struct {
	name string `json:"name"`
}

//type student struct{
//	Name string
//}
//
//func zhoujielun(v interface{}){
//	switch msg:= v.(type){
//	case *student,student:
//		msg.Name
//	}
//}

func main(){
	js := `{
		"name":"11"
	}`
	var p People
	err := json.Unmarshal([]byte(js),&p)
	if err != nil {
		fmt.Println("err:",err)
		return
	}
	fmt.Println("people:",p)//{}
	fmt.Println([]byte(js))


}


