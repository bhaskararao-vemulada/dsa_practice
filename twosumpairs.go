package main
import "fmt"

func sumpairs(arr []int, target int) [][]int {
	result := [][]int{}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				result = append(result, []int{arr[i], arr[j]})
			}
		}
	}

	return result
}

func main() {
	arr := []int{1, 5, 7, -1, 5}
	target := 6

	output := sumpairs(arr, target)
	fmt.Println(output)
}