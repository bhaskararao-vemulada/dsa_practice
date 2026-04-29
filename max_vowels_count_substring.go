package main
import "fmt"


// s := "abciiidef"
// k := 3


func isvowel(ch byte)bool {
	if ch=='a'||ch=='e'||ch=='i'||ch=='o'||ch=='u'{
		return true 
	}
	return false 
}

func max_vowel_in_substring(s string, k int)int{
	count:=0
	for i:=0;i<k;i++{
	    if isvowel(s[i]){
		  count++
	    }

	}
	max:=count
	for i:=k;i<len(s);i++ {
		if isvowel(s[i]){
			count++

		}
		if isvowel(s[i-k]) {     // remove old
            count--
        }
		 if count > max {
            max= count
        }
		
	}
	return max
}

func main(){
	s:="abciiidef"
	k:=3
	x:=max_vowel_in_substring(s,k)
	fmt.Println(x)
}