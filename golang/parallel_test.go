package golang

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// concurrency-safe Map
// theories, refer to：https://blog.csdn.net/u011957758/article/details/96633984
func TestSyncMap(t *testing.T) {
	key, subKey := "", 50
	m := &sync.Map{}
	SetMap(m, key, subKey)
}

func SetMap(m *sync.Map, key string, subKey int) {
	if subM, ok := m.Load(key); ok {
		e, ok1 := subM.(*sync.Map).Load(subKey)
		if ok1 {
			subM.(*sync.Map).Store(subKey, e.(int)+1)
		} else {
			subM.(*sync.Map).Store(subKey, 1)
		}
	} else {
		newSubM := &sync.Map{}
		newSubM.Store(subKey, 1)
		m.Store(key, newSubM)
	}
}

// synchronization
// mutually exclusive
func TestMutex(t *testing.T) {
	var m sync.Mutex

	m.Lock()
	go func() {
		defer m.Unlock()
		println("hello world")
	}()

	m.Lock()
}

func TestChannelSync(t *testing.T) {
	done := make(chan int)
	go func() {
		time.Sleep(time.Second)
		println("hello world")
		done <- 1 // < done. using `<- done` as an alternative also works, but it can be better to let the recipient block main threading.
	}()

	<-done // done <- 1
}

func TestChannelSync2(t *testing.T) {
	done := make(chan int, 10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("goroutine: %d\n", i)
			done <- 1
		}(i)
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}

// channel itself actually is a queue. It's easy(convenient) to implement `Observer`(also known(referred to) as producer-consumer) model using it.
func TestProduceConsume(t *testing.T) {
	ch := make(chan int, 2)
	go Producer(5, ch)
	go Producer(2, ch)
	go Consumer(ch)

	time.Sleep(4 * time.Second)
}

func Producer(factor int, out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * factor
	}
}

func Consumer(in <-chan int) {
	for v := range in {
		fmt.Printf("consume: %d\n", v)
	}
}

// Publish-Subscribe Model
func TestPubSub(t *testing.T) {
}

// limit(restrain,restrict,control) parallel count, by channel with buffer.
// TryLock，can implement it by channel with a buffer size of 1. it's like semaphore.

// goroutine
func TestGoroutineBasic(t *testing.T) {
	// reserved word `go`, it creates a coroutine.
	// the created coroutine will be added to a wait queue. so goroutine is not totally/completely real-time, but has the characteristic of delayed execution(delayed-execution featured), like `defer`.
	// so it will copy needed parameters while being established.

	// `WaitGroup`'s function `wait`. you can call it in multiple places.
	// it's a for-loop which checks a variable if it hits particular conditions.
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		wg.Wait()
		println("go exit")
	}()

	go func() {
		<-time.After(time.Second)
		println("done")
		wg.Done()
	}()

	wg.Wait()
	println("main arrived\n")

	// runtime.NumCPU() the physical machine's amount of CPU.
	// runtime.GOMAXPROCS() Set Processor `P`'s count.
	// runtime.Gosched() Let current goroutine give up execution right, so other goroutines can be scheduled.
	// runtime.Goexit() Terminate present goroutine. all `defer` registered before are supposed to be executed. if it's called in main.main, the program will panic after all tasks is finished.
	// os.Exit() exit the program directly(straightly, instantly, immediately). no `defer` shall be run.

	// Using Factory Pattern to bind goroutine to a channel.
	// some goroutines is permanently waiting for receiving messages from or sending messages to a channel due to certain bug in our program. it causes memory leaks.
}
