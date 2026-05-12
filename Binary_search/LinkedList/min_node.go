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
		data: 5,
		next: n2,
	}
	return n1
}

func min_node()int{
	head:=creating_linked_list()
	curr:=head
	min:=curr.data
	for curr!=nil {
		if curr.data<min{
			min=curr.data
		}
		curr=curr.next
	}
	return min
}
func main(){
	x:=min_node()
	fmt.Println(x)
}