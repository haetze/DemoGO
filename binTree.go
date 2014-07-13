package main

import (
	"fmt"
)


type Node struct{
	val int
	node1 *Node
	node2 *Node
}

func (n *Node) init(val int){
	n.val = val
}

func main(){
	node := Node{}
	node.init(12)
	node2 := Node{}
	node2.init(32)
	node.node2 = &node2
	fmt.Println(*node.node2)
	node.node2.val = 34
	fmt.Println(node2)
}
