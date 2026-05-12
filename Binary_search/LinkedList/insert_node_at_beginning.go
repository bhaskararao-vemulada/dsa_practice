package main
import "fmt"
type Node struct {
	data int
	next *Node
}

func creating_linkedlist()*Node{
	n4:=&Node{
		data:40,
		next:nil,
	}
	n3:=&Node{
		data:30,
		next:n4,
	}
	n2:=&Node{
		data:20,
		next:n3,
	}
	n1:=&Node{
		data:10,
		next:n2,
	}
	return n1
}


func insert_node()*Node{
	n1:=creating_linkedlist()
	n:=&Node{
		data:1,
		next:n1,
	}
	return n

	
	
}

func traverse(head *Node){
	
	curr:=head
	for curr!=nil {
		fmt.Print(curr.data,"->")
		curr=curr.next
	}
	fmt.Println("nil")
}
func main(){
	x:=insert_node()
	fmt.Println(x)
	traverse(x)
}

