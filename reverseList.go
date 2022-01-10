package main

import "fmt"

/**
翻转链表
https://mp.weixin.qq.com/s/pnvVP-0ZM7epB8y3w_Njwg

206：反转一个单链表。

示例: 输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL

如果再定义一个新的链表，实现链表元素的反转，其实这是对内存空间的浪费。
其实只需要改变链表的next指针的指向，直接将链表反转 ，而不用重新定义一个新的链表

*/

type listNode struct {
	Val int
	Next *listNode
}

func NewListNode() listNode{
	return listNode{
		Val:  0,
		Next: nil,
	}
}

func main(){
	//简单赋值

	listNode0 := NewListNode()
	listNode0.Val = 1

	listNode1 := NewListNode()
	listNode1.Val = 2
	listNode0.Next = &listNode1

	listNode2 := NewListNode()
	listNode2.Val = 3
	listNode1.Next = &listNode2

	listNode3 := NewListNode()
	listNode3.Val = 4
	listNode2.Next = &listNode3

	listNode4 := NewListNode()
	listNode4.Val = 5
	listNode3.Next = &listNode4

	fmt.Println("--原始链表--")

	for i := &listNode0; i != nil; i = i.Next{
		fmt.Println(i.Val)
	}

	fmt.Println("--反转链表--")
	reversedNode0 := reverseList(&listNode0)
	for i := reversedNode0; i != nil; i = i.Next{
		fmt.Println(i.Val)
	}
}

func reverseList(head *listNode) *listNode{
	cur := head
	var pre *listNode // 初始化为nil

	// 方案1：判断tmp（临时链表尾指针）
	tmp := cur.Next // cur.Next指向会变化，在此暂存，同时还用来判断是否为nil
	for tmp != nil {
		cur.Next = pre
		pre = cur
		cur = tmp
		tmp = tmp.Next
	}
	cur.Next = pre

	return cur

	// 方案2：返回pre
}