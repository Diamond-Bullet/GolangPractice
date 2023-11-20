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

func TestChannelSync1(t *testing.T) {
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

// channel is actually a queue. It's easy(convenient) to implement `Observer`(also known(referred to) as producer-consumer) model using it.
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
func TestWaitGroup(t *testing.T) {
	const goroutineNum = 2

	wg := sync.WaitGroup{}
	wg.Add(2)

	for i := 0; i < goroutineNum; i++ {
		// reserved word `go`, it creates a coroutine.
		// the created coroutine will be added to a wait queue. so goroutine is not totally/completely real-time, but has the characteristic of delayed execution(delayed-execution featured), like `defer`.
		// so it will copy needed parameters while being established.
		go func(index int) {
			defer wg.Done()
			<-time.After(time.Second)
			println("done", index)
		}(i)
	}

	// `wg.Wait()` can be called in multiple places.
	// it's a for-loop which checks a variable if it hits particular conditions.
	go func() {
		wg.Wait()
		println("goroutine exits")
	}()

	wg.Wait()
	println("main exits\n")

	// os.Exit() exit the program directly(straightly, instantly, immediately). no `defer` shall be run.

	// Using Factory Pattern to bind goroutine to a channel.
	// some goroutines are permanently waiting for receiving messages from or sending messages to a channel due to certain bug in our program. it causes memory leaks.
}
