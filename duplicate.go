package main
import "fmt"
func findduplicate(arr []int)bool {
	m:=make(map[int]bool)
	for _,num:=range arr{
		if m[num]{
			return true
			}
		m[num] = true}
	return false 	
		}
func main(){
	arr:=[]int{1,2,3,1}
	x:=findduplicate(arr)
	fmt.Println(x)

}