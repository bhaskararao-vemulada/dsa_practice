package main
import "fmt"
// arr = []int{1,1,1}
// k = 2

func count_of_subarray_sum(arr []int,k int)int{

    count:=0
   
    sum:=0
    for left:=0;left<len(arr);left++{
        sum=0
        for right:=left;right<len(arr);right++{
       
           sum+=arr[right]
           if sum==k{
               count++
        }

    }
   



}
 return count }

func main(){
    arr:=[]int{1,1,1}
    k:=2
    x:=count_of_subarray_sum(arr,k)
    fmt.Println(x)
}