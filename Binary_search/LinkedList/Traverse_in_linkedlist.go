package main
import "fmt"
type Node struct {
	data int
	next *Node
}

func create_linked_list()(*Node){
	n3:=&Node{
		data:10,
		next:nil,
	}	
	n2:=&Node{
		data:20,
		next:n3,
	}
	n1:=&Node{
		data:30,
		next:n2,
	}
	n:=&Node{
		data:40,
		next:n1,
	}
	
	
	return n 

}
func traverse()*Node{
	head:=create_linked_list()
	curr:=head
	for curr!=nil{
		fmt.Println(curr.data)
		
		curr=curr.next
		
	}
	return curr
}
func main(){
	x:=traverse()
	
	

}