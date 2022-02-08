package linklist

import "fmt"

type Node struct {
	val  int
	next *Node
}
type LinkList struct {
	head *Node
}

func (l *LinkList) AddNode(val int) {
	cur := l.head
	if cur == nil {
		cur = &Node{val: val, next: nil}
	} else {
		for cur != nil {
			cur = cur.next
		}
		cur.next = &Node{val: val, next: nil}
	}
}

func (l *LinkList) Print() {
	cur := l.head
	for cur != nil {
		fmt.Println(cur.val)
		cur = cur.next
	}
}

func main() {
	l := LinkList{}
	l.AddNode(2)
	l.Print()
}
