package main
import "fmt"
type Node struct {
	data int
	next *Node
}

func create_linked_list() *Node {
	n4 := &Node{
		data: 40,
		next: nil,
	}
	n3 := &Node{
		data: 30,
		next: n4,
	}
	n2 := &Node{
		data: 20,
		next: n3,
	}
	n1 := &Node{
		data: 10,
		next: n2,
	}
	return n1
}

func insert_end(head *Node) *Node {
	n5 := &Node{
		data: 50,
		next: nil,
	}

	if head == nil {
		return n5
	}

	curr := head

	for curr.next != nil {
		curr = curr.next
	}

	curr.next = n5

	return head
}

func traverse(head *Node) {
	curr := head

	for curr != nil {
		fmt.Print(curr.data, "->")
		curr = curr.next
	}

	fmt.Println("nil")
}

func main() {
	head := create_linked_list()

	traverse(head)

	head = insert_end(head)

	traverse(head)
}