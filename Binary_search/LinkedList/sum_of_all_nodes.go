package main
import "fmt"

//  sum of all nodes in linkedlist 
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


func sum_of_all_nodes()int{
	head:=create_linked_list()
	curr:=head
	sum:=0
	for curr!=nil {
		sum+=curr.data
		curr=curr.next
	}
	return sum

}
func main(){
	x:=sum_of_all_nodes()
	fmt.Println(x)
}
