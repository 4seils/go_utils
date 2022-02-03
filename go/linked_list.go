package 4seils_utils

import (
	"fmt"
	"sync"
)

//implement a doubly linked list

//define a node structure
//Each node has a key, value, the address of the upper node, and the address of the lower node
type DoubleNode struct {
	Key   int         //Key
	Value interface{} //value
	Prev  *DoubleNode //The previous node pointer
	Next  *DoubleNode //The next node pointer
}

//define the structure of a double linked list
type DoubleList struct {
	lock     *sync.RWMutex //lock
	Capacity int           // capacit
	Size     int
	Head     *DoubleNode //Head node
	Tail     *DoubleNode //Tail node
}

//implement doubly linked list

// Initialize double linked list
func NewDLL(capacity int) *DoubleList {
	list := new(DoubleList)
	list.Capacity = capacity
	list.lock = new(sync.RWMutex)
	list.Size = 0
	list.Head = nil
	list.Tail = nil
	return list
}

func Node(key int, value interface{}) *DoubleNode {
	return &DoubleNode{
		Key:   key,
		Value: value,
		Prev:  nil,
		Next:  nil,
	}
}

/*
//Add a node to the tail by default
func (list *DoubleList) Append(node *DoubleNode) bool {
	return list.AddTail(node)
}

//add header element
// 1. First determine whether there is a capacity size
// 2. Determine if the header is empty,
// a. add new node if empty
// b. If it is not empty, change the existing node and add the
func (list *DoubleList) AddHead(node *DoubleNode) bool {
	// Determine if the capacity is 0
	list.lock.Lock()
	defer list.lock.Unlock()

	if list.Capacity == list.Size {
		return false
	}

	// Determine whether the head node is nil
	if list.Head == nil {
		list.Head = node
		list.Tail = node
	} else { //There is a head node
		list.Head.Prev = node //Point the previous node of the old head node to the new node
		node.Next = list.Head //The next node of the new head node points to the old head node
		list.Head = node      //Set the new head node
		list.Head.Prev = nil  //Set NIL to the previous node of the new head node
	}
	list.Size++
	return true
}
*/
//add tail element
// append element
// 1. First determine whether there is a capacity size
// 2. Determine if the tail is empty,
// a. add new node if empty
// b. If it is not empty, change the existing node and add the
func (list *DoubleList) AddTail(node *DoubleNode) bool {
	// Determine whether there is capacity,
	list.lock.Lock()
	defer list.lock.Unlock()

	if list.Capacity == list.Size {
		return false
	}

	//Check if the tail is empty
	if list.Tail == nil {
		list.Tail = node
		list.Head = node
	} else {
		//The next node at the old tail points to the new node
		list.Tail.Next = node
		//When appending a new node, first point the upper node of the node to the old tail node
		node.Prev = list.Tail
		//set new tail node
		list.Tail = node
		//The new tail next node is set to empty
		list.Tail.Next = nil
	}
	//Double linked list size +1
	list.Size++
	return true
}

/*
//Add any position element
func (list *DoubleList) Insert(index int, node *DoubleNode) bool {
	// full capacity
	if list.Size == list.Capacity {
		return false
	}
	//if there is no node
	if list.Size == 0 {
		return list.Append(node)
	}
	//If the inserted position is greater than the current length, the tail will be added
	if index > list.Size {
		return list.AddTail(node)
	}
	//If the inserted position is equal to 0, the header is added
	if index == 0 {
		return list.AddHead(node)
	}
	//Remove the node to be inserted at the position
	nextNode := list.Get(index)
	list.lock.Lock()
	defer list.lock.Unlock()
	//Intermediate insertion requires:
	//Assuming that there are A and C nodes, now we want to insert the B node
	// nextNode is the C node,
	//A's lower node should be B, that is, the lower node of C's upper node is B
	nextNode.Prev.Next = node
	//The upper node of B is the upper node of C
	node.Prev = nextNode.Prev
	//The next node of B is C
	node.Next = nextNode
	//The upper node of C is B
	nextNode.Prev = node
	list.Size++
	return true
}
*/

/*
// delete any element
func (list *DoubleList) Remove(node *DoubleNode) *DoubleNode {
	// Determine whether it is the head node
	if node == list.Head {
		return list.RemoveHead()
	}
	// Determine whether it is a tail node
	if node == list.Tail {
		return list.RemoveTail()
	}
	list.lock.Lock()
	defer list.lock.Unlock()
	//The node is an intermediate node
	// then you need:
	// Point the next node pointer of the previous node to the next node
	// Point the previous node pointer of the next node to the previous node
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	list.Size--
	return node
}
*/

// delete the head node
func (list *DoubleList) RemoveHead() *DoubleNode {
	// Determine whether the head node is empty
	if list.Head == nil {
		return nil
	}
	list.lock.Lock()
	defer list.lock.Unlock()
	// take out the head node
	node := list.Head
	// Determine whether the head has the next node
	if node.Next != nil {
		list.Head = node.Next
		list.Head.Prev = nil
	} else { //If there is no next node, it means there is only one node
		list.Head, list.Tail = nil, nil
	}
	list.Size--
	return node
}

/*
// delete the tail node
func (list *DoubleList) RemoveTail() *DoubleNode {
	// Determine whether the tail node is empty
	if list.Tail == nil {
		return nil
	}
	list.lock.Lock()
	defer list.lock.Unlock()
	//Remove the tail node
	node := list.Tail
	//Determine whether the previous one of the tail node exists
	if node.Prev != nil {
		list.Tail = node.Prev
		list.Tail.Next = nil
	} else {
		list.Tail, list.Head = nil, nil
	}
	list.Size--
	return node
}

// get an element
func (list *DoubleList) Get(index int) *DoubleNode {
	//If index = 0, return the head
	if index == 0 {
		return list.Head
	}
	//If it exceeds or equals the current chain size, return the tail
	if index >= list.Size {
		return list.Tail
	}
	//If it is in the middle, you need to loop the linked list of index times
	node := list.Head
	for i := 1; i < index; i++ {
		node = node.Next
	}
	return node
}

//Search for a certain data. Traverse all tables
func (list *DoubleList) Search(key int) *DoubleNode {
	if list.Size == 0 {
		return nil
	}
	// search from the head
	node := list.Head
	// Determine if it is header data.
	if node.Key == key {
		return node
	}
	// Search down for non-head data
	for node != nil {
		node = node.Next
		if node.Key == key {
			return node
		}
	}
	return node
}

//Get the size of the linked list
func (list *DoubleList) GetSize() int {
	return list.Size
}
*/

//print linked list
func (list *DoubleList) Print() {
	if list == nil {
		fmt.Println("list is nil")
		return
	} else if list.Size == 0 {
		fmt.Println("list size is zero")
		return
	}
	p := list.Head
	line := ""
	for p != nil {
		line += fmt.Sprintf("key:%d, value: %v", p.Key, p.Value)
		p = p.Next
		if p != nil {
			line += " => "
		}
	}
	fmt.Println(line)
}

// reverse the linked list
func (list *DoubleList) Reverse() {
	if list == nil || list.Size == 0 {
		fmt.Println("data is nil")
		return
	}
	p := list.Tail
	line := ""
	for p != nil {
		line += fmt.Sprintf("key:%d, value: %v", p.Key, p.Value)
		p = p.Prev
		if p != nil {
			line += " => "
		}
	}
	fmt.Println(line)
}

/*
func main() {
	list := NewDL(10)
	nodes := make([]*DoubleNode, 0)
	for i := 0; i < 10; i++ {
		var node *DoubleNode
		if i%2 == 0 {
			node = Node(i, "test"+strconv.Itoa(i))
		} else {
			node = Node(i, i)
		}

		nodes = append(nodes, node)
	}
	list.Append(nodes[0])
	list.Print()
	list.Append(nodes[1])
	list.Print()
	list.Append(nodes[2])
	list.Print()
	list.RemoveTail()
	list.Print()
	list.RemoveHead()
	list.Print()
	list.Append(nodes[3])
	list.Print()
	list.Insert(2, nodes[4])
	list.Print()
	list.Append(nodes[5])
	list.Print()
	list.Append(nodes[6])
	list.Print()
	list.Reverse()
}


func DoubleList_Print() {
	l := NewDL(5)
	for i := 0; i < 5; i++ {
		node := Node(i, i)
		l.Append(node)
	}
	if l.GetSize() != 5 {
		t.Errorf("want: %d, errValue: %d", 5, l.GetSize())
	}
	l.Print()
}
*/
