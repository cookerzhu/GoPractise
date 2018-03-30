package main


type student struct {
	name string
}

func main(){
	m:=map[string]student{"people":{"zhoujielun"}}
	m["people"].name = "wuyanzu"
}
