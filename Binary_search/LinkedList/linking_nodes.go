package main
import "fmt"
type Node struct {
	data int
	next *Node 
}

func linking_nodes()*Node{
	n:=&Node{
		data: 10,
		next: nil,
	}
		n1:=&Node{
		data: 10,
		next: nil,
	}
		n2:=&Node{
		data: 10,
		next: nil,
	}
	n.next=n1
	n1.next=n2
	
	return n
}
func main(){
	x:=linking_nodes()
	fmt.Println(x)
}

