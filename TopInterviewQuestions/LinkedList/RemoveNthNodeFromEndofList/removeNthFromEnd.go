package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Append(val int) *ListNode {
	newNode := &ListNode{Val: val}

	current := l
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode

	return newNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pos := 0
	var deleteNode *ListNode

	if head.Next == nil {
		return nil
	}

	for current := head; current != nil; current = current.Next {
		pos++
		if pos-n-1 >= 0 {
			if deleteNode == nil {
				deleteNode = head
			} else {
				deleteNode = deleteNode.Next
			}
		}
	}

	if deleteNode == nil {
		return head.Next
	}

	if deleteNode.Next != nil && deleteNode.Next.Next == nil {
		deleteNode.Next = nil
	} else if deleteNode.Next != nil && deleteNode.Next.Next != nil {
		deleteNode.Next = deleteNode.Next.Next
	}

	return head
}

func PrintList(l *ListNode) {
	if l == nil {
		return
	}
	current := l
	for {
		fmt.Print(current.Val, " ")
		if current.Next == nil {
			break
		}
		current = current.Next
	}
}

func main() {
	list := &ListNode{Val: 1}
	list.Append(2)
	list.Append(6)

	list = removeNthFromEnd(list, 3)
	PrintList(list)
}
