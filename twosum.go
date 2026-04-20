package main 


import "fmt"

// TWO SUM PROBLEM 

// Input:
// nums = [2,7,11,15]
// target = 9

// Output:
// [0,1]

// Explanation:
// nums[0] + nums[1] = 2 + 7 = 9
// BRUTE FORCE
func twosum(arr []int,target int)(int,int){
	for i:=0;i<len(arr);i++ {
		for j:=i+1;j<len(arr);j++{
			if arr[i]+arr[j]==target{
				return i,j
			}

	}
		

	}
	return -1,-1
		
	}

func main(){
	arr:=[]int{2,7,11,15}
	target:=9
	x,y:=twosum(arr,target)
	fmt.Println(x,y)
}


