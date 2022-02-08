package Tree

import "fmt"

/*
	          10
			/    \
		   5      15
		  /\      /\
		 3  6   12  18
*/

func insert(head *Node, val int) *Node {
	if head == nil {
		return &Node{val, nil, nil}
	}
	if val <= head.val {
		head.left = insert(head.left, val)
	} else {
		head.right = insert(head.right, val)
	}
	return head
}

func inOrder(head *Node) {
	if head == nil {
		return
	}
	inOrder(head.left)
	fmt.Println(head.val)
	inOrder(head.right)
}

type Node struct {
	val   int
	left  *Node
	right *Node
}

func main() {
	bt := Node{10, nil, nil}
	fmt.Println(bt)

	insert(&bt, 15)
	insert(&bt, 12)
	insert(&bt, 18)
	insert(&bt, 5)
	insert(&bt, 3)
	insert(&bt, 6)

	// bt.insert(5)
	// bt.insert(15)
	// bt.insert(3)

	fmt.Println(bt.left)
	fmt.Println("in order -> ")
	inOrder(&bt)
}
