package main

/*
You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.



Example 1:

Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6
Example 2:

Input: lists = []
Output: []
Example 3:

Input: lists = [[]]
Output: []


Constraints:

k == lists.length
0 <= k <= 104
0 <= lists[i].length <= 500
-104 <= lists[i][j] <= 104
lists[i] is sorted in ascending order.
The sum of lists[i].length will not exceed 104.

*/

/**
* Definition for singly-linked list.
 */
type ListNode struct {
    Val  int
    Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }

    if len(lists) == 1 {
        return lists[0]
    }

    for len(lists) > 1 {
        list1 := lists[0]
        list2 := lists[1]
        lists = lists[2:]

        if list1 == nil && list2 == nil {
            continue
        }

        var head *ListNode
        ptr := head
        var prev *ListNode

        for list1 != nil && list2 != nil {
            if head == nil {
                head = new(ListNode)
                ptr = head
            }

            if ptr == nil {
                ptr = new(ListNode)
                prev.Next = ptr
            }

            if list1.Val < list2.Val {
                ptr.Val = list1.Val
                list1 = list1.Next
            } else {
                ptr.Val = list2.Val
                list2 = list2.Next
            }

            prev = ptr
            ptr = ptr.Next
        }

        if list1 != nil {
            if head == nil {
                head = list1
            } else if ptr == nil {
                prev.Next = list1
            }
        }

        if list2 != nil {
            if head == nil {
                head = list2
            } else if ptr == nil {
                prev.Next = list2
            }
        }

        lists = append(lists, head)
        head = nil
    }

    if len(lists) == 1 {
        return lists[0]
    }
    return nil
}