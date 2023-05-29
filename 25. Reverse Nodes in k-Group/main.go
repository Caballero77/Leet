package main

type ListNode struct {
  Val int
  Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
  result := &ListNode{ Val: -1, Next: head }
  groupPrev := result

  for true {
    groupEnd := groupPrev
    i := k
    for groupEnd != nil && i > 0 {
      groupEnd = groupEnd.Next
      i--
    }
    if groupEnd == nil {
      break
    }
    groupNext := groupEnd.Next

    curr := groupPrev.Next
    prev := groupEnd.Next
    for curr != groupNext {
      tmp := curr.Next
      curr.Next = prev
      prev = curr
      curr = tmp
    }

    tmp := groupPrev.Next
    groupPrev.Next = groupEnd
    groupPrev = tmp

  }

  return result.Next
}


func main() {

}
