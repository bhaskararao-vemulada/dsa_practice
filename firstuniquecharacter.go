package main
import "fmt"
func firstunique(s string)int{
	m:=make(map[rune]int)
	for _,ch:=range s{
		m[ch]++
	}
	for i,ch:=range s {
		if m[ch]==1{
			return i
		}
	}
	return -1
}
func main(){
	s:="lol"
	x:=firstunique(s)
	fmt.Println(x)
}
// optimized  
// o(n) time complexity 