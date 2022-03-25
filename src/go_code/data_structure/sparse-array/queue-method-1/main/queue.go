package main

import "fmt"

type queue struct {
	queueArray [5]int
	front      int
	rear       int
	maxSize    int
}

func (q *queue) SetQueueElement(num int) {
	if q.rear+1 > q.maxSize {
		fmt.Println("index out of bounds")
		return
	}
	q.rear++
	q.queueArray[q.rear] = num
}

func (q *queue) GetQueueElement() int {
	if q.front+1 > q.rear {
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
	for i := 1; i <= 5; i++ {
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
	//queue.SetQueueElement(6)

}
