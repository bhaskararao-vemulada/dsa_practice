package main
import "fmt"

// nums = []int{44,22,33,11,1}
// threshold = 5

// Output:

// 44

func max_in_divisor(arr []int)int{
	max:=arr[0]
	for i:=1;i<len(arr);i++ {
		if arr[i]>max{
			max=arr[i]
		}
		
	}
	return max
}


func find_sum_divisor(arr []int,divisor int)int{
	sum:=0
	for i:=0;i<len(arr);i++ {
		sum+=(arr[i]+divisor-1)/divisor
	}
	return sum
}
func divisor(arr []int, k int)int{
	left:=1
	right:=max_in_divisor(arr)
	total:=0
	ans:=-1
	for left<=right {
		mid:=left+((right-left)/2)

		total=find_sum_divisor(arr,mid)
		if total<=k {
			ans=mid
			
			right=mid-1

		}else{
			left=mid+1
		}

	}
	return ans

}
func main(){
	arr:=[]int{100}
	k:=10
	x:=divisor(arr,k)
	fmt.Println(x)
}
