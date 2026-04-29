package main
import "fmt"
func non_repeating_substring_length(s string)int{
	left:=0
	length:=0
	maxlength:=0

	m:=make(map[byte]int)
	for right:=0;right<len(s);right++ {
		m[s[right]]++
	
	
		for m[s[right]]>1{
			m[s[left]]--
		
			
			
			left++
			
		}
		length=right-left+1
		if length>maxlength{
			maxlength=length
		}
	}
	return maxlength

}


func main(){
	s:="abbbbcdefabcbb"
	x:=non_repeating_substring_length(s)
	fmt.Println(x)

}