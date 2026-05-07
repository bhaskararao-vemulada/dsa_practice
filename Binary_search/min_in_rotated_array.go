package main
import "fmt"
func min_in_rotated_array(arr []int)int{
	left:=0
	right:=len(arr)-1

	for left<right {
		mid:=left+((right-left)/2)

	
		if arr[mid]<arr[right]{
			
			right=mid
		}else if arr[mid]>arr[right]{
			left=mid+1
		}
	}
	return arr[left]
	
	

}
func main(){
	arr:=[]int{4,5,6,7,0,1,2}
	x:=min_in_rotated_array(arr)
	fmt.Println(x)
}