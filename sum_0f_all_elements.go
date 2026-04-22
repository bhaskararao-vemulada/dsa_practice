package main
import "fmt"
func sum_of_all_elements(arr []int)int{
	if len(arr)==0{
		return 0
	}
	sum:=0
	for _,v:=range arr{
		sum+=v
	}
	return sum
}
func main(){
	arr:=[]int{1,2,3,4,5}
	x:=sum_of_all_elements(arr)
	fmt.Println(x)
}