package main

import "fmt"

/**
203 移除链表元素
https://mp.weixin.qq.com/s/slM1CH5Ew9XzK93YOQYSjA

题意：删除链表中等于给定值 val 的所有节点。

Input: head = [1,2,6,3,4,5,6], val = 6
Output: [1,2,3,4,5]


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
	listNode2.Val = 6
	listNode1.Next = &listNode2

	listNode3 := NewListNode()
	listNode3.Val = 3
	listNode2.Next = &listNode3

	listNode4 := NewListNode()
	listNode4.Val = 4
	listNode3.Next = &listNode4

	listNode5 := NewListNode()
	listNode5.Val = 5
	listNode4.Next = &listNode5

	listNode6 := NewListNode()
	listNode6.Val = 6 // listNode6的Next已经默认为null
	listNode5.Next = &listNode6

	fmt.Println("--原始链表--")

	for i := &listNode0; i != nil; i = i.Next{
		fmt.Println(i.Val)
	}

	fmt.Println("--删除val后的链表--")

	listNodeChanged := removeElements(&listNode0, 6)

	for i := listNodeChanged; i != nil; i = i.Next{
		fmt.Println(i.Val)
	}
}

func removeElements(listHead *listNode, target int) *listNode{
	// 创建虚拟节点作为head，head后连接到原来的链表上。可以避免删除头节点时，需要让头结点指向下一个节点的单独处理过程
	dummyHead := NewListNode()
	dummyHead.Next = listHead // 但是不能用dummyHead作为起点向后遍历，因为需要返回dummyHead->next作为头结点

	cur := dummyHead.Next

	for cur.Next != nil {
		if cur.Next.Val == target {
			cur.Next = cur.Next.Next
			// Go不需要像cpp：利用tmp指针指向cur.Next，然后回收cur.Next的内存
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}
