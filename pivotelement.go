//  pivot element sum 
// nums := []int{1,7,3,6,5,6}
package main
import "fmt"

func prefix(n int,nums []int)int{
	
	sum:=0
	for i:=0;i<n;i++{
		sum+=nums[i]
	}
	return sum
}
func totalsum(nums []int)int{
	sum:=0
	for i:=0;i<len(nums);i++ {
		sum+=nums[i]
	}
	return sum
}
func pivotsum(nums []int)int{
	total:=totalsum(nums)
	leftsum:=0
	for i:=0;i<len(nums);i++ {
		
		rightsum:=total-leftsum-nums[i]
		if leftsum==rightsum{
			return i
		}
		leftsum+=nums[i]
	}
	return -1
}
func main(){
	nums:=[]int{1,7,3,6,5,6}
	x:=pivotsum(nums)
	fmt.Println(x)
}