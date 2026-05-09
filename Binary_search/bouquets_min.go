package main
import "fmt"

// bloomDay = []int{1, 2, 4, 9, 3, 4, 1}
// m = 2
// k = 2

// Output:

// 4
func max_days(arr []int)int {
	max:=0
	for i:=0;i<len(arr);i++ {
		if arr[i]>max {
			max=arr[i]
		}
		
	}
	return max 
}

func counting_days(arr []int,mid int,m,k int)bool{
	count:=0
	bouquets:=0

	for i:=0;i<len(arr);i++ {
		if arr[i]<=mid{
			count++
		} else{
			count=0
		}
		if count==k{
			bouquets++
			count=0
		}
	}
	return bouquets>=m
	


}
func blooming_days_bouquets(bloomDay []int,m int,k int)int{
	if len(bloomDay)<m*k{
		return -1
	}
	left:=1
	right:=max_days(bloomDay)
	ans:=-1
	for left<=right {
		mid:=left+((right-left)/2)
		if counting_days(bloomDay,mid,m,k){
			ans=mid
			right=mid-1

		}else{
			left=mid+1
		}

	}
	return ans 


}

func main(){
	bloomDay:=[]int{1, 2, 4, 9, 3, 4, 1}
	m:=2
	k:=2
	x:=blooming_days_bouquets(bloomDay,m,k)
	fmt.Println(x)
}