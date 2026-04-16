// arr = [1, 3, 5, 7, 9, 11]
// target = 7
package main
import ("fmt")

func binarysearch(arr []int,target int)int{
	left:=0
	right:=len(arr)-1
	mid:=0
	for left <=right {
		mid=left+(right-left)/2
		if arr[mid]==target{
			return mid
		}
		if arr[mid]>target {
			right =mid-1
		}else {
			left=mid+1
		}
		
	}
	return -1
	
}

func main(){
	arr:=[]int{1, 3, 5, 7, 9, 11}
	target:=7
	fmt.Println(binarysearch(arr,target))
}