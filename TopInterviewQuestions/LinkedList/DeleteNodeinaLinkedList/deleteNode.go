package main

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

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func main() {
	list := &ListNode{Val: 1}
	addr := list.Append(5)
	list.Append(6)

	deleteNode(addr)
}
