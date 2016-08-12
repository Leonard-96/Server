// FIFOQueue project FIFOQueue.go
package FIFOQueue

import (
	"container/list"
	"fmt"
	"reflect"
)

type Queue struct {
	sem  chan int
	list *list.List
}

var tFunc func(val interface{}) bool

func NewQueue() *Queue {
	sem := make(chan int, 1)
	list := list.New()
	return &Queue{sem, list}
}

func (q *Queue) Size() int {
	return q.list.Len()
}

func (q *Queue) Enqueue(val interface{}) *list.Element {
	q.sem <- 1
	e := q.list.PushFront(val)
	<-q.sem
	return e
}

func (q *Queue) Dequeue() *list.Element {
	q.sem <- 1
	e := q.list.Back()
	q.list.Remove(e)
	<-q.sem
	return e
}

func (q *Queue) Query(queryFunc interface{}) *list.Element {
	q.sem <- 1
	e := q.list.Front()
	for e != nil {
		if reflect.TypeOf(queryFunc) == reflect.TypeOf(tFunc) {
			if queryFunc.(func(val interface{}) bool)(e.Value) {
				<-q.sem
				return e
			}
		} else {
			<-q.sem
			return nil
		}
		e = e.Next()
	}
	<-q.sem
	return nil
}

func (q *Queue) Contain(val interface{}) bool {
	q.sem <- 1
	e := q.list.Front()
	for e != nil {
		if e.Value == val {
			<-q.sem
			return true
		} else {
			e = e.Next()
		}
	}
	<-q.sem
	return false
}
