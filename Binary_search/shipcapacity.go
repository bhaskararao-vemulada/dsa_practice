package main
import "fmt"
// weights = []int{1,2,3,4,5,6,7,8,9,10}
// days = 5
// Output
// 15
func finding_capacity(weights []int, capacity int) int {
	days := 1          // start with first day
	currentLoad := 0   // weight loaded today

	for i := 0; i < len(weights); i++ {

		// If current package fits in today's ship
		if currentLoad+weights[i] <= capacity {
			currentLoad += weights[i]

		} else {
			// Start a new day
			days++
			currentLoad = weights[i]
		}
	}

	return days
}
func max_weight(weights []int)int{
	sum:=0
	maxsum:=0
	for i:=0;i<len(weights);i++ {
		sum+=weights[i]
		if sum> maxsum{
			maxsum=sum
		}
	}
	return maxsum
}

func min_weight(weights []int)int {
	maxweight:=0
	for i:=0;i<len(weights);i++ {
		
		if weights[i]> maxweight{
			maxweight=weights[i]
		}
	}
	return maxweight
}





func ship_capacity(weights []int,days int)int{
	min_capacity:=min_weight(weights)
	max_capacity:=max_weight(weights)
	ans:=0
	for min_capacity<= max_capacity{
		capacity:=min_capacity+((max_capacity-min_capacity)/2)
		y:=finding_capacity(weights,capacity)
		if y<=days{
			ans=capacity
			max_capacity=capacity-1
			

		}else {
			
			min_capacity=capacity+1
		}





	}
	return ans

}
func main(){
	weights:=[]int{1,2,3,4,5,6,7,8,9,10}
	days:=5
	ans:=ship_capacity(weights,days)
	fmt.Println(ans)
	
}

