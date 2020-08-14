package main

import "time"

type WebhookMessage struct {
	message   string
	timestamp time.Time
	next      *WebhookMessage
}

type List struct {
	size uint
	head *WebhookMessage
	tail *WebhookMessage
}

func (list *List) Init() {
	(*list).size = 0
	(*list).head = nil
	(*list).tail = nil
}

func (list *List) Append(node *WebhookMessage) bool {
	if node == nil {
		return false
	}
	(*node).next = nil
	if (*list).size == 0 {
		(*list).head = node
	} else {
		oldTail := (*list).tail
		(*oldTail).next = node
	}
	(*list).tail = node
	(*list).size++
	return true
}

func (list *List) Get(i uint) *WebhookMessage {
	if i >= (*list).size {
		return nil
	}
	item := (*list).head
	var j = uint(0)
	for ; j < i; j++ {
		item = (*item).next
	}
	return item
}

func (list *List) RemoveHead() bool {
	if (*list).size == 0 {
		return false
	}
	(*list).head = (*list).head.next
	(*list).size--
	return true
}
