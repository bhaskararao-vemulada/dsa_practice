package main
import "fmt"
func Firstinteger(arr []int,target int)int {
	left:=0
	right:=len(arr)-1
	ans:=-1
	for left<=right{
		mid:=left+(right-left)/2
      
		if arr[mid]>=target{
			ans = mid 
			right=mid-1
			
			
		}else {
			left=mid+1
		}
	}
	return ans

}

func main(){
	arr:=[]int{1, 2, 2, 2, 3, 4}
	target:=2
	fmt.Println(Firstinteger(arr,target))
}