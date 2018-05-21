/**
@See http://www.cnblogs.com/mushroom/p/8998538.html
 */
package main
import "fmt"
func main() {
	s := []byte("")//cap is 32 ,len is 0 没有下面的注释，slice分配在栈上面，初始化大小是32的数组
	s1 := append(s, 'a')
	s2 := append(s, 'b')
	//fmt.Println(s1, "==========", s2)   去掉注释，会发生逃（slice数组会分配到堆上，是空数组，s1,s2都是引用复制的新slice）
	fmt.Println(string(s1), "==========", string(s2))


	news := []byte{} // cap and len is 0
	fmt.Println(cap(news), len(news))
}