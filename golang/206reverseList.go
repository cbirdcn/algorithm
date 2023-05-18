/*
206. 反转链表

给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
 

示例 1：


输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]
示例 2：


输入：head = [1,2]
输出：[2,1]
示例 3：

输入：head = []
输出：[]
 

提示：

链表中节点的数目范围是 [0, 5000]
-5000 <= Node.val <= 5000
 

进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/reverse-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"

type listNode struct{
	Val int
	Next *listNode
}

func New() listNode{
	return listNode{Val:0, Next:nil}
}

func main(){

	head2 := New()
	head2.Val = 2
	head2.Next = nil

	head := New()
	head.Val = 1
	head.Next = &head2

	// for listNode
	for i := &head; i != nil; i = i.Next {
		fmt.Println(i)
	}


	cur := &head
	var pre *listNode
	tmp := cur.Next // (*cur).Next
	for tmp != nil {
		cur.Next = pre
		pre = cur
		cur = tmp
		tmp = tmp.Next
	}
	cur.Next = pre

	// another for listNode
	for cur.Next != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
	fmt.Println(cur.Val)
}
