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

// TestLink
func TestLinkListGeneral(t *testing.T) {
	// https://leetcode.com/problems/odd-even-linked-list/description
	// 這個範例很好解釋了 dummy 與 tail的用法 
	func oddEvenList(head *ListNode) *ListNode {
	    // 因為dummy的關系，就算head是nil，dummy與tail也會好好的指在dummy上，所以不用判斷head == nil
	    // if head == nil || head.Next == nil { return head }
	
	    // dummy 用來讓loop變優雅，主要是可以去掉如果第一個值是空的情況，有了dummy的話至少都會是dummy。
	    dummy0, dummy1 := &ListNode{}, &ListNode{}
	    tail0, tail1 := dummy0, dummy1
	    // dummy 和 tail 可以指到同一個，新增的時候控制tail增加即可
	
	    
	    i := 0
	    for cur := head; cur != nil; cur = cur.Next {
	        i++
	        if i % 2 == 1 {
	            tail1.Next, tail1 = cur, cur
	        } else {
	            tail0.Next, tail0 = cur, cur
	        }
	    }
	
	    // 接上
	    // 因為dummy的關系，tail不會是nil，所以這裡不需要額外的判斷
	    tail1.Next = dummy0.Next
	    tail0.Next = nil
	
	    return dummy1.Next 
	}
}

/*
反轉
// https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list/submissions/?envType=study-plan-v2&envId=leetcode-75
var pre *ListNode
cur := head
for cur != nil {
	cur.Next, cur, pre = pre, cur.Next, cur
}
1>2>3>4>5
1<2<3<4<5
        p c

*/

/*
// https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list/submissions/?envType=study-plan-v2&envId=leetcode-75
快慢指針
for f != nil && f.Next != nil {
	s, f = s.Next, f.Next.Next
}

s, f = head, head
奇數 s: 中點
1 2 3 4 5
    s   f
偶數 s: 右中點
1 2 3 4 5 6
      s     f
- s永遠是中點或後半的head
- f == nil { even node } 
→ 操作後半

s, f = dummy, dummy
奇數 s: 中點
d 1 2 3 4 5
      s     f
偶數 s: 左中點
d 1 2 3 4 5 6
      s     f
- s永遠是中點或前半的tail
- f != nil { even node }
→ 用於切鏈表

*/
