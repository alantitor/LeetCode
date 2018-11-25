package main

func main() {
	circularQueue := Constructor(3)

	circularQueue.EnQueue(1)
	circularQueue.EnQueue(2)
	//circularQueue.EnQueue(3)
	//circularQueue.EnQueue(4)
	circularQueue.DeQueue()
}

type MyCircularQueue struct {
	data     []int
	capacity int
	length   int
	rear     int
	front    int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	queue := MyCircularQueue{}
	queue.data = make([]int, k)
	queue.capacity = k
	queue.length = 0
	queue.rear = -1
	queue.front = -1
	return queue
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	this.rear += 1
	this.rear %= len(this.data)
	this.data[this.rear] = value
	this.length++

	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	this.front += 1
	this.front %= len(this.data)
	this.data[this.front] = 0
	this.length--

	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	} else {
		return this.data[(this.front+1)%this.capacity]
	}
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	} else {
		return this.data[(this.rear)]
	}
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.length == 0
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.length == this.capacity
}
