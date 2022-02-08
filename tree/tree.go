package tree

import "fmt"

/*
	          10
			/    \
		   5      15
		  /\      /\
		 3  6   12  18
*/
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func Insert(head *Node, val int) *Node {
	if head == nil {
		return &Node{val, nil, nil}
	}
	if val <= head.Val {
		head.Left = Insert(head.Left, val)
	} else {
		head.Right = Insert(head.Right, val)
	}
	return head
}

func InOrder(head *Node) {
	if head == nil {
		return
	}
	InOrder(head.Left)
	fmt.Println(head.Val)
	InOrder(head.Right)
}
