package utils

import (
	"fmt"
)

type MessageQueue struct {
	Pending chan bool
	list    *DoubleList
}

func (queue *MessageQueue) Initialize(capacity int) {
	queue.list = NewDLL(capacity)
	queue.Pending = make(chan bool, capacity)
}

func (queue *MessageQueue) Push(key int, val interface{}) bool {
	n := Node(key, val)
	rslt := queue.list.AddTail(n)
	if rslt {
		queue.Pending <- true
		fmt.Println("Push OK")
	} else {
		<-queue.Pending
		fmt.Println("Push NG")
	}
	return rslt
}

func (queue *MessageQueue) Pop() *DoubleNode {
	return queue.list.RemoveHead()
}
