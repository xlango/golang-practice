package pool

import (
	"fmt"
)

type Task struct {
	f func() error
}

func NewTask(f func() error) *Task {
	return &Task{
		f: f,
	}
}

func (t *Task) Execute() {
	t.f()
}

type Pool struct {
	EntryChan   chan *Task
	JobChan     chan *Task
	WorkerCount int
}

func NewPool(cap int) *Pool {
	return &Pool{
		EntryChan:   make(chan *Task),
		JobChan:     make(chan *Task),
		WorkerCount: cap,
	}
}

func (p *Pool) Worker(workerId int) {
	for task := range p.JobChan {
		task.Execute()
		fmt.Printf("Worker Id:%d excuted!", workerId)
	}
}

func (p *Pool) Run() {
	for i := 0; i < p.WorkerCount; i++ {
		go p.Worker(i)
	}

	for task := range p.EntryChan {
		p.JobChan <- task
	}
}
