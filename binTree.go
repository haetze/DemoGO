package main

import (
	"fmt"
)


type Node struct{
	val int
	left *Node
	right *Node
}

func (n *Node) inOrderPrint(){
	if n.left != nil {
		n.left.inOrderPrint();
	}
	fmt.Println(n.val);
	if n.right != nil {
		n.right.inOrderPrint();
	}
}

func (n *Node) init(val int){
	n.val = val
}

func (n *Node) add(m *Node){
	if n.val > m.val {
		if n.left == nil {
			n.left = m
		}else{
			n.left.add(m)
		}
	}else if n.val < m.val {
		if n.right == nil {
			n.right = m
		}else{
			n.right.add(m)
		}
	}
}


func main(){
	n := Node{}
	n.init(12)
	n1 := Node{}
	n1.init(6)
	n2 := Node{}
	n2.init(24)
	n3 := Node{}
	n3.init(16)
	n4 := Node{}
	n4.init(9)
	n.add(&n1)
	n.add(&n2)
	n.add(&n3)
	n.add(&n4)
	n.inOrderPrint()
	//fmt.Println("Hello World!")
	
}
