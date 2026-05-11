package main
import "fmt"
// max node 

type Node struct {
	data int
	next *Node
}

func create_linked_list()*Node{
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

func max_node()int{
	if head==nil{
		return -1
	}
	head:=create_linked_list()
	curr:=head
	max:=curr.data
	for curr!=nil {
		if curr.data>max {
			max=curr.data
		}

		curr=curr.next
	}
	return max 
}

func main(){
	x:=max_node()
	fmt.Println(x)
}