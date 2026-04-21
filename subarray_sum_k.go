package main
import "fmt"
func subarray_sum_k(arr []int, k int) int {
	m := make(map[int]int)
	m[0] = 1 

	sum := 0
	count := 0

	for _, num := range arr {
		sum += num

		if val, ok := m[sum-k]; ok {
			count += val
		}

		m[sum]++ 
	}

	return count
}

func main() {
	arr := []int{1, 2, 1}
	k := 3
	fmt.Println(subarray_sum_k(arr, k))
}