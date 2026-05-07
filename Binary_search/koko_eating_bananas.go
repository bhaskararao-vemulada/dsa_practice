package main
import "fmt"
// Input:
// piles = [3,6,7,11]
// h = 8

// Output:
// 4

func max_of_array(arr []int)int{
	max:=arr[0]
	
	for i:=0;i<len(arr);i++ {
		if arr[i]>max{
			max=arr[i]
		}
		
		
	}
	return max

}
func finding_hours(arr []int,speed int)int{
	hours:=0
	for i:=0;i<len(arr);i++ {
		hours+=(arr[i]+speed-1)/speed
	}
	return hours
}
func koko_eating_bananas(arr []int,h int)int{
	left:=1
	right:= max_of_array(arr)
	speed:=1
	ans := right
	for left<=right{
		speed=left+((right-left)/2)
		x:=finding_hours(arr,speed)
		if x<=h {
			ans=speed
			right=speed-1
		}else {
			left=speed+1
		}
		
		
		
	}
	return ans
		
	}
func main(){
	arr:=[]int{3,6,7,11}
	h:=8
	y:=koko_eating_bananas(arr,h)
	fmt.Println(y)
}

