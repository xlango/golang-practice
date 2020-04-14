package main

import "fmt"

//循环队列

const maxSize = 3

//SqQueue 结构体定义
type SqQueue struct {
	data  [maxSize]interface{}
	front int
	rear  int
}

//New 新建空队列
func New() *SqQueue {
	return &SqQueue{
		front: 0,
		rear:  0,
	}
}

// Length 队列长度
func (q *SqQueue) Length() interface{} {
	return (q.rear - q.front + maxSize) % maxSize
}

// Enqueue 入队
func (q *SqQueue) Enqueue(e interface{}) error {
	if (q.rear+1)%maxSize == q.front {
		return fmt.Errorf("quque is full")
	}
	q.data[q.rear] = e
	q.rear = (q.rear + 1) % maxSize
	return nil
}

// Dequeue 出队
func (q *SqQueue) Dequeue() (e interface{}, err error) {
	if q.rear == q.front {
		return e, fmt.Errorf("quque is empty")
	}
	e = q.data[q.front]
	q.front = (q.front + 1) % maxSize
	return e, nil
}

func main() {
	q := New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	e, _ := q.Dequeue()
	fmt.Println(e)
	e1, _ := q.Dequeue()
	fmt.Println(e1)
	e2, _ := q.Dequeue()
	fmt.Println(e2)
}
