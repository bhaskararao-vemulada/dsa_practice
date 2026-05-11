package main
import "fmt"
type Node struct {
	data int
	next *Node
}

func creation_of_multiple_nodes()(*Node,*Node,*Node){
	n2:=&Node{
		data:30,
		next:nil,
		
	}
	n1:=&Node{
		data:20,
		next:nil,
	
	}

	n:=&Node{
		data:10,
		next:nil,
	}

	return n,n1,n2}
func main(){
	x,y,z:=creation_of_multiple_nodes()
	fmt.Println(x,y,z)

}