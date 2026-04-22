package main 

import "fmt"

func removeduplicates(arr []int)([]int,int){
	slow:=0
	for fast:=1;fast<len(arr);fast++{
		if arr[fast]!=arr[slow]{
			slow++
			arr[slow]=arr[fast]

		}

	}
	return arr[:slow+1],slow+1
}

func main(){
	arr:=[]int{1,1,2,2,3}
	x,y:=removeduplicates(arr)
	fmt.Println(x,y)
}