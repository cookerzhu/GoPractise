package main

/**
在Go的for…range循环中，Go始终使用值拷贝的方式代替被遍历的元素本身，
简单来说，就是for…range中那个value，是一个值拷贝，而不是元素本身。
这样一来，当我们期望用&获取元素的指针地址时，实际上只是取到了value这个临时变量的指针地址
而在整个for…range循环中，value这个临时变量会被重复使用
 */
type Student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*Student)
	stus := []Student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	//错误写法
	for _, stu := range stus {
		//stu其实是stus元素的值拷贝，并且在每次循环被重复利用，其内存地址是同一个地址
		m[stu.Name] = &stu
	}

	//正确写法
	for i, _ := range stus {
		m[stus[i].Name] = &stus[i]
	}

	//for k,v:=range m{
	//	println(k,"=>",v.Name)
	//}

}

func main(){
	pase_student()
}