package main
import "fmt"

func peak_element(arr []int) int {

	if len(arr) == 0 {
		return -1
	}

	left := 0
	right := len(arr) - 1

	for left < right {

		mid := left + (right-left)/2

		if arr[mid] < arr[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func main(){
	arr:=[]int{1,2,1,3,5,6,4}
	x:=peak_element(arr)
	fmt.Println(x)
}