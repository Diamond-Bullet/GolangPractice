package behavioral_pattern

/*
	Object Pool
*/

import (
	"GolangPractice/pkg/logger"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestObjectPool(t *testing.T) {
	pool := NewPool(10)
	go pool.Run()

	for i := 0; i < 40; i++ {
		pool.Wg.Add(1)
		pool.PushTask(&Task{ID: i, work: func(i ...interface{}) error {
			time.Sleep(time.Second)
			return nil
		}})
	}

	pool.Wg.Wait()
	logger.Info("goroutines: ", runtime.NumGoroutine())
	pool.Stop()

	time.Sleep(time.Second)
	logger.Info("goroutines: ", runtime.NumGoroutine())
}

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
			logger.Info(workerID, "run error:", err)
		} else {
			logger.Info(workerID, "run success!")
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
