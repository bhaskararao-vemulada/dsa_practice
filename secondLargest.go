package main

import "math"
import "fmt"

func secondLargest(arr []int) int {
    if len(arr) < 2 {
        return -1
    }

    first := math.MinInt
    second := math.MinInt

    for _, v := range arr {
        if v > first {
            second = first
            first = v
        } else if v > second && v != first {
            second = v
        }
    }

    if second == math.MinInt {
        return -1
    }

    return second
}

func main(){
	arr:=[]int{10, 20, 4, 45, 99}
	x:=secondLargest(arr)
	fmt.Println(x)
}