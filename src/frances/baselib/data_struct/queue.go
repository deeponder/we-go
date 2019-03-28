package data_struct

import "sync"

type node struct {
	value interface{}
	prev  *node
	next  *node
}

type LinkedQueue struct {
	head *node
	tail *node
	size int
	lock *sync.Mutex
}

//size
func (queue *LinkedQueue) Size() int {
	return queue.size
}

//peak
func (queue *LinkedQueue) Peak() interface{} {
	if queue.head == nil {
		panic("empty queue")
	}
	return queue.head.value
}

//add
func (queue *LinkedQueue) Add(v interface{}) {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	newNode := &node{v, queue.tail, nil}
	if queue.head == nil {
		queue.head = newNode
		queue.tail = newNode
	} else {
		queue.tail.next = newNode
		queue.tail = newNode
	}
	queue.size++
}

//del
func (queue *LinkedQueue) Del() {
	if queue.head == nil {
		panic("empty queue")
	}
	headNode := queue.head
	queue.head = headNode.next
	queue.head.prev = headNode.next

	queue.size--
}

func NewQueue() *LinkedQueue {
	return &LinkedQueue{}
}

//1.考虑线程安全， 加锁
//2.链表、环形数组实现
