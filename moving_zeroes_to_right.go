package main
import "fmt"

func moving_all_zeroes_to_right(arr[]int)[]int{
	slow:=0
	for fast:=0;fast<len(arr);fast++{
		if arr[fast]!=0{
			arr[fast],arr[slow]=arr[slow],arr[fast]
			slow++
		}
		
	}
	return arr
}
func main(){
	arr:=[]int{0, 1, 0, 5,0,0,0,0,0,0,0,0,0,0,0,3, 12}
	x:=moving_all_zeroes_to_right(arr)
	fmt.Println(x)
}