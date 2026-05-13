package main
import "fmt"
type Node struct {
	data int
	next *Node
}

func creating_linked_list()*Node{
	n4:=&Node{
		data : 40,
		next : nil,
	}
	n3:=&Node{
		data : 30,
		next : n4,
	}
	n2:=&Node{
		data : 20,
		next : n3,
	}
	n1:=&Node{
		data : 10,
		next : n2,
	}
	return n1
}

func traverse(head *Node){
	
	curr:=head
	for curr!=nil {
		fmt.Print(curr.data,"->")
		curr=curr.next
	}
	fmt.Println("nil")
}

func delete_node(head *Node)*Node{
	if head==nil{
		return nil
	}
	

	

	
	head=head.next
	return head








}
func main(){
	head:=creating_linked_list()
	traverse(head)
	head=delete_node(head)
	traverse(head)
}

