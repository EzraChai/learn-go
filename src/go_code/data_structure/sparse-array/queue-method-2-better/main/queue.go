package main

import "fmt"

type queue struct {
	queueArray [5]int
	front      int
	rear       int
	maxSize    int
}

func (q *queue) SetQueueElement(num int) {
	rearBefore := q.rear
	if q.rear == q.maxSize {
		q.rear = -12
	}
	if q.rear == q.front && q.queueArray[q.rear+1] != 0 {
		q.rear = rearBefore
		fmt.Println("index out of bounds")
		return
	}
	q.rear++
	q.queueArray[q.rear] = num
}

func (q *queue) GetQueueElement() int {
	if q.front == q.maxSize {
		q.front = -1
	}
	if q.front == q.rear {
		fmt.Println("no more elements")
		return 0
	}

	q.front++
	num := q.queueArray[q.front]
	q.queueArray[q.front] = 0
	return num
}

func InitQueue() *queue {
	q := &queue{
		front: -1,
		rear:  -1,
	}
	q.maxSize = len(q.queueArray) - 1
	return q
}

//	Using array to demonstrate the queue
func main() {
	queue := InitQueue()
	for i := 1; i <= 6; i++ {
		queue.SetQueueElement(i)
	}
	fmt.Println(queue.queueArray)
	for i := 0; i < 3; i++ {
		fmt.Println(queue.GetQueueElement())
	}
	fmt.Println(queue.queueArray)
	fmt.Println(queue.GetQueueElement())
	fmt.Println(queue.queueArray)
	fmt.Println(queue.GetQueueElement())
	queue.SetQueueElement(6)
	queue.SetQueueElement(7)
	queue.SetQueueElement(8)
	queue.SetQueueElement(9)
	fmt.Println(queue.queueArray)
	for i := 0; i < 4; i++ {
		fmt.Println(queue.GetQueueElement())
	}

}
