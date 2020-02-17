package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	elements []int
}

func (q *Queue) Push(elem ...int) {
	q.elements = append(q.elements, elem...)
}

func (q *Queue) Pop() (int, error) {

	elemAmount := len(q.elements)

	if elemAmount == 0 {
		return 0, errors.New("No elements in queue")
	}

	popElement := q.elements[0]
	q.elements = q.elements[1:]

	return popElement, nil
}

func main() {
	queue := Queue{
		elements: []int{},
	}

	queue.Push(3, 4, 5)
	fmt.Println(queue)

	if firstElement, err := queue.Pop(); err == nil {
		fmt.Println(firstElement)
	}

	if secondElement, err := queue.Pop(); err == nil {
		fmt.Println(secondElement)
	}

	if thirdElement, err := queue.Pop(); err == nil {
		fmt.Println(thirdElement)
	}

	if notExistElem, err := queue.Pop(); err != nil {
		fmt.Println(notExistElem, err)
	}

}
