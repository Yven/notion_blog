package lib

import (
	"container/list"
	"fmt"
)

type ReqQueue struct {
	queue *list.List
}

type ReqInfo struct {
	Id    string
	times int16
}

func (c *ReqQueue) Enqueue(value *ReqInfo) {
	c.queue.PushBack(value)
}

func (c *ReqQueue) Dequeue() (value *ReqInfo, error error) {
	if c.queue.Len() > 0 {
		ele := c.queue.Front()
		value = c.queue.Remove(ele).(*ReqInfo)
		return value, nil
	}
	return &ReqInfo{}, fmt.Errorf("Pop Error: Queue is empty")
}

func (c *ReqQueue) Front() (*ReqInfo, error) {
	if c.queue.Len() > 0 {
		if val, ok := c.queue.Front().Value.(*ReqInfo); ok {
			return val, nil
		}
		return &ReqInfo{}, fmt.Errorf("Peep Error: Queue Datatype is incorrect")
	}
	return &ReqInfo{}, fmt.Errorf("Peep Error: Queue is empty")
}

func (c *ReqQueue) Size() int {
	return c.queue.Len()
}

func (c *ReqQueue) Empty() bool {
	return c.queue.Len() == 0
}
