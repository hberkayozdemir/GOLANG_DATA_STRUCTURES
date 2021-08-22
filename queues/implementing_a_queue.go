package main

import (
	"errors"
	"fmt"
)

// Queue - our representation of a queue data structure
type Queue struct {
	Elements []int
}

// Enqueue -add en element of type int to the end of our queue
func (q *Queue) Enqueue(elem int) {
	q.Elements = append(q.Elements, elem)
}

func (q *Queue) Deqeueue() (int, error) {
	// we need the return first elements
	// update elements slice
	// validate queue not empty control we have to check

	if (len(q.Elements)) == 0 {
		return 0, errors.New("queue is empty")
	}

	var firstElement int
	fmt.Println(firstElement, "   ", q.Elements)
	firstElement, q.Elements = q.Elements[0], q.Elements[1:]
	fmt.Println(firstElement, "   ", q.Elements)

	return firstElement, nil
}

// get the length of the queue

func (q *Queue) Length() int {
	return len(q.Elements)
}

// is empty method

func (q *Queue) IsEmpty() bool {
	/* 	if len(q.Elements) == 0 {
	   		return true
	   	}
	   	return false
	   	 KISALTMA ZAMANI
	*/
	return q.Length() == 0
}

// peek- return first element forum our queue with out updating queue
func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {

		return 0, errors.New("empty queue")
	}
	return q.Elements[0], nil
}
func main() {

	fmt.Println("Queues Section")

	queue := Queue{}
	fmt.Println(queue)
	queue.Enqueue(1)
	fmt.Println(queue)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	fmt.Println(queue)
	elem, _ := queue.Deqeueue()
	fmt.Println(elem)
	fmt.Println(queue)
	peak, _ := queue.Peek()
	fmt.Println(peak)
	fmt.Println(queue.IsEmpty())
	fmt.Println(queue.Length())

}
