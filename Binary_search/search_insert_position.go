package main
import "fmt"

func search_insert_position(arr []int,k int)int{
	left:=0
	right:=len(arr)-1
	for left<=right {
		mid:=left+((right-left)/2)
		if arr[mid]==k{
			return mid
		}else if arr[mid]>k {
			right=mid-1
		}else if arr[mid]<k {
			
			left=mid+1
		}
	}
	return left

}
func main(){
	arr:=[]int{1,3,5,6}
	k:=7
	x:=search_insert_position(arr,k)
	fmt.Println(x)
}