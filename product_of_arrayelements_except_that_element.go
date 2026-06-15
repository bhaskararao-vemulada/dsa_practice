package main

import "fmt"


// nums := []int{1,2,3,4}
// product of array brute force 

func productofarray(arr []int)[]int{
	
	result:=[]int{}
	for i:=0;i<len(arr);i++ {
		product:=1
		for j:=0;j<len(arr);j++ {
			if j!=i{
			product*=arr[j]}


		}
		result=append(result,product)
	}
	return result
	

}
func main(){
	arr:=[]int{1,2,3,4}
	z:=productofarray(arr)
	fmt.Println(z)
}
//  optimized way 
// package main

// import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)

	prefix := make([]int, n)
	suffix := make([]int, n)
	result := make([]int, n)

	// Base cases
	prefix[0] = 1
	suffix[n-1] = 1

	// Build prefix array
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}

	// Build suffix array
	for i := n - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}

	// Build answer
	for i := 0; i < n; i++ {
		result[i] = prefix[i] * suffix[i]
	}

	return result
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}