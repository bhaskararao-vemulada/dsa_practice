package main
import "fmt"
func search_in_rotatedarray(arr []int, k int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == k {
			return mid
		}

		// left half sorted
		if arr[left] <= arr[mid] {

			// target inside left half
			if k >= arr[left] && k < arr[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}

		} else {

			// right half sorted
			if k > arr[mid] && k <= arr[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

func main(){
	arr:=[]int{4,5,6,7,0,1,2}
	k:=0
	x:=search_in_rotatedarray(arr,k)
	fmt.Println(x)
}