package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {

  for head != nil && head.Val == val {
    head = head.Next
  }

  curr := head

  for curr != nil && curr.Next != nil {
    if curr.Next.Val == val {
      curr.Next = curr.Next.Next
    }

    if curr.Next != nil && curr.Next.Val != val {
      curr = curr.Next
    }

  }

	return head
}

func main() {
	fmt.Println("Hello World")
}
