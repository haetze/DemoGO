package main

import "fmt"

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
	fmt.Println(node)
}

