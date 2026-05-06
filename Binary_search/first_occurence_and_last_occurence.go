package main
import "fmt"


func first_occurence(arr []int,k int)int{
	left:=0
	right:=len(arr)-1
	ans:=-1
	for left<=right {
		mid:=left+(right-left)/2
		if arr[mid]==k{
			ans=mid
			right=mid-1
		}else if arr[mid] < k {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}
func last_occurence(arr []int,k int)int{
	left:=0
	right:=len(arr)-1
	ans:=-1
	for left<=right{
		mid:=left+((right-left)/2)
		if arr[mid]==k{
			ans=mid
			left=mid+1
		}else if arr[mid]>k{
			right=mid-1

		}else {
			left=mid+1
		}
	}
	return ans 

}
func main(){
	arr:=[]int{5,7,7,8,8,10}
	k:=8
	x:=first_occurence(arr,k)
	y:=last_occurence(arr,k)
	fmt.Println(x,y)
}


// arr = [5,7,7,8,8,10]
// target = 8

// Output = [3,4]