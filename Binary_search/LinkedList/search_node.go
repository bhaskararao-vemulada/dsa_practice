package main
import "fmt"

type Node struct {
	data int
	next *Node 
}
func creating_linked_list()*Node{
	n4:=&Node{
		data: 40,
		next: nil,
	}
	n3:=&Node{
		data: 30,
		next: n4,
	}
	n2:=&Node{
		data: 20,
		next: n3,
	}
	n1:=&Node{
		data: 10,
		next: n2,
	}
	
	return n1
}

func search(target int)bool{
	head:=creating_linked_list()
	curr:=head
	ans:=false
	for curr!=nil {
		if curr.data==target{
			ans=true
			return ans
		}
		curr=curr.next
		 

	}
	return false 

}
func main(){
	x:=search(50)
	fmt.Println(x)
}