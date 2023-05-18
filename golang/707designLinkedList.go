/*
707. 设计链表

你可以选择使用单链表或者双链表，设计并实现自己的链表。

单链表中的节点应该具备两个属性：val 和 next 。val 是当前节点的值，next 是指向下一个节点的指针/引用。

如果是双向链表，则还需要属性 prev 以指示链表中的上一个节点。假设链表中的所有节点下标从 0 开始。

实现 MyLinkedList 类：

MyLinkedList() 初始化 MyLinkedList 对象。
int get(int index) 获取链表中下标为 index 的节点的值。如果下标无效，则返回 -1 。
void addAtHead(int val) 将一个值为 val 的节点插入到链表中第一个元素之前。在插入完成后，新节点会成为链表的第一个节点。
void addAtTail(int val) 将一个值为 val 的节点追加到链表中作为链表的最后一个元素。
void addAtIndex(int index, int val) 将一个值为 val 的节点插入到链表中下标为 index 的节点之前。如果 index 等于链表的长度，那么该节点会被追加到链表的末尾。如果 index 比长度更大，该节点将 不会插入 到链表中。
void deleteAtIndex(int index) 如果下标有效，则删除链表中下标为 index 的节点。
 

示例：

输入
["MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get", "deleteAtIndex", "get"]
[[], [1], [3], [1, 2], [1], [1], [1]]
输出
[null, null, null, null, 2, null, 3]

解释
MyLinkedList myLinkedList = new MyLinkedList();
myLinkedList.addAtHead(1);
myLinkedList.addAtTail(3);
myLinkedList.addAtIndex(1, 2);    // 链表变为 1->2->3
myLinkedList.get(1);              // 返回 2
myLinkedList.deleteAtIndex(1);    // 现在，链表变为 1->3
myLinkedList.get(1);              // 返回 3
 

提示：

0 <= index, val <= 1000
请不要使用内置的 LinkedList 库。
调用 get、addAtHead、addAtTail、addAtIndex 和 deleteAtIndex 的次数不超过 2000 。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/design-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

这道题目设计链表的五个接口：

获取链表第index个节点的数值
在链表的最前面插入一个节点
在链表的最后面插入一个节点
在链表第index个节点前面插入一个节点
删除链表的第index个节点
可以说这五个接口，已经覆盖了链表的常见操作，是练习链表操作非常好的一道题目
*/

package main

import "fmt"


type SingleNode struct {
	Val int
	Next *SingleNode
}

type SingleLinkedList struct {
	dummyHead *SingleNode
	Size int
}

func Init() (*SingleLinkedList){
	list := new(SingleLinkedList)
	list.dummyHead = new(SingleNode)
	list.Size = 0
	return list
}

func (list *SingleLinkedList) printLinkedList() {
	for cur:=list.dummyHead.Next; cur!=nil; cur=cur.Next {
		fmt.Println(cur.Val)
	}
}

func (list *SingleLinkedList) addAtHead(val int) {
	newNode := new(SingleNode)
	newNode.Val = val
	newNode.Next = list.dummyHead.Next
	list.dummyHead.Next = newNode
	list.Size++
}

func (list *SingleLinkedList) addAtTail(val int) {
	cur := list.dummyHead
	for ;cur.Next!=nil;cur=cur.Next {}
	newNode := new(SingleNode)
	newNode.Val = val
	cur.Next = newNode
	list.Size++
}

func (list *SingleLinkedList) addAtIndex(val int, index int) {
	if index < 0 {
		index = 0
	} else if index > list.Size {
		return 
	}
	cur := list.dummyHead
	for i:=0;cur.Next!=nil && i<index-1;cur, i=cur.Next, i+1 {}
	newNode := &SingleNode{Val:val}
	newNode.Next = cur.Next
	cur.Next = newNode
	list.Size++
}

func (list *SingleLinkedList) get(index int) int{
	cur := list.dummyHead.Next
	for ; cur!=nil && index>1; cur=cur.Next{
		index--
	}
	if cur == nil {
		return -1
	}
	return cur.Val
}

func main() {
	l := Init()
	l.addAtHead(1)
	l.addAtTail(3)
	l.addAtIndex(2, 2)
	l.addAtIndex(0, 0)
	l.printLinkedList()
}
