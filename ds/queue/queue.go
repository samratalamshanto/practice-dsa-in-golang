package main

import "fmt"

type Queue struct {
	data []int
}

func (q *Queue) Enqueue(val int) {
	q.data = append(q.data, val)
}

func (q *Queue) Dequeue() {
	if len(q.data) == 0 {
		return
	}
	q.data = q.data[1:]
}

func (q *Queue) Front() int {
	if len(q.data) == 0 {
		return -1
	}
	return q.data[0]
}

func main() {
	queue := &Queue{}
	fmt.Println(queue.data)

	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	fmt.Println(queue.data)

	front := queue.Front()
	fmt.Println(front)

	queue.Dequeue()
	fmt.Println(queue.data)
}
