package main 

import (
	"fmt"
	"net/http"
	"log"
)

func say(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println(r.URL.Path)
	fmt.Fprintf(w,"Hello,Jason")
}

func main() {
	http.HandleFunc("/",say)
	err := http.ListenAndServe(":9090",nil)
	if err != nil { //if error occurs
		log.Fatal("ListenAndServe: ",err)
	}
	
}

