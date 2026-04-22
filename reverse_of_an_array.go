package main
import "fmt"


// arr = []int{1, 2, 3, 4, 5}

func reverse_of_array(arr []int)[]int{
	left:=0
	right:=len(arr)-1

	for left<=right {
		arr[left],arr[right]=arr[right],arr[left] 
		left++
		right--
	}
	return arr
}
func main(){
	arr:=[]int{1, 2, 3, 4, 5}
	x:=reverse_of_array(arr)
	fmt.Println(x)
}