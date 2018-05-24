package main

import (
	"encoding/json"
	"fmt"
	//"sync"
)

type People struct {
	Name string //`json:"name"`
	// 结构体里的字段首字母必须大写,如果我们想让struct转json后的首字母小写，我们可以通过字段的tag指定
}

type Result struct {
	Status int  `json:"status"`
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
	js := []byte(`{"name":"11"}`)
	var p People
	err := json.Unmarshal(js,&p)
	if err != nil {
		fmt.Println("err:",err)
		return
	}
	fmt.Println("people:",p)//{}
	fmt.Println([]byte(js))

	//var data = []byte(`{"status": 200}`)
	//result := &Result{}
	//
	//if err := json.Unmarshal(data, result); err != nil {
	//	fmt.Println("error:", err)
	//	return
	//}
	//
	//fmt.Printf("result=%+v", result)

}


