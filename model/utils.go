package model

import "sync"

// This job stores response and error
type Job struct {
	Responses interface{}
	Errors    error
}

type Resource struct {
	sync.Mutex
	count int
}

func (r *Resource) Add() {
	r.Lock()
	r.count++
	r.Unlock()
}

func (r *Resource) Sum() int {
	return r.count
}

type TaskQueue struct {
	items []interface{}
	lock  sync.Mutex
}

func (t *TaskQueue) Add(items interface{}) {
	t.lock.Lock() // lock other writes
	defer t.lock.Unlock()
	t.items = append(t.items, items)
}

func (t *TaskQueue) Remove() interface{} {
	t.lock.Lock()
	defer t.lock.Unlock()

	item := t.items[0]
	t.items = t.items[1:]
	return item
}
