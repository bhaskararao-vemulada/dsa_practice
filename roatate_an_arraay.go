package main
import "fmt"



func reverse(arr []int, start, end int) {
    for start < end {
        arr[start], arr[end] = arr[end], arr[start]
        start++
        end--
    }
}

func rotate(arr []int, k int) {
    n := len(arr)
    if n == 0 {
        return 
    }

    k = k % n

    reverse(arr, 0, n-1)
    reverse(arr, 0, k-1)
    reverse(arr, k, n-1)
}


func main(){
	arr:=[]int{1,2,3,4,5,6,7}
	k:=3
	rotate(arr,k)
	fmt.Println(arr)

}