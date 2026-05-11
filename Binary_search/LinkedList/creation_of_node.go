package main

import"fmt"

type Node struct {
	data int
	next *Node
}

func creation_of_node(data int)*Node{
	n:=&Node{
		data: 10,
		next: nil,
}
fmt.Println(n.data)
fmt.Println(n.next)
	return n
}
func main(){
	y:=creation_of_node(10)
	fmt.Println(y)

}