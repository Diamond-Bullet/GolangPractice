package design_pattern

/*
	对象池模式
*/

import (
	"fmt"
	"sync"
)

type Task struct {
	ID   int
	work func(...interface{}) error
}

func (t *Task) Execute() error {
	return t.work()
}

func NewTask(work func(...interface{}) error, ID int) *Task {
	return &Task{
		ID:   ID,
		work: work,
	}
}

// Pool 对象池
type Pool struct {
	WorkerNum   int
	JobsChannel chan *Task
	Wg          *sync.WaitGroup
}

func NewPool(workerNum int) *Pool {
	return &Pool{
		WorkerNum:   workerNum,
		JobsChannel: make(chan *Task, workerNum),
		Wg:          &sync.WaitGroup{},
	}
}

func (p *Pool) PushTask(task *Task) {
	p.JobsChannel <- task
}

func (p *Pool) Work(workerID int) {
	for task := range p.JobsChannel {
		err := task.Execute()
		if err != nil {
			fmt.Println(workerID, "run error:", err)
		} else {
			fmt.Println(workerID, "run success!")
		}

		p.Wg.Done()
	}
}

func (p *Pool) Run() {
	for i := 0; i < p.WorkerNum; i++ {
		go p.Work(i)
	}
}

func (p *Pool) Stop() {
	close(p.JobsChannel)
}
