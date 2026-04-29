package main
import "fmt"
func averageofksubarray(arr []int, k int) []float64 {
    result := []float64{}
    windowSum := 0

    // build first window
    for i := 0; i < k; i++ {
        windowSum += arr[i]
    }

    result = append(result, float64(windowSum)/float64(k))

    // slide window
    for i := k; i < len(arr); i++ {
        windowSum += arr[i]
        windowSum -= arr[i-k]

        result = append(result, float64(windowSum)/float64(k))
    }

    return result
}
func main(){
	arr:=[]int{1, 3, 2, 6, -1, 4, 1, 8, 2}
	k:=5
	x:=averageofksubarray(arr,k)
	fmt.Println(x)
}