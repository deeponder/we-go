/*
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
 */
package main

type node struct {
	value int
	prev *node
	next *node
}
func mergeTwoSortedLists(l1 *node, l2 *node) *node {
	//判断有否为nil的链
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	//初始化， 输出链的头
	var head, node *node
	if l1.value < l2.value {
		head = l1
		node = l1
		l1 = l1.next
	} else {
		head = l2
		node = l2
		l2 = l2.next
	}

	//循环比较
	for l1 != nil && l2 != nil {
		if l1.value < l2.value {
			node.next = l1
			l1 =  l1.next
		} else {
			node.next = l2
			l2 = l2.next
		}
		node = node.next
	}

	//接上不为nil的链
	if l1 == nil {
		node.next = l2
	}

	if l2 == nil {
		node.next = l1
	}
	return head
}
func main() {
		
}
