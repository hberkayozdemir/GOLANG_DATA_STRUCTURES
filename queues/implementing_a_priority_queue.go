package main

import (
	"errors"
	"fmt"
)

type PriorityQueue struct {
	High []int
	Low  []int
}

// enpritortiyqueue
func (q *PriorityQueue) Enqueue(elem int, highPriority bool) {
	if highPriority {
		q.High = append(q.High, elem)
	} else {
		q.Low = append(q.Low, elem)
	}
}

// depriorityqueue

func (q *PriorityQueue) Deqeueue() (int, error) {

	if len(q.High) != 0 {
		var firstElement int
		firstElement, q.High = q.High[0], q.High[1:]
		return firstElement, nil
	}
	if len(q.Low) != 0 {
		var firstElement int
		firstElement, q.Low = q.Low[0], q.Low[1:]
		return firstElement, nil
	}

	return 0, errors.New("priorityqueue empty")

}

func (q *PriorityQueue) Peek() (int, error) {

	if len(q.High) != 0 {
		return q.High[0], nil
	}
	if len(q.Low) != 0 {
		return q.Low[0], nil
	}

	return 0, errors.New("priorityqueue empty")
}

// length

func (q *PriorityQueue) Length() int {
	return len(q.Low) + len(q.High)
}

func (q *PriorityQueue) isEmpty() bool {

	return q.Length() == 0
}

func main() {

	fmt.Println(" PriorityQueues section")

	priorityQueue := PriorityQueue{}
	fmt.Println(priorityQueue)
	priorityQueue.Enqueue(1, true)
	priorityQueue.Enqueue(1, false)
	priorityQueue.Enqueue(2, false)
	priorityQueue.Enqueue(3, false)
	priorityQueue.Enqueue(4, true)
	fmt.Println(priorityQueue)
	elem, _ := priorityQueue.Deqeueue()
	fmt.Println(elem)
	priorityQueue.Deqeueue()
	fmt.Println(priorityQueue)
	priorityQueue.Deqeueue()
	fmt.Println(priorityQueue)
	priorityQueue.Deqeueue()
	fmt.Println(priorityQueue)
}
