# 刷题内容
> 原题链接
- https://leetcode.com/problems/add-two-numbers
> 内容描述

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.

# 解题方案
## 思路1
全部变成数字做加法再换回去呗，这多暴力，爽！
> 注意：是返回头结点，且要判断最后一位数相加时carry是否大于0
```golang
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var head, tail *ListNode
    var carry int = 0
    for l1 != nil || l2 != nil{
        n1, n2 := 0, 0
        if l1 != nil{
            n1 = l1.Val
            l1 = l1.Next
        }
        if l2 != nil{
            n2 = l2.Val
            l2 = l2.Next
        }
        sum := n1 + n2 + carry
        sum, carry = sum%10, sum/10
        if head == nil{
            tail = &ListNode{Val:sum}
            head = tail
        } else {
            tail.Next = &ListNode{Val:sum}
            tail = tail.Next
        }
    }
    if carry > 0{
        tail.Next = &ListNode{Val:carry}
    }
    return head
}
```
> 复杂度分析
> - 时间复杂度：O(max⁡(m,n))，其中 m 和 n 分别为两个链表的长度。我们要遍历两个链表的全部位置，而处理每个位置只需要 O(1) 的时间。
> - 空间复杂度：O(1)。注意返回值不计入空间复杂度。
