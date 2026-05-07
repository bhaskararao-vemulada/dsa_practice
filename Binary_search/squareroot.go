package main
import "fmt"
func squareroot(x int)int{
	left:=0
	right:=x
	ans:=-1
	for left<=right{
		mid:=left+((right-left)/2)
		if mid*mid==x{
			ans=mid
			return ans
		}
		if mid  <= x/mid{
			ans=mid
			left=mid+1
		}else{
			right=mid-1
		}
	}
	return ans
}
func main(){
	x:=8
	y:=squareroot(x)
	fmt.Println(y)
}