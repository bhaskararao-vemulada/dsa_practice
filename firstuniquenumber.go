package main
import "fmt"


func firstunique(arr []int)int {
	m:=make(map[int]int)
	for _,v:=range arr{
		m[v]++
	}
	for _,v :=range arr{
		if m[v]==1{
			return v

		}
	}
	return 0
   
}

func main(){
	arr:=[]int{2, 3,1,9,9,2,1,5, 4,4, 2, 3, 5}
	x:=firstunique(arr)
	fmt.Println(x)
}