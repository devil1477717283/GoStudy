package workerpool

import (
	"sync"
)

type Poll struct {
	capacity int
	active   chan struct{}
	tasks    chan Task
	wg       sync.WaitGroup
	quit     chan struct{}
}

func New(capacity int) *Pool {
	if capacity <= 0 {
		cap
	}
}
