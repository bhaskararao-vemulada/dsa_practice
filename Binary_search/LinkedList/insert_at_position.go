package main

import "fmt"
type Node struct{
	data int
	next *Node

}

func creating_linked_list()*Node{
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

func insert_at_position(position int)*Node{
	n:=&Node{
		data: 1,
		next: nil,
	}
	head:=creating_linked_list()
	curr:=head
	count:=0
	for curr !=nil{
		count++
		if count+1==position{
			n.next=curr.next
			curr.next=n
			return head
			
		}
		curr=curr.next
		


	}
	return head



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
	head:=insert_at_position(3)
	
	traverse(head)
}

