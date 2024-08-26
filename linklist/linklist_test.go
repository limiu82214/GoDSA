package linklist_test

import (
	"container/list"
	"log"
	"strings"
	"testing"
)

// 這是一個簡單的linklist實作，一般情況使用list.New()即可
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

// TestBuiltinLinkedList 內建的linklist範例
func TestBuiltinLinkedList(t *testing.T) {
	logList := func(l *list.List) {
		res := []string{}
		for e := l.Front(); e != nil; e = e.Next() {
			res = append(res, e.Value.(string))
		}
		log.Println(strings.Join(res, ","))
	}

	l := list.New()
	log.Println(l.Len()) // 0

	l.PushBack("1")
	logList(l) // 1

	l.PushBack("2")
	logList(l) // 1,2

	l.PushFront("0")
	logList(l) // 0,1,2

	l.InsertAfter("1.5", l.Front().Next())
	logList(l) // 0,1,1.5,2

	l.Remove(l.Front().Next().Next())
	logList(l) // 0,1,2

	// log.Println(l.Back().Next().Value.(string)) // panic
	log.Println(l.Back().Prev().Value.(string)) // 1
}
