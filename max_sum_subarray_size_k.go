package main
import "fmt"

func max_subarray_of_size_k(arr []int, k int) int {
    n := len(arr)
    if n == 0 || k > n {
        return 0
    }

    sum := 0

    // Step 1: first window
    for i := 0; i < k; i++ {
        sum += arr[i]
    }

    maxsum := sum

    // Step 2: slide window
    for j := k; j < n; j++ {
        sum += arr[j]
        sum -= arr[j-k]

        if sum > maxsum {
            maxsum = sum
        }
    }

    return maxsum
}
func main(){
	arr:=[]int{2, 1, 5, 1, 3, 2}
	k:=3
	x:=max_subarray_of_size_k(arr,k)
	fmt.Println(x)
}