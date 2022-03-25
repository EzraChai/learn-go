package main

import "fmt"

type queue struct {
	queueArray [5]int
	front      int
	rear       int
	maxSize    int
}

func (q *queue) setQueueElement(num int) {
	q.rear++
	if q.rear > q.maxSize {
		fmt.Println("index out of bounds")
		return
	}
	q.queueArray[q.rear] = num
}

func (q *queue) getQueueElement() int {
	q.front++
	if q.front > q.rear {
		fmt.Println("no more elements")
		return -1
	}
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
	for i := 0; i < 5; i++ {
		queue.setQueueElement(i)
	}
	fmt.Println(queue.queueArray)
	for i := 0; i < 3; i++ {
		fmt.Println(queue.getQueueElement())
	}
	fmt.Println(queue.queueArray)
	fmt.Println(queue.getQueueElement())
	fmt.Println(queue.queueArray)

}
