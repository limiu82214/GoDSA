package linklist_test

type Node struct {
	Prev, Next *Node
	Key, Val   int
}

var Head *Node
var Tail *Node

func InitList() {
	//init
	Head, Tail = &Node{}, &Node{}
	Head.Next = Tail
	Tail.Prev = Head
}

func InsertHead(n *Node) {
	Head.Next.Prev, Head.Next, n.Next, n.Prev = n, n, Head.Next, Head
}
func Remove(n *Node) {
	n.Prev.Next, n.Next.Prev = n.Next, n.Prev
}
