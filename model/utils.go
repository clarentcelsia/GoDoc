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
