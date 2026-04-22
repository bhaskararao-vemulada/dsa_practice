package main
import "fmt"
func findingmax(arr []int)int{
	
	if len(arr)==0{
		return -1
	}
	max:=arr[0]
	for i:=0;i<len(arr);i++ {
		if arr[i]>max {
			max=arr[i]
		}



	}
	return max 
}
func main(){
	arr:=[]int{-3, -7, -2, -1, -4}
	x:=findingmax(arr)
	fmt.Println(x)
}