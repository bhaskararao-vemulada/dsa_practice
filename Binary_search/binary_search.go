package main
import "fmt"
func binarysearch(arr []int,k int)int{
	left:=0
	right:=len(arr)-1
	for left<=right {
		mid:= left+(right-left)/2
		if arr[mid]==k{
			return mid
		}else if arr[mid]<k {
			left=mid+1
		}else if arr[mid]>k {
			right=mid-1
		}
	}
	return -1
}

func main(){
	
	arr:=[]int{-1, 0, 3, 5, 9, 12}
	k:=9
	x:=binarysearch(arr,k)
	fmt.Println(x)
}