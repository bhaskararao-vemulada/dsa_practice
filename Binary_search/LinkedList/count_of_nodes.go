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
func traverse(){
	head:=creating_linked_list()
	curr:=head
	count:=0
	for curr!=nil{
		fmt.Print(curr.data,"->")
		count++
		curr=curr.next

	}
	fmt.Println("nil")
}

func count_of_nodes()int{
	head:=creating_linked_list()
	curr:=head
	count:=0
	for curr!=nil{
		
		count++
		curr=curr.next

	}
	
	return count
}

func main(){
	traverse()
	x:=count_of_nodes()
	fmt.Println(x)
}