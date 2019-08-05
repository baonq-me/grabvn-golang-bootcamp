package main

import (
	"fmt"
	"time"
)

type Printer interface {
	Print(index int)
}

type Job Printer

type Worker struct {
	index int
	job   chan Job
	pool  chan chan Job
}

func (w Worker) start() {
	for {
		w.pool <- w.job
		value := <-w.job
		value.Print(w.index)
	}
}

type Pool struct {
	pool    chan chan Job
	job     chan Job
	workers []Worker
}

func NewWorkerPool(count int) *Pool {
	p := Pool{
		pool: make(chan chan Job, count),
		job:  make(chan Job),
	}

	for i := 0; i < count; i++ {
		p.workers = append(p.workers, Worker{
			index: i + 1,
			pool:  p.pool,
			job:   make(chan Job),
		})
	}

	for _, worker := range p.workers {
		go worker.start()
	}

	go p.start()

	return &p
}

func (p *Pool) dispatch(job Job) {
	p.job <- job
}

func (p *Pool) start() {
	for {
		job := <-p.job     // Get job
		worker := <-p.pool // Get worker
		worker <- job      // Assign job to worker
	}
}

type Int int

func (i Int) Print(index int) {
	fmt.Println(index, "- got:", i)
}

func main() {
	p := NewWorkerPool(10)
	for i := 1; i < 20; i++ {
		p.dispatch(Int(i))
	}

	time.Sleep(2 * time.Second)
}
